package api

import (
	"database/sql"
	"enfermeria_go/service/home"
	"enfermeria_go/service/paciente"
	"enfermeria_go/service/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	adress string
	db     *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		adress: addr,
		db:     db,
	}

}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	home.RegisterRoutes(router)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	pacienteStore := paciente.NewStore(s.db)
	pacienteHandler := paciente.NewHandlerPaciente(pacienteStore)
	pacienteHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.adress)

	return http.ListenAndServe(s.adress, router)

}
