package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func TodoRouter(r *mux.Router) {
	router := r.PathPrefix("/todo-items").Subrouter()
	router.HandleFunc("", controllers.GetAllTodo).Methods("GET")
	router.HandleFunc("", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetOneTodo).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteTodo).Methods("DELETE")
}
