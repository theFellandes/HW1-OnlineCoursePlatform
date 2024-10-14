package web

import (
	"github.com/gorilla/mux"
	"internal/config"
	"internal/services"
	"log"
	"net/http"
)

const (
	STUDENTS = "/students"
	INSTRUCTORS = "/instructors"
	COURSES = "/courses"
	ENROLLMENT = "/enrollments"
)

func Routes(conf *config.Conf) {
	route := mux.NewRouter()

	// Students endpoints, Add Student, Get Student
	route.HandleFunc(STUDENTS, services.GetUser).Methods(http.MethodGet)
	route.HandleFunc(STUDENTS, services.AddUser).Methods(http.MethodPost)

	route.HandleFunc(INSTRUCTORS, services.GetUser).Methods(http.MethodGet)
	route.HandleFunc(INSTRUCTORS, services.AddUser).Methods(http.MethodPost)

	route.HandleFunc(COURSES, services.GetUser).Methods(http.MethodGet)
	route.HandleFunc(COURSES, services.AddUser).Methods(http.MethodPost)

	route.HandleFunc(ENROLLMENT, services.GetUser).Methods(http.MethodGet)
	route.HandleFunc(ENROLLMENT, services.AddUser).Methods(http.MethodPost)


	// Start the server and handle any potential errors
	server := &http.Server{
		Addr:    conf.App.Port, // Address to listen on
		Handler: route,         // The router to handle incoming requests
	}

	// Log fatal error if http.ListenAndServe fails
	log.Fatal(server.ListenAndServe())
}
