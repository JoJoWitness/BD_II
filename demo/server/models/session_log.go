package models

import (
	"time"

	"demo/config"
)

const (
	EventLogin  = "LOGIN"
	EventLogout = "LOGOUT"
)

type SessionLog struct {
	ID        int       `json:"id"`
	EventType string    `json:"event_type"`
	EventTime time.Time `json:"event_time"`
}

func InsertSessionLog(userID int, eventType string) error {
	_, err := config.DB.Exec(
		`INSERT INTO session_logs (user_id, event_type) VALUES ($1, $2)`,
		userID, eventType,
	)
	return err
}

func FindSessionLogsByUserID(userID int) ([]SessionLog, error) {
	rows, err := config.DB.Query(
		`SELECT id, event_type, event_time FROM session_logs WHERE user_id=$1 ORDER BY event_time DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := make([]SessionLog, 0)
	for rows.Next() {
		var l SessionLog
		if err := rows.Scan(&l.ID, &l.EventType, &l.EventTime); err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, rows.Err()
}
