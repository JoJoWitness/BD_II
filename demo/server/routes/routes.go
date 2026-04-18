package routes

import (
	"net/http"

	"demo/controllers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /phrase", controllers.GetPhrase)
	mux.HandleFunc("GET /phrases", controllers.ListPhrases)
	mux.HandleFunc("POST /phrase", controllers.AddPhrase)
	mux.HandleFunc("DELETE /phrase", controllers.DeletePhrase)
}
