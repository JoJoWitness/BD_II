package routes

import (
	"net/http"

	"demo/controllers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/login", controllers.Login)
	mux.HandleFunc("/api/logs", controllers.Logs)
	mux.HandleFunc("/api/logout", controllers.Logout)
}
