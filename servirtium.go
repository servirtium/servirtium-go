package servirtium

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var templateContent = `
{{- $RequestMethod := .RequestMethod }}
{{- $RequestURLPath := .RequestURLPath }}
{{- $RequestHeader := .RequestHeader }}
{{- $RequestBody := .RequestBody }}
{{- $ResponseHeader := .ResponseHeader }}
{{- $ResponseBody := .ResponseBody }}
{{- $ResponseStatus := .ResponseStatus }}
{{- $RecordSequence := .RecordSequence }}
{{- $ResponseContentType := .ResponseContentType -}}
## Interaction {{ $RecordSequence }}: {{ $RequestMethod }} {{- $RequestURLPath }}

### Request headers recorded for playback:

` + "```" + `
{{ range $key, $values := $RequestHeader }}
  {{- $key}}: {{ range $value := $values}} {{- $value }} {{- end }}
{{ end }}
` + "```" + `

### Request body recorded for playback ():

` + "```" + `
{{ $RequestBody }}
` + "```" + `


### Response headers recorded for playback:

` + "```" + `
{{ range $key, $values := $ResponseHeader }}
  {{- $key}}: {{ range $value := $values}} {{- $value }} {{- end }}
{{ end }}
` + "```" + `

### Response body recorded for playback ({{- $ResponseStatus }}: {{ $ResponseContentType }}):

` + "```" + `
{{ $ResponseBody }}
` + "```" + `

`

// Impl ...
type Impl struct {
	ServerPlayback      *http.Server
	ServerRecord        *http.Server
	interactionSequence int64
	content             string
	// caller request
	callerRequestHeadersRemoval     []string
	callerRequestHeaderReplacements map[*regexp.Regexp]string
	callerRequestBodyReplacement    map[*regexp.Regexp]string
	// record request
	recordRequestHeadersRemoval     []string
	recordRequestHeaderReplacements map[*regexp.Regexp]string
	recordRequestBodyReplacement    map[*regexp.Regexp]string
	// caller response
	callerResponseHeadersRemoval     []string
	callerResponseHeaderReplacements map[*regexp.Regexp]string
	callerResponseBodyReplacement    map[*regexp.Regexp]string
	// record response
	recordResponseHeadersRemoval     []string
	recordResponseHeaderReplacements map[*regexp.Regexp]string
	recordResponseBodyReplacement    map[*regexp.Regexp]string
}

// NewServirtium ...
func NewServirtium() *Impl {
	lastError = nil
	return &Impl{
		interactionSequence:          0,
		content:                      "",
		callerRequestHeadersRemoval:  []string{},
		recordRequestHeadersRemoval:  []string{},
		callerResponseHeadersRemoval: []string{},
		recordResponseHeadersRemoval: []string{},
	}
}

func (s *Impl) StartPlayback(recordFileName string, servirtiumPort int) {
	s.initPlaybackServerOnPort(recordFileName, servirtiumPort)
	go s.ServerPlayback.ListenAndServe()
}

func (s *Impl) EndPlayback() {
	s.ServerPlayback.Shutdown(context.TODO())
}

func (s *Impl) GetPlaybackURL() string {
	return fmt.Sprintf("http://%s", s.ServerPlayback.Addr)
}

func (s *Impl) initPlaybackServerOnPort(recordFileName string, servirtiumPort int) {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.playbackHandler(recordFileName))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "OPTION", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})
	srv := &http.Server{
		Handler: c.Handler(r),
		Addr:    "127.0.0.1:" + strconv.Itoa(servirtiumPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.ServerPlayback = srv
}

func (s *Impl) getInteraction(data string, interactionSequence int64) string {
	interactions := strings.Split(fmt.Sprintf("\n%s", data), "\n## Interaction ")
	return interactions[interactionSequence+1]
}

func (s *Impl) parseHeaders(content string) map[string]string {
	var (
		headers = map[string]string{}
	)
	newLineReplaceByCommaContent := strings.ReplaceAll(content, "\n", ",")
	headersContents := strings.Split(newLineReplaceByCommaContent, ",")
	for _, v := range headersContents {
		replaceByComma := strings.ReplaceAll(v, ": ", ",")
		keysAndValues := strings.Split(replaceByComma, ",")
		keyIndex := 0
		valueIndex := 1
		isValidHeaders := len(keysAndValues) > 1
		if isValidHeaders {
			headers[keysAndValues[keyIndex]] = keysAndValues[valueIndex]
		}
	}
	return headers
}

func (s *Impl) getPlaybackResponse(data string) (string, map[string]string, string) {
	var (
		responseHeaders map[string]string
		responseBody    string
		sections        = strings.Split(data, "\n### ")
		statusCode      string
	)
	for _, v := range sections {
		if strings.HasPrefix(v, "Response headers recorded for playback") {
			headerContent := strings.Split(v, "```")[1]
			responseHeaders = s.parseHeaders(headerContent)
		}
		if strings.HasPrefix(v, "Response body recorded for playback") {
			responseBody = strings.TrimSpace(strings.Split(v, "```")[1])
			statusPart := strings.Split(v, "(")[1]
			statusCode = strings.Split(statusPart, ": ")[0]
		}
	}
	return responseBody, responseHeaders, statusCode
}

func removeHeadersPlayback(headers map[string]string, deleteItems []string) map[string]string {
	for _, v := range deleteItems {
		delete(headers, v)
	}
	return headers
}

func replaceHeadersPlayback(headers map[string]string, replaceItems map[*regexp.Regexp]string) map[string]string {
	for regex, replaceValue := range replaceItems {
		for k, v := range headers {
			line := fmt.Sprintf("%s: %s", k, v)
			line1 := regex.ReplaceAllString(line, replaceValue)
			if line1 != line {
				headerReplacedValue := strings.Split(line1, ": ")[1]
				headers[k] = headerReplacedValue
			}
		}
	}
	return headers
}

var lastError error = nil

func (s *Impl) GetLastError() error {
	return lastError
}

func (s *Impl) playbackHandler(recordFileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		workingPath, err := os.Getwd()
		if err != nil {
			lastError = err
			log.Fatal(err)
		}
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName))
		if err != nil {
			lastError = err
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}
		interaction := s.getInteraction(string(data), s.interactionSequence)
		body, headers, status := s.getPlaybackResponse(interaction)
		callerResponseBody := replaceContent(body, s.callerResponseBodyReplacement)
		callerResponseHeaders := removeHeadersPlayback(headers, s.callerResponseHeadersRemoval)
		callerResponseHeaders = replaceHeadersPlayback(callerResponseHeaders, s.callerResponseHeaderReplacements)
		for k, v := range callerResponseHeaders {
			w.Header().Set(k, v)
		}
		w.Header().Del("Content-Length")
		s.interactionSequence = s.interactionSequence + 1
		statusCode, err := strconv.Atoi(status)
		if err != nil {
			lastError = err
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(statusCode)

		switch w.Header().Get("Content-Encoding") {
		case "gzip":
			var b bytes.Buffer
			gz := gzip.NewWriter(&b)
			if _, err := gz.Write([]byte(callerResponseBody)); err != nil {
				lastError = err
				log.Fatal(err)
			}
			if err := gz.Close(); err != nil {
				lastError = err
				log.Fatal(err)
			}
			i := b.Bytes()
			if w.Header().Get("Content-Length") != "" {
				w.Header().Set("Content-Length", fmt.Sprintf("%d", i))
			}
			w.Write(i)
		default:
			if w.Header().Get("Content-Length") != "" {
				w.Header().Set("Content-Length", fmt.Sprintf("%d", utf8.RuneCountInString(callerResponseBody)))
			}
			w.Write([]byte(callerResponseBody))
		}
		_, _ = w.Write([]byte(callerResponseBody))
	}
}

func (s *Impl) StartRecord(apiURL string, servirtiumPort int) {
	s.initRecordServerOnPort(apiURL, servirtiumPort)
	go s.ServerRecord.ListenAndServe()
}

func (s *Impl) WriteRecord(recordFileName string) {
	workingPath, err := os.Getwd()
	if err != nil {
		lastError = err
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName)
	markdownExists := s.checkMarkdownExists(filePath)
	if !markdownExists {
		os.Create(filePath)
	}
	err = ioutil.WriteFile(filePath, []byte(s.content), os.ModePerm)
	if err != nil {
		lastError = err
		log.Fatal(err)
	}
}

func (s *Impl) EndRecord() {
	s.ServerRecord.Shutdown(context.TODO())
}

func (s *Impl) GetRecordURL() string {
	return fmt.Sprintf("http://%s", s.ServerRecord.Addr)
}

func (s *Impl) initRecordServerOnPort(apiURL string, servirtiumPort int) {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.recordHandler(apiURL))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})
	srv := &http.Server{
		Handler: c.Handler(r),
		Addr:    "127.0.0.1:" + strconv.Itoa(servirtiumPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.ServerRecord = srv
}

type recordData struct {
	RecordSequence      int64
	RequestMethod       string
	RequestURLPath      string
	RequestHeader       map[string][]string
	RequestBody         string
	ResponseHeader      map[string][]string
	ResponseBody        string
	ResponseStatus      int
	ResponseContentType string
}

func removeHeaders(header http.Header, deleteItems []string) http.Header {
	for _, v := range deleteItems {
		header.Del(v)
	}
	return header
}

func replaceHeaders(headers http.Header, replaceItems map[*regexp.Regexp]string) http.Header {
	for regex, replaceValue := range replaceItems {
		for k, v := range headers {
			headerValue := strings.Join(v, ",")
			line := fmt.Sprintf("%s: %s", k, headerValue)
			line1 := regex.ReplaceAllString(line, replaceValue)
			if line1 != line {
				headerReplacedValue := strings.Split(line1, ": ")[1]
				headers.Set(k, headerReplacedValue)
			}
		}
	}
	return headers
}

func replaceContent(content string, maskItems map[*regexp.Regexp]string) string {
	newContent := content
	for k, v := range maskItems {
		newContent = k.ReplaceAllString(newContent, v)
	}
	return newContent
}

func (s *Impl) recordHandler(apiURL string) func(w http.ResponseWriter, r *http.Request) {
	lastError = nil
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone request body
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			lastError = err
			http.Error(w, "Servirtium Recorder/1: "+err.Error(), http.StatusInternalServerError)
			return
		}
		url := fmt.Sprintf("%s%s", apiURL, r.RequestURI)

		// Mutate caller request headers and body
		proxyRequestBody := replaceContent(string(requestBody), s.callerRequestBodyReplacement)
		newCallerProxyRequestHeader := removeHeaders(r.Header, s.callerRequestHeadersRemoval)
		newCallerProxyRequestHeader = replaceHeaders(newCallerProxyRequestHeader, s.callerRequestHeaderReplacements)
		if newCallerProxyRequestHeader.Get("Content-Length") != "" {
			newCallerProxyRequestHeader.Set("Content-Length", fmt.Sprintf("%d", utf8.RuneCountInString(proxyRequestBody)))
		}

		// Make a request
		proxyRequest, err := http.NewRequest(r.Method, url, bytes.NewReader([]byte(proxyRequestBody)))
		if err != nil {
			lastError = err
			log.Fatal(err)
		}
		proxyRequest.Header = newCallerProxyRequestHeader
		response, err := http.DefaultClient.Do(proxyRequest)
		if err != nil {
			lastError = err
			http.Error(w, "Servirtium Recorder/2: "+err.Error(), http.StatusInternalServerError)
			return
		}
		var rspReader io.ReadCloser
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			rspReader, _ = gzip.NewReader(response.Body)
		default:
			rspReader = response.Body
		}
		responseBody, err := io.ReadAll(rspReader)
		if err != nil {
			lastError = err
			http.Error(w, "Servirtium Recorder/3: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Mutate record request headers and body
		newRecordRequestBody := replaceContent(proxyRequestBody, s.recordRequestBodyReplacement)
		newRecordRequestHeaders := removeHeaders(proxyRequest.Header, s.recordRequestHeadersRemoval)
		newRecordRequestHeaders = replaceHeaders(newRecordRequestHeaders, s.recordRequestHeaderReplacements)

		// Modify record response headers and body
		newRecordResponseBody := replaceContent(string(responseBody), s.recordResponseBodyReplacement)
		newRecordResponseHeaders := removeHeaders(response.Header, s.recordResponseHeadersRemoval)
		newRecordResponseHeaders = replaceHeaders(newRecordResponseHeaders, s.recordResponseHeaderReplacements)

		s.record(recordData{
			RequestURLPath:      r.URL.Path,
			RequestMethod:       r.Method,
			RequestHeader:       newRecordRequestHeaders,
			RequestBody:         newRecordRequestBody,
			ResponseHeader:      newRecordResponseHeaders,
			ResponseBody:        newRecordResponseBody,
			ResponseContentType: response.Header.Get("Content-Type"),
			ResponseStatus:      response.StatusCode,
		})

		// Mutate caller response headers and body
		newCallerResponseBody := replaceContent(newRecordResponseBody, s.callerResponseBodyReplacement)
		newCallerResponseHeaders := removeHeaders(newRecordResponseHeaders, s.recordResponseHeadersRemoval)
		newCallerResponseHeaders = replaceHeaders(newCallerResponseHeaders, s.callerResponseHeaderReplacements)

		// response.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(newCallerResponseBody)))
		defer func() {
			_ = response.Body.Close()
		}()
		for k, v := range newCallerResponseHeaders {
			responseHeaderValue := strings.Join(v, ",")
			w.Header().Set(k, responseHeaderValue)
		}
		w.WriteHeader(response.StatusCode)
		switch w.Header().Get("Content-Encoding") {
		case "gzip":
			var bytesBuffer bytes.Buffer
			gz := gzip.NewWriter(&bytesBuffer)
			if _, err := gz.Write([]byte(newCallerResponseBody)); err != nil {
				lastError = err
				http.Error(w, "Servirtium Recorder/4: "+err.Error(), http.StatusInternalServerError)
				return
			}
			if err := gz.Close(); err != nil {
				lastError = err
				http.Error(w, "Servirtium Recorder/5: "+err.Error(), http.StatusInternalServerError)
				return
			}
			contentBytes := bytesBuffer.Bytes()
			w.Header().Set("Content-Length", fmt.Sprintf("%d", utf8.RuneCount(contentBytes)))
			w.Write(contentBytes)
		default:
			w.Header().Set("Content-Length", fmt.Sprintf("%d", utf8.RuneCountInString(newCallerResponseBody)))
			w.Write([]byte(newCallerResponseBody))
		}
	}
}

func (s *Impl) checkMarkdownExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		lastError = err
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (s *Impl) appendContentInFile(currentContent, newContent string) string {
	if s.interactionSequence == 0 {
		return newContent
	}
	finalContent := fmt.Sprintf("%s\n%s", currentContent, newContent)
	return finalContent
}

func (s *Impl) record(params recordData) {
	tmpl, err := template.New("template").Parse(templateContent)
	if err != nil {
		lastError = err
		log.Fatal(err)
	}
	data := recordData{
		RecordSequence:      s.interactionSequence,
		RequestMethod:       params.RequestMethod,
		RequestURLPath:      params.RequestURLPath,
		RequestHeader:       params.RequestHeader,
		RequestBody:         params.RequestBody,
		ResponseHeader:      params.ResponseHeader,
		ResponseBody:        params.ResponseBody,
		ResponseContentType: params.ResponseContentType,
		ResponseStatus:      params.ResponseStatus,
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	newContent := buffer.Bytes()
	finalContent := s.appendContentInFile(s.content, string(newContent))
	s.content = finalContent
	s.interactionSequence = s.interactionSequence + 1
}

func (s *Impl) SetCallerRequestHeadersRemoval(headers []string) {
	s.callerRequestHeadersRemoval = headers
}

func (s *Impl) SetCallerRequestHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.callerRequestHeaderReplacements = replaceRegex
}

func (s *Impl) SetCallerRequestBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.callerRequestBodyReplacement = replaceRegex
}

func (s *Impl) SetRecordRequestHeadersRemoval(headers []string) {
	s.recordRequestHeadersRemoval = headers
}

func (s *Impl) SetRecordRequestHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.recordRequestHeaderReplacements = replaceRegex
}

func (s *Impl) SetRecordRequestBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.recordRequestBodyReplacement = replaceRegex
}

func (s *Impl) SetCallerResponseHeadersRemoval(headers []string) {
	s.callerResponseHeadersRemoval = headers
}

func (s *Impl) SetCallerResponseHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.callerResponseHeaderReplacements = replaceRegex
}

func (s *Impl) SetCallerResponseBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.callerResponseBodyReplacement = replaceRegex
}

func (s *Impl) SetRecordResponseHeadersRemoval(headers []string) {
	s.recordResponseHeadersRemoval = headers
}

func (s *Impl) SetRecordResponseHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.recordResponseHeaderReplacements = replaceRegex
}

func (s *Impl) SetRecordResponseBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.recordResponseBodyReplacement = replaceRegex
}
