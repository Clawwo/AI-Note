package repository

import (
	"ai-notetaking-be/internal/entity"
	"ai-notetaking-be/internal/pkg/serverutils"
	"ai-notetaking-be/pkg/database"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type INotebookRepository interface {
	UsingTx(ctx context.Context, tx database.DatabaseQueryer) INotebookRepository
	Create(ctx context.Context, notebook *entity.Notebook) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Notebook, error)
	Update(ctx context.Context, notebook *entity.Notebook) error
	DeleteById(ctx context.Context, id uuid.UUID) error
	NullifyParentById(ctx context.Context, parent_id uuid.UUID) error
}

type notebookRepository struct {
	db database.DatabaseQueryer
}

func (n *notebookRepository) UsingTx(ctx context.Context, tx database.DatabaseQueryer) INotebookRepository {
	return &notebookRepository{
		db: tx,
	}
}

// Membuat notebook baru
func (n *notebookRepository) Create(ctx context.Context, notebook *entity.Notebook) error {
	_, err := n.db.Exec(
		ctx,
		`INSERT INTO notebook (id, name, parent_id, created_at, updated_at, deleted_at, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		notebook.Id,
		notebook.Name,
		notebook.Parent_id,
		notebook.Created_at,
		notebook.Updated_at,
		notebook.Deleted_at,
		notebook.IsDeleted,
	)
	if err != nil {
		return err
	}

	return nil
}

// Mengambil notebook berdasarkan ID
func (n *notebookRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Notebook, error) {
	row := n.db.QueryRow(
		ctx,
		`SELECT id, name, parent_id, created_at, updated_at, deleted_at, is_deleted FROM notebook n WHERE n.is_deleted = false AND n.id = $1`,
		id,
	
	)	
	var notebook entity.Notebook
	err := row.Scan(
		&notebook.Id,
		&notebook.Name,
		&notebook.Parent_id,
		&notebook.Created_at,
		&notebook.Updated_at,
		&notebook.Deleted_at,
		&notebook.IsDeleted,
	)

	if err != nil {
		if errors .Is(err, pgx.ErrNoRows) {
			return nil, serverutils.ErrNotFound
		}
		return nil, err
	}
	return &notebook, nil
}

// Memperbarui notebook berdasarkan ID
func (n *notebookRepository) Update(ctx context.Context, notebook *entity.Notebook) error {
	_, err := n.db.Exec(
		ctx,
		`Update notebook SET name = $1, parent_id = $2, updated_at = $3 WHERE id = $4`,
		notebook.Name,
		notebook.Parent_id,
		notebook.Updated_at,
		notebook.Id,
	)
		
	if err != nil {
		return err
	}

	return nil
}

// Menghapus notebook berdasarkan ID
func (n *notebookRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	_, err := n.db.Exec(
		ctx,
		`Update notebook SET deleted_at = $1, is_deleted = true WHERE id = $2`,
		time.Now(),
		id,
	)
		
	if err != nil {
		return err
	}

	return nil
}

// Menghapus notebook berdasarkan ID
func (n *notebookRepository) NullifyParentById(ctx context.Context, parent_id uuid.UUID) error {
	_, err := n.db.Exec(
		ctx,
		`Update notebook SET parent_id = null, updated_at = $1 WHERE id = $2`,
		time.Now(),
		parent_id,
	)
		
	if err != nil {
		return err
	}

	return nil
}

func NewNotebookRepository(db *pgxpool.Pool) INotebookRepository {
	return &notebookRepository{
		db: db,
	}
}
