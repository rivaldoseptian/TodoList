package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/routes"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	config.ConectDB()

	r := mux.NewRouter()
	routes.ActivityRouter(r)
	routes.TodoRouter(r)

	log.Println("Server Running On Port:", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
