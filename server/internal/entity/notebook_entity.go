package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notebook struct {
	Id uuid.UUID
	Name string
	Parent_id *uuid.UUID
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
	IsDeleted bool
}
