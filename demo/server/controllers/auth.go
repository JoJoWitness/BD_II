package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"demo/models"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type logoutRequest struct {
	UserID int `json:"user_id"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if req.Username == "" || req.Password == "" {
		writeError(w, http.StatusBadRequest, "username y password son requeridos")
		return
	}

	user, err := models.FindUserByUsername(req.Username)
	switch {
	case errors.Is(err, models.ErrUserNotFound):
		user, err = models.CreateUser(req.Username, req.Password)
		if err != nil {
			log.Printf("create user: %v", err)
			writeError(w, http.StatusInternalServerError, "error creando usuario")
			return
		}
	case err != nil:
		log.Printf("find user: %v", err)
		writeError(w, http.StatusInternalServerError, "error consultando usuario")
		return
	default:
		if user.Password != req.Password {
			writeError(w, http.StatusUnauthorized, "credenciales inválidas")
			return
		}
	}

	if err := models.InsertSessionLog(user.ID, models.EventLogin); err != nil {
		log.Printf("insert login log: %v", err)
		writeError(w, http.StatusInternalServerError, "error registrando login")
		return
	}

	writeJSON(w, http.StatusOK, loginResponse{UserID: user.ID, Username: user.Username})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	var req logoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if req.UserID <= 0 {
		writeError(w, http.StatusBadRequest, "user_id inválido")
		return
	}

	if err := models.InsertSessionLog(req.UserID, models.EventLogout); err != nil {
		log.Printf("insert logout log: %v", err)
		writeError(w, http.StatusInternalServerError, "error registrando logout")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
