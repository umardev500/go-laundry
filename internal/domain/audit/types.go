package audit

import "github.com/google/uuid"

type Action string

const (
	CreateAction Action = "create"
	UpdateAction Action = "update"
	DeleteAction Action = "delete"
)

type Create struct {
	TableName  string
	RecordID   uuid.UUID
	Action     Action
	ModifiedBy uuid.UUID
	OldData    map[string]any
	NewData    map[string]any
}
