package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func ActivityRouter(r *mux.Router) {
	router := r.PathPrefix("/activity-groups").Subrouter()
	router.HandleFunc("", controllers.GetActivity).Methods("GET")
	router.HandleFunc("", controllers.CrateActivity).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetOne).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateActivity).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteActivity).Methods("DELETE")
}
