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
	"github.com/kr/pretty"
)

// IServirtium ...
type IServirtium interface {
	StartRecord(apiURL string)
	WriteRecord(recordFileName string)
	CheckMarkdownIsDifferentToPreviousRecording(recordFileName string) bool
	EndRecord()
	StartPlayback(recordFileName string)
	EndPlayback(recordFileName string)
}

// Impl ...
type Impl struct {
	ServerPlayback  *httptest.Server
	ServerRecord    *httptest.Server
	RequestSequence int64
	Content         string
}

// NewServirtium ...
func NewServirtium() *Impl {
	return &Impl{
		RequestSequence: 0,
		Content:         "",
	}
}

// StartPlayback ...
func (s *Impl) StartPlayback(recordFileName string) {
	s.initServerPlayback(recordFileName)
}

// EndPlayback ...
func (s *Impl) EndPlayback(recordFileName string) {
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
		w.Header().Set("Content-Type", "application/xml")
		_, _ = w.Write(data)
		return
	}
}

// StartRecord ...
func (s *Impl) StartRecord(apiURL string) {
	s.initRecordServer()
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
		pretty.Println(filePath)
		os.Create(filePath)
	}
	err = ioutil.WriteFile(filePath, []byte(s.Content), os.ModePerm)
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
	r.PathPrefix("/").HandlerFunc(s.manInTheMiddleHandler(apiURL string))
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

		// We may want to filter some headers, otherwise we could just use a shallow copy
		// proxyReq.Header = r.Header
		proxyReq.Header = make(http.Header)
		for h, val := range r.Header {
			proxyReq.Header[h] = val
		}
		proxyReq.Header.Set("User-Agent", "Servirtium-Testing")

		resp, err := http.DefaultClient.Do(proxyReq)
		// Clone resp
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newRespHeader := resp.Header
		newRespHeader.Del("Set-Cookie")
		newRespHeader.Del("Date")
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
		s.record(recordData{
			RequestURLPath:      r.URL.Path,
			RequestMethod:       r.Method,
			RequestHeader:       r.Header,
			RequestBody:         string(reqBody),
			ResponseHeader:      newRespHeader,
			ResponseBody:        string(respBody),
			ResponseContentType: resp.Header.Get("Content-Type"),
			ResponseStatus:      resp.Status,
		})

		defer func() {
			_ = resp.Body.Close()
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
	if s.RequestSequence == 0 {
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
		RecordSequence:      s.RequestSequence,
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
	finalContent := s.appendContentInFile(s.Content, string(newContent))
	s.Content = finalContent
	s.RequestSequence = s.RequestSequence + 1
}

// CheckMarkdownIsDifferentToPreviousRecording ...
func (s *Impl) CheckMarkdownIsDifferentToPreviousRecording(recordFileName string) bool {
	workingPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/mock/%s.md", workingPath, recordFileName)
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return s.Content == string(fileContent)
}
