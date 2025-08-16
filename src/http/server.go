package http

import (
	"fmt"
	"net/http"

	"github.com/PopSquad/BalloonField/src/util"
)

type Server struct {
	Address string
	Logger  util.Logger
}

func NewHTTPServer(address string) *Server {
	return &Server{
		Address: address,
		Logger: util.Logger{
			Tag: "HTTPServer",
		},
	}
}

func (s *Server) httpDefaultHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Log("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("svr is running\n"))
}

func (s *Server) userScoreListHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Log("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO\n"))
}

func (s *Server) userFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Log("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO2\n"))
}

func (s *Server) Start() {
	http.HandleFunc("/", s.httpDefaultHandler)
	http.HandleFunc("/Idea/GameUserScoresList.aspx", s.userScoreListHandler)
	http.HandleFunc("/Idea/autologin.aspx", s.userFeedbackHandler)

	s.Logger.Log("running at %s", s.Address)
	if err := http.ListenAndServe(s.Address, nil); err != nil {
		panic(fmt.Errorf("HTTP server failed to start: %v", err))
	}
}
