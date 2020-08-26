package servirtium

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// IServerRecord ...
type IServerRecord interface {
	Start()
	Close()
}

// ServerRecordTestImpl ...
type ServerRecordTestImpl struct {
	serverRecord *httptest.Server
	handler      func(w http.ResponseWriter, r *http.Request)
}

// NewServerRecordTest ...
func NewServerRecordTest(handler func(w http.ResponseWriter, r *http.Request)) *ServerRecordTestImpl {
	return &ServerRecordTestImpl{
		handler: handler,
	}
}

// initServer ...
func (s *ServerRecordTestImpl) initServer() {
	l, err := net.Listen("tcp", "127.0.0.1:61417")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.handler)
	ts := httptest.NewUnstartedServer(r)

	// NewUnstartedServer creates a listener. Close that listener and replace
	// with the one we created.
	ts.Listener.Close()
	ts.Listener = l
	s.serverRecord = ts
}

// Start ...
func (s *ServerRecordTestImpl) Start() {
	s.initServer()
	s.serverRecord.Start()
}

// Close ...
func (s *ServerRecordTestImpl) Close() {
	s.serverRecord.Close()
}

// ServerRecordImpl ...
type ServerRecordImpl struct {
	serverRecord *http.Server
	handler      func(w http.ResponseWriter, r *http.Request)
	port         int
}

// NewServerRecord ...
func NewServerRecord(handler func(w http.ResponseWriter, r *http.Request), port int) *ServerRecordImpl {
	return &ServerRecordImpl{
		handler: handler,
		port:    port,
	}
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PATCH, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		// I added this for another handler of mine,
		// but I do not think this is necessary for GraphQL's handler
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// initServer ...
func (s *ServerRecordImpl) initServer() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(s.handler)
	srv := &http.Server{
		Handler: disableCors(r),
		Addr:    "127.0.0.1:" + strconv.Itoa(s.port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.serverRecord = srv
}

// Start ...
func (s *ServerRecordImpl) Start() {
	s.initServer()
	log.Fatal(s.serverRecord.ListenAndServe())
}

// Close ...
func (s *ServerRecordImpl) Close() {
	s.serverRecord.Close()
}
