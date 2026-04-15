package main

import (
	"log"
	"net/http"

	"demo/config"
	"demo/controllers"
	"demo/routes"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	mux := http.NewServeMux()
	routes.Register(mux)

	addr := ":8080"
	log.Printf("API escuchando en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, controllers.CORS(mux)); err != nil {
		log.Fatal(err)
	}
}
