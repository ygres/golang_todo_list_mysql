package apiserver

import (
	"app/internal/app/handler"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initializes the app with predefined configuration
func (s *Server) Initialize(DatabaseURL string) {

	db, err := sql.Open("mysql", DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s.DB = db
	s.Router = mux.NewRouter()
	s.setRouters()
}

// Configure the router
func (s *Server) setRouters() {
	s.Router.HandleFunc("/projects", s.handleRequest(handler.GetAllProject)).Methods("GET")
	s.Router.HandleFunc("/projects", s.handleRequest(handler.CreateProject)).Methods("POST")
	s.Router.HandleFunc("/projects/{title}", s.handleRequest(handler.GetProject)).Methods("GET")
	s.Router.HandleFunc("/projects/{title}", s.handleRequest(handler.UpdateProject)).Methods("PUT")
	s.Router.HandleFunc("/projects/{title}", s.handleRequest(handler.DeleteProject)).Methods("DELETE")

	s.Router.HandleFunc("/tasks", s.handleRequest(handler.GetAllTasks)).Methods("GET")
	s.Router.HandleFunc("/projects/{title}/task", s.handleRequest(handler.GetProjectTasks)).Methods("GET")
	s.Router.HandleFunc("/projects/{title}/task", s.handleRequest(handler.CreateTask)).Methods("POST")

	s.Router.HandleFunc("/task/{id:[0-9]+}", s.handleRequest(handler.UpdateTask)).Methods("PUT")
	s.Router.HandleFunc("/task/{id:[0-9]+}", s.handleRequest(handler.DeleteTask)).Methods("DELETE")
}

// Run the app on it's router
func (s *Server) Start(host string) {
	log.Fatal(http.ListenAndServe(host, s.Router))
}

type RequestHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler(s.DB, w, r)
	}
}
