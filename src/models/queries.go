package models

import (
	"database/sql"
	"logs-service/src/db"
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