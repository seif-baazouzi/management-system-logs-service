package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	LogID       uuid.UUID `json:"logID"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Value       uint      `json:"value"`
	Date        time.Time `json:"date"`
	UserID      uuid.UUID `json:"userID"`
	WorkspaceID uuid.UUID `json:"workspaceID"`
	CreatedAt   time.Time `json:"createdAt"`
}

type LogBody struct {
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Value       uint      `json:"value"`
	Date        time.Time `json:"date"`
}
