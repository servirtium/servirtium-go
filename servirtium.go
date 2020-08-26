package servirtium

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// Impl ...
type Impl struct {
	ServerPlayback            IServerPlayback
	ServerRecord              IServerRecord
	requestSequence           int64
	content                   string
	requestHeadersNeedDelete  []string
	requestHeadersNeedReplace map[string]string
	requestHeadersNeedMask    map[string]string
	requestBodyNeedMask       map[*regexp.Regexp]string
	responseHeadersNeedDelete []string
	responseHeaderNeedReplace map[string]string
	responseHeadersNeedMask   map[string]string
	responseBodyNeedMask      map[*regexp.Regexp]string
}

// NewServirtium ...
func NewServirtium() *Impl {
	return &Impl{
		requestSequence:           0,
		content:                   "",
		requestHeadersNeedDelete:  []string{},
		responseHeadersNeedDelete: []string{},
	}
}

// StartPlayback ...
func (s *Impl) StartPlayback() {
	s.ServerPlayback.Start()
}

// EndPlayback ...
func (s *Impl) EndPlayback() {
	s.ServerPlayback.Close()
}

func (s *Impl) getInteraction(data string) string {
	interactions := strings.Split(fmt.Sprintf("\n%s", data), "\n## Interaction ")
	return interactions[s.requestSequence+1]
}

func (s *Impl) parseHeaders(content string) map[string]string {
	var (
		headers = map[string]string{}
	)
	newLineReplaceByCommaContent := strings.ReplaceAll(content, "\n", ",")
	headersContents := strings.Split(newLineReplaceByCommaContent, ",")
	for _, v := range headersContents {
		replaceByComma := strings.ReplaceAll(v, ": ", ",")
		keysAndValues := strings.Split(replaceByComma, ": ")
		keyIndex := 0
		valueIndex := 1
		isValidHeaders := len(keysAndValues) > 1
		if isValidHeaders {
			headers[keysAndValues[keyIndex]] = keysAndValues[valueIndex]
		}
	}
	return headers
}

func (s *Impl) getPlaybackResponse(data string) (string, map[string]string) {
	var (
		headers      map[string]string
		responseBody string
		sections     = strings.Split(data, "\n### ")
	)
	for _, v := range sections {
		if strings.HasPrefix(v, "Response headers recorded for playback") {
			headerContent := strings.Split(v, "```")[1]
			headers = s.parseHeaders(headerContent)
		}
		if strings.HasPrefix(v, "Response body recorded for playback") {
			responseBody = strings.Split(v, "```")[1]
			headers["content-length"] = fmt.Sprintf("%d", len(responseBody))
		}
	}
	return responseBody, headers
}

// AnualAvgHandlerPlayback ...
func (s *Impl) AnualAvgHandlerPlayback(recordFileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		workingPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName))
		interaction := s.getInteraction(string(data))
		body, headers := s.getPlaybackResponse(interaction)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
		return
	}
}

// StartRecord ...
func (s *Impl) StartRecord() {
	s.ServerRecord.Start()
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
	s.ServerRecord.Close()
}

type recordData struct {
	RecordSequence      int64
	RequestMethod       string
	RequestURLPath      string
	RequestHeader       map[string][]string
	RequestBody         string
	ResponseHeader      map[string][]string
	ResponseBody        string
	ResponseStatus      string
	ResponseContentType string
}

func removeHeader(header http.Header, deleteItems []string) http.Header {
	for _, v := range deleteItems {
		header.Del(v)
	}
	return header
}

func replaceHeader(header http.Header, replaceItems map[string]string) http.Header {
	for k, v := range replaceItems {
		if _, isFound := header[k]; isFound {
			header.Set(k, v)
		}
	}
	return header
}

func maskBody(content string, maskItems map[*regexp.Regexp]string) string {
	newContent := content
	for k, v := range maskItems {
		newContent = k.ReplaceAllString(content, v)
	}
	return newContent
}

// ManInTheMiddleHandler ...
func (s *Impl) ManInTheMiddleHandler(apiURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone Request Body
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		url := fmt.Sprintf("%s%s", apiURL, r.RequestURI)
		proxyRequest, err := http.NewRequest(r.Method, url, bytes.NewReader(requestBody))
		// Modify request headers
		proxyRequest.Header = replaceHeader(r.Header, s.requestHeadersNeedReplace)
		proxyRequest.Header = removeHeader(proxyRequest.Header, s.requestHeadersNeedDelete)
		response, err := http.DefaultClient.Do(proxyRequest)
		// Clone response
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Modify Request Header before write
		newRequestHeader := replaceHeader(proxyRequest.Header, s.requestHeadersNeedMask)
		maskedRequestBody := maskBody(string(requestBody), s.requestBodyNeedMask)

		// Modify Response Header before write
		newResponseHeader := replaceHeader(response.Header, s.responseHeaderNeedReplace)
		removedResponseHeader := removeHeader(newResponseHeader, s.responseHeadersNeedDelete)
		maskedResponseHeader := replaceHeader(removedResponseHeader, s.responseHeadersNeedMask)
		maskedResponseBody := maskBody(string(responseBody), s.responseBodyNeedMask)

		s.record(recordData{
			RequestURLPath:      r.URL.Path,
			RequestMethod:       r.Method,
			RequestHeader:       newRequestHeader,
			RequestBody:         maskedRequestBody,
			ResponseHeader:      maskedResponseHeader,
			ResponseBody:        maskedResponseBody,
			ResponseContentType: response.Header.Get("Content-Type"),
			ResponseStatus:      response.Status,
		})
		response.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
		defer func() {
			_ = response.Body.Close()
		}()
		w.Write(responseBody)
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
	if s.requestSequence == 0 {
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
		RecordSequence:      s.requestSequence,
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
	s.requestSequence = s.requestSequence + 1
}

// CheckMarkdownIsDifferentToPreviousRecording ...
func (s *Impl) CheckMarkdownIsDifferentToPreviousRecording(recordFileName string) bool {
	workingPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName)
	fileContent, err := ioutil.ReadFile(filePath)
	isCanReadFile := err == nil
	if !isCanReadFile {
		return true
	}
	newContent := maskBody(s.content, s.requestBodyNeedMask)
	newContent = maskBody(newContent, s.requestBodyNeedMask)

	return newContent == string(fileContent)
}

// DeleteRequestHeaders ...
func (s *Impl) DeleteRequestHeaders(headers []string) {
	for _, v := range headers {
		s.requestHeadersNeedDelete = append(s.requestHeadersNeedDelete, v)
	}
}

// MaskRequestHeaders ...
func (s *Impl) MaskRequestHeaders(headers map[string]string) {
	s.requestHeadersNeedMask = headers
}

// ReplaceRequestHeaders ...
func (s *Impl) ReplaceRequestHeaders(headers map[string]string) {
	s.requestHeadersNeedReplace = headers
}

// MaskRequestBody ...
func (s *Impl) MaskRequestBody(replaceRegex map[*regexp.Regexp]string) {
	s.requestBodyNeedMask = replaceRegex
}

// DeleteResponseHeaders ...
func (s *Impl) DeleteResponseHeaders(headers []string) {
	for _, v := range headers {
		s.responseHeadersNeedDelete = append(s.responseHeadersNeedDelete, v)
	}
}

// MaskResponseHeaders ...
func (s *Impl) MaskResponseHeaders(headers map[string]string) {
	s.responseHeadersNeedMask = headers
}

// ReplaceResponseHeaders ...
func (s *Impl) ReplaceResponseHeaders(headers map[string]string) {
	s.responseHeaderNeedReplace = headers
}

// MaskResponseBody ...
func (s *Impl) MaskResponseBody(replaceRegex map[*regexp.Regexp]string) {
	s.responseBodyNeedMask = replaceRegex
}
