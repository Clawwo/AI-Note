package dto

import (
	"time"

	"github.com/google/uuid"
)

// Request dan Response DTO untuk Notebook
type CreateNotebookRequest struct {
	Name string `json:"name" validate:"required"`
	ParentId *uuid.UUID `json:"parent_id"`
}

type CreateNotebookResponse struct {
	Id uuid.UUID `json:"id"`
}

// DTO untuk menampilkan detail notebook
type ShowNotebookResponse struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	ParentId *uuid.UUID `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DTO untuk memperbarui notebook
type UpdateNotebookRequest struct {
	Id uuid.UUID 
	Name string `json:"name" validate:"required"`
}

type UpdateNotebookResponse struct {
	Id uuid.UUID `json:"id"`
}

// DTO untuk memindahkan notebook
type MoveNotebookRequest struct {
	Id uuid.UUID 
	ParentId *uuid.UUID `json:"parent_id"`
}

type MoveNotebookResponse struct {
	Id uuid.UUID `json:"id"`
}