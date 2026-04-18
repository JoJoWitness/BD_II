package main

import (
	"log"
	"net/http"

	"demo/config"
	"demo/routes"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	mux := http.NewServeMux()
	routes.Register(mux)

	addr := ":8080"
	log.Printf("API escuchando en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, cors(mux)); err != nil {
		log.Fatal(err)
	}
}
