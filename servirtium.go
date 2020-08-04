package servirtium

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

// Impl ...
type Impl struct {
	ServerPlayback            *httptest.Server
	ServerRecord              *httptest.Server
	requestSequence           int64
	content                   string
	requestHeadersNeedDel     []string
	requestHeadersNeedMask    map[string]string
	requestHeadersNeedReplace map[string]string
	responseHeadersNeedDel    []string
	responseHeadersNeedMask   map[string]string
}

// NewServirtium ...
func NewServirtium() *Impl {
	return &Impl{
		requestSequence:        0,
		content:                "",
		requestHeadersNeedDel:  []string{},
		responseHeadersNeedDel: []string{},
	}
}

// StartPlayback ...
func (s *Impl) StartPlayback(recordFileName string) {
	s.initServerPlayback(recordFileName)
}

// EndPlayback ...
func (s *Impl) EndPlayback() {
	s.ServerPlayback.Close()
}

func (s *Impl) initServerPlayback(recordFileName string) {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.anualAvgHandlerPlayback(recordFileName))
	srv := httptest.NewServer(r)
	s.ServerPlayback = srv
}

func (s *Impl) anualAvgHandlerPlayback(recordFileName string) func(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
		return
	}
}

// StartRecord ...
func (s *Impl) StartRecord(apiURL string) {
	s.initRecordServer(apiURL)
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

func (s *Impl) initRecordServer(apiURL string) {
	l, err := net.Listen("tcp", "127.0.0.1:61417")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.manInTheMiddleHandler(apiURL))
	ts := httptest.NewUnstartedServer(r)

	// NewUnstartedServer creates a listener. Close that listener and replace
	// with the one we created.
	ts.Listener.Close()
	ts.Listener = l
	s.ServerRecord = ts
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

func maskHeader(header http.Header, maskItems map[string]string) http.Header {
	for k, v := range maskItems {
		if _, isFound := header[k]; isFound {
			header.Set(k, v)
		}
	}
	return header
}

func replaceHeader(header http.Header, maskItems map[string]string) http.Header {
	for k, v := range maskItems {
		if _, isFound := header[k]; isFound {
			header.Set(k, v)
		}
	}
	return header
}

func (s *Impl) manInTheMiddleHandler(apiURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone Request Body
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
		url := fmt.Sprintf("%s%s", apiURL, r.RequestURI)
		proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(reqBody))
		// Modify request headers
		proxyReq.Header = replaceHeader(r.Header, s.requestHeadersNeedReplace)
		proxyReq.Header = removeHeader(proxyReq.Header, s.requestHeadersNeedDel)
		response, err := http.DefaultClient.Do(proxyReq)
		// Clone response
		respBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Modify Request Header before write
		newReqHeader := maskHeader(proxyReq.Header, s.requestHeadersNeedMask)
		// Modify Response Header before write
		newRespHeader := removeHeader(response.Header, s.responseHeadersNeedDel)
		newRespHeader = maskHeader(newRespHeader, s.responseHeadersNeedMask)

		response.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
		s.record(recordData{
			RequestURLPath:      r.URL.Path,
			RequestMethod:       r.Method,
			RequestHeader:       newReqHeader,
			RequestBody:         string(reqBody),
			ResponseHeader:      newRespHeader,
			ResponseBody:        string(respBody),
			ResponseContentType: response.Header.Get("Content-Type"),
			ResponseStatus:      response.Status,
		})

		defer func() {
			_ = response.Body.Close()
		}()
		w.Write(respBody)
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
	return s.content == string(fileContent)
}

// DelRespHeaders ...
func (s *Impl) DelRespHeaders(headers []string) {
	for _, v := range headers {
		s.responseHeadersNeedDel = append(s.responseHeadersNeedDel, v)
	}
}

// DelRequestHeaders ...
func (s *Impl) DelRequestHeaders(headers []string) {
	for _, v := range headers {
		s.requestHeadersNeedDel = append(s.requestHeadersNeedDel, v)
	}
}

// MaskRequestHeaders ...
func (s *Impl) MaskRequestHeaders(headers map[string]string) {
	s.requestHeadersNeedMask = headers
}

// MaskResponseHeaders ...
func (s *Impl) MaskResponseHeaders(headers map[string]string) {
	s.responseHeadersNeedMask = headers
}

// ReplaceRequestHeaders ...
func (s *Impl) ReplaceRequestHeaders(headers map[string]string) {
	s.requestHeadersNeedReplace = headers
}
