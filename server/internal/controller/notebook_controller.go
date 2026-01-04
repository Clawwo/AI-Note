package controller

import (
	"ai-notetaking-be/internal/dto"
	"ai-notetaking-be/internal/pkg/serverutils"
	"ai-notetaking-be/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type INotebookController interface {
	RegisterRoutes(r fiber.Router)
	Create(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	MoveNotebook(ctx *fiber.Ctx) error
}

type notebookController struct {
	service service.INotebookService
}

func NewNotebookController(service service.INotebookService) INotebookController {
	return &notebookController{service: service}
}

func (c *notebookController) RegisterRoutes(r fiber.Router) {
	h := r.Group("/notebook/v1")
	h.Get("", c.GetAll)
	h.Post("", c.Create)
	h.Get("/:id", c.Show)
	h.Put("/:id", c.Update)
	h.Delete("/:id", c.Delete)
	h.Put("/:id/move", c.MoveNotebook)
}

// Menampilkan semua notebook
func (c *notebookController) GetAll(ctx *fiber.Ctx) error {
	res, err := c.service.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse("Berhasil menampilkan semua Notebook", res))
}

// Membuat notebook baru
func (c *notebookController) Create(ctx *fiber.Ctx) error {
	var req dto.CreateNotebookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := serverutils.ValidateRequest(req)
	if err != nil {
		return err
	}

	res, err := c.service.Create(ctx.Context(), &req)
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse("Berhasil membuat Notebook", res))
}

// Menampilkan detail notebook berdasarkan ID
func (c *notebookController) Show(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	id, _ := uuid.Parse(idParams)


	res, err := c.service.Show(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse("Berhasil menampilkan Notebook", res))
}
// Memperbarui notebook berdasarkan ID
func (c *notebookController) Update(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	id, _ := uuid.Parse(idParams)

	var req dto.UpdateNotebookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	req.Id = id

	res, err := c.service.Update(ctx.Context(), &req)
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse("Berhasil memperbarui Notebook", res))
}

// Menghapus notebook berdasarkan ID
func (c *notebookController) Delete(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	id, _ := uuid.Parse(idParams)

	err := c.service.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse[any]("Berhasil menghapus Notebook", nil))
}

func (c *notebookController) MoveNotebook(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	id, _ := uuid.Parse(idParams)

	var req dto.MoveNotebookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	req.Id = id

	res, err := c.service.MoveNotebook(ctx.Context(), &req)
	if err != nil {
		return err
	}

	return ctx.JSON(serverutils.SuccessResponse("Berhasil memindahkan Notebook", res))
}