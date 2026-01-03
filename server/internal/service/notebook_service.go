package service

import (
	"ai-notetaking-be/internal/dto"
	"ai-notetaking-be/internal/entity"
	"ai-notetaking-be/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type INotebookService interface {
	Create(ctx context.Context, req *dto.CreateNotebookRequest) (*dto.CreateNotebookResponse, error)
	Show(ctx context.Context, id uuid.UUID) (*dto.ShowNotebookResponse, error)
	Update(ctx context.Context, req *dto.UpdateNotebookRequest) (*dto.UpdateNotebookResponse, error)
}

type notebookService struct {
	notebookRepository repository.INotebookRepository
}

func NewNotebookService(notebookRepository repository.INotebookRepository) INotebookService {
	return &notebookService{
		notebookRepository: notebookRepository,
	}
}

// Membuat notebook baru
func (c *notebookService) Create(ctx context.Context, req *dto.CreateNotebookRequest) (*dto.CreateNotebookResponse, error) {

	notebook := entity.Notebook{
		Id: uuid.New(),
		Name: req.Name,
		Parent_id: req.ParentId,
		Created_at: time.Now(),
	}



	err := c.notebookRepository.Create(ctx, &notebook)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNotebookResponse{
		Id: notebook.Id,
	}, nil
}

// Menampilkan detail notebook berdasarkan ID
func (c* notebookService) Show(ctx context.Context, id uuid.UUID) (*dto.ShowNotebookResponse, error) {
	notebook, err := c.notebookRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := dto.ShowNotebookResponse{
		Id: notebook.Id,
		Name: notebook.Name,
		ParentId: notebook.Parent_id,
		CreatedAt: notebook.Created_at,
		UpdatedAt: notebook.Updated_at,
	}
	return &res, nil
}

// Memperbarui notebook berdasarkan ID
func (c* notebookService) Update(ctx context.Context, req *dto.UpdateNotebookRequest) (*dto.UpdateNotebookResponse, error) {
	notebook, err := c.notebookRepository.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	notebook.Name = req.Name
	notebook.Updated_at = now

	err = c.notebookRepository.Update(ctx, notebook)

	if err != nil {
		return nil, err
	}

	res := dto.UpdateNotebookResponse{
		Id: notebook.Id,
	}
	return &res, nil
}

