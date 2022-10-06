package models

import (
	"database/sql"
	"logs-service/src/db"

	"github.com/google/uuid"
)

func GetLogs(workspaceID string, startingDate string, endingDate string) ([]Log, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var logs []Log

	var rows *sql.Rows
	var err error

	rows, err = conn.Query(
		`SELECT logID, label, description, value, date, userID, workspaceID, createdAt FROM logs
			WHERE workspaceID = $1 AND date >= $2 AND date <= $3
			ORDER BY createdAt`,
		workspaceID,
		startingDate,
		endingDate,
	)

	if err != nil {
		return logs, err
	}

	defer rows.Close()

	for rows.Next() {
		var log Log
		rows.Scan(&log.LogID, &log.Label, &log.Description, &log.Value, &log.Date, &log.UserID, &log.WorkspaceID, &log.CreatedAt)

		logs = append(logs, log)
	}

	return logs, nil
}

func CreateLog(log LogBody, workspaceID string, userID string) (string, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	logID := uuid.New()

	_, err := conn.Exec(
		"INSERT INTO logs (logID, label, description, value, date, workspaceID, userID) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		logID,
		log.Label,
		log.Description,
		log.Value,
		log.Date,
		workspaceID,
		userID,
	)

	if err != nil {
		return "", err
	}

	return logID.String(), nil
}

func UpdateLog(log LogBody, logID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"UPDATE logs SET label = $1, description = $2, value = $3, date = $4 WHERE logID = $5 AND userID = $6",
		log.Label,
		log.Description,
		log.Value,
		log.Date,
		logID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}
