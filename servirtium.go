package servirtium

import (
	"bytes"
	"context"
	"fmt"
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
	return &Impl{
		interactionSequence:          0,
		content:                      "",
		callerRequestHeadersRemoval:  []string{},
		recordRequestHeadersRemoval:  []string{},
		callerResponseHeadersRemoval: []string{},
		recordResponseHeadersRemoval: []string{},
	}
}

// StartPlayback ...
func (s *Impl) StartPlayback(recordFileName string, servirtiumPort int) {
	s.initServerPlaybackOnPort(recordFileName, servirtiumPort)
	log.Fatal(s.ServerPlayback.ListenAndServe())
}

// EndPlayback ...
func (s *Impl) EndPlayback() {
	s.ServerPlayback.Shutdown(context.TODO())
}

func (s *Impl) initServerPlayback(recordFileName string) {
	s.initServerPlaybackOnPort(recordFileName, 61417)
}

func (s *Impl) initServerPlaybackOnPort(recordFileName string, servirtiumPort int) {
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
			statusCode = strings.Split(strings.Split(v, "(")[1], ": ")[0]
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

func (s *Impl) playbackHandler(recordFileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		workingPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName))
		if err != nil {
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
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(callerResponseBody))
		return
	}
}

// StartRecord ...
func (s *Impl) StartRecord(apiURL string, servirtiumPort int) {
	s.initRecordServerOnPort(apiURL, servirtiumPort)
	log.Fatal(s.ServerRecord.ListenAndServe())
}

// WriteRecord ...
func (s *Impl) WriteRecord(recordFileName string) {
	workingPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName)
	markdownExists := s.checkMarkdownExists(filePath)
	if !markdownExists {
		os.Create(filePath)
	}
	err = ioutil.WriteFile(filePath, []byte(s.content), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// EndRecord ...
func (s *Impl) EndRecord() {
	s.ServerRecord.Shutdown(context.TODO())
}

func (s *Impl) initRecordServer(apiURL string) {
	s.initRecordServerOnPort(apiURL, 61417)
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
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone request body
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		proxyRequest.Header = newCallerProxyRequestHeader
		response, err := http.DefaultClient.Do(proxyRequest)

		// Clone response body
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		if response.Header.Get("Content-Length") != "" {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", utf8.RuneCountInString(newCallerResponseBody)))
		}
		w.WriteHeader(response.StatusCode)
		w.Write([]byte(newCallerResponseBody))
	}
}

// checkMarkdownExists ...
func (s *Impl) checkMarkdownExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
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
	workingPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/template.tmpl", workingPath))
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New("template").Parse(string(content))
	if err != nil {
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

// SetCallerRequestHeadersRemoval ...
func (s *Impl) SetCallerRequestHeadersRemoval(headers []string) {
	s.callerRequestHeadersRemoval = headers
}

// SetCallerRequestHeaderReplacements ...
func (s *Impl) SetCallerRequestHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.callerRequestHeaderReplacements = replaceRegex
}

// SetCallerRequestBodyReplacement ...
func (s *Impl) SetCallerRequestBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.callerRequestBodyReplacement = replaceRegex
}

// SetRecordRequestHeadersRemoval ...
func (s *Impl) SetRecordRequestHeadersRemoval(headers []string) {
	s.recordRequestHeadersRemoval = headers
}

// SetRecordRequestHeaderReplacements ...
func (s *Impl) SetRecordRequestHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.recordRequestHeaderReplacements = replaceRegex
}

// SetRecordRequestBodyReplacement ...
func (s *Impl) SetRecordRequestBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.recordRequestBodyReplacement = replaceRegex
}

// SetCallerResponseHeadersRemoval ...
func (s *Impl) SetCallerResponseHeadersRemoval(headers []string) {
	s.callerResponseHeadersRemoval = headers
}

// SetCallerResponseHeaderReplacements ...
func (s *Impl) SetCallerResponseHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.callerResponseHeaderReplacements = replaceRegex
}

// SetCallerResponseBodyReplacement ...
func (s *Impl) SetCallerResponseBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.callerResponseBodyReplacement = replaceRegex
}

// SetRecordResponseHeadersRemoval ...
func (s *Impl) SetRecordResponseHeadersRemoval(headers []string) {
	s.recordResponseHeadersRemoval = headers
}

// SetRecordResponseHeaderReplacements ...
func (s *Impl) SetRecordResponseHeaderReplacements(replaceRegex map[*regexp.Regexp]string) {
	s.recordResponseHeaderReplacements = replaceRegex
}

// SetRecordResponseBodyReplacement ...
func (s *Impl) SetRecordResponseBodyReplacement(replaceRegex map[*regexp.Regexp]string) {
	s.recordResponseBodyReplacement = replaceRegex
}
