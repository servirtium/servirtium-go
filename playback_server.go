package servirtium

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// IServerPlayback ...
type IServerPlayback interface {
	Start()
	Close()
}

// ServerPlaybackTestImpl ...
type ServerPlaybackTestImpl struct {
	serverPlayback *httptest.Server
	handler        func(w http.ResponseWriter, r *http.Request)
}

// NewServerPlaybackTest ...
func NewServerPlaybackTest(handler func(w http.ResponseWriter, r *http.Request)) *ServerPlaybackTestImpl {
	return &ServerPlaybackTestImpl{
		handler: handler,
	}
}

// initServer ...
func (s *ServerPlaybackTestImpl) initServer() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.handler)
	srv := httptest.NewServer(r)
	s.serverPlayback = srv
}

// Start ...
func (s *ServerPlaybackTestImpl) Start() {
	s.initServer()
}

// Close ...
func (s *ServerPlaybackTestImpl) Close() {
	s.serverPlayback.Close()
}

// ServerPlaybackImpl ...
type ServerPlaybackImpl struct {
	serverPlayback *http.Server
	handler        func(w http.ResponseWriter, r *http.Request)
	port           int
}

// NewServerPlayback ...
func NewServerPlayback(handler func(w http.ResponseWriter, r *http.Request), port int) *ServerPlaybackImpl {
	return &ServerPlaybackImpl{
		handler: handler,
		port:    port,
	}
}

// initServer ...
func (s *ServerPlaybackImpl) initServer() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.handler)
	srv := &http.Server{
		Handler: disableCors(r),
		Addr:    "127.0.0.1:" + strconv.Itoa(s.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.serverPlayback = srv
}

// Start ...
func (s *ServerPlaybackImpl) Start() {
	s.initServer()
	log.Fatal(s.serverPlayback.ListenAndServe())
}

// Close ...
func (s *ServerPlaybackImpl) Close() {
	s.serverPlayback.Close()
}
