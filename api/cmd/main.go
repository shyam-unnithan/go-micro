package main

import (
	"log"
	"net/http"

	"github.com/shyam-unnithan/eduwiz/api/pkg/router"
	_ "github.com/shyam-unnithan/eduwiz/api/pkg/bootstrapper"
)

func main() {

	r := router.InitRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
