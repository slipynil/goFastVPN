package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	httpHandlers handlers
}

func New(awg awg) *server {
	return &server{
		httpHandlers: handlers{awg},
	}
}

func (s *server) Start(endpoint string) {
	r := mux.NewRouter()
	r.HandleFunc("/peers", s.httpHandlers.DeletePeer).Methods("DELETE")
	r.HandleFunc("/peers", s.httpHandlers.AddPeer).Methods("POST")

	http.ListenAndServe(endpoint, r)
}
