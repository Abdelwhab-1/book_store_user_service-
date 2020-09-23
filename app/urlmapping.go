package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func mapper(mux *mux.Router) {
	get := mux.Methods(http.MethodGet).Subrouter()
	post := mux.Methods(http.MethodPost).Subrouter()
	put := mux.Methods(http.MethodPut).Suprouter()
	del := mux.Methods(http.MethodDelete).Suprouter()

	get.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser)
	post.HandleFunc("/user", handlers.CreateUser)
	put.HandleFunc("/user/{id:[0-9]+}", handlers.EditUser)
	del.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUser)

}
