package main

import (
	"fmt"
	"net/http"
	"server/config"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	config.ConectDB()

	r := mux.NewRouter()

	log.Println("Server Running On Port:", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
