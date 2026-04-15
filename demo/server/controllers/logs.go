package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"demo/models"
)

type logResponse struct {
	ID        int    `json:"id"`
	EventType string `json:"event_type"`
	EventTime string `json:"event_time"`
}

func Logs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil || userID <= 0 {
		writeError(w, http.StatusBadRequest, "user_id inválido")
		return
	}

	logs, err := models.FindSessionLogsByUserID(userID)
	if err != nil {
		log.Printf("find logs: %v", err)
		writeError(w, http.StatusInternalServerError, "error consultando logs")
		return
	}

	out := make([]logResponse, len(logs))
	for i, l := range logs {
		out[i] = logResponse{
			ID:        l.ID,
			EventType: l.EventType,
			EventTime: l.EventTime.Format(time.RFC3339),
		}
	}
	writeJSON(w, http.StatusOK, out)
}
