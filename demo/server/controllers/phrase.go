package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"demo/config"
)

type phraseRow struct {
	ID   int    `json:"id"`
	Text string `json:"phrase"`
}

type phraseInput struct {
	Phrase string `json:"phrase"`
}

func ListPhrases(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id, text FROM phrases ORDER BY id`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	list := []phraseRow{}
	for rows.Next() {
		var row phraseRow
		if err := rows.Scan(&row.ID, &row.Text); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		list = append(list, row)
	}
	writeJSON(w, http.StatusOK, list)
}

func GetPhrase(w http.ResponseWriter, r *http.Request) {
	var row phraseRow
	err := config.DB.QueryRow(
		`SELECT id, text FROM phrases ORDER BY RANDOM() LIMIT 1`,
	).Scan(&row.ID, &row.Text)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "no hay frases")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, row)
}

func AddPhrase(w http.ResponseWriter, r *http.Request) {
	var in phraseInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	text := strings.TrimSpace(in.Phrase)
	if text == "" {
		writeError(w, http.StatusBadRequest, "phrase requerido")
		return
	}
	var row phraseRow
	err := config.DB.QueryRow(
		`INSERT INTO phrases (text) VALUES ($1) RETURNING id, text`,
		text,
	).Scan(&row.ID, &row.Text)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, row)
}

func DeletePhrase(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id inválido")
		return
	}
	res, err := config.DB.Exec(`DELETE FROM phrases WHERE id = $1`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		writeError(w, http.StatusNotFound, "frase no encontrada")
		return
	}
	writeJSON(w, http.StatusOK, map[string]int{"deleted": id})
}
