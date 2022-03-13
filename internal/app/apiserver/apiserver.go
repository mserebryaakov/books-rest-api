package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mserebryaakov/books-rest-api/internal/app/store"
)

//API server
type APIServer struct {
	config *Config
	router *mux.Router
	Store  *store.Store
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

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	err := st.Open()
	if err != nil {
		return nil
	}

	s.Store = st
	return nil
}

//Start ...
func (s *APIServer) Start() error {
	s.config.Log.Println("Start server")
	s.configureRouter()

	err := s.configureStore()
	if err != nil {
		return nil
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
