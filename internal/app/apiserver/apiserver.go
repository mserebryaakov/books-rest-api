package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//API server
type APIServer struct {
	config *Config
	router *mux.Router
}

//New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello rw")
	}
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

//Start ...
func (s *APIServer) Start() error {
	s.config.Log.Println("Start server")
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
