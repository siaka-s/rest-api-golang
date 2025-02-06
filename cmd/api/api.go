package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {

	return &ApiServer{
		addr: addr,
	}

}

func (s *ApiServer) Run() error {

	// creer le routeur

	router := mux.NewRouter()

	// Port d'ecoute

	log.Println("ecoute sur le port", s.addr)

	return http.ListenAndServe(s.addr, router)

}
