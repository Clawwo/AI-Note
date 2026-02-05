package main

import (
	"ai-notetaking-be/internal/controller"
	"ai-notetaking-be/internal/pkg/serverutils"
	"ai-notetaking-be/internal/repository"
	"ai-notetaking-be/internal/service"
	"ai-notetaking-be/pkg/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(cors.New())

	app.Use(serverutils.ErrorHandlerMiddleware())

	db := database.ConnectDB(os.Getenv("DB_CONNECTION_STRING"))

	notebookRepository := repository.NewNotebookRepository(db)

	notebookService := service.NewNotebookService(
		notebookRepository, db,
	)

	notebookController := controller.NewNotebookController(
		notebookService,
	)
	

	api := app.Group("/api")
	notebookController.RegisterRoutes(api)

	log.Fatal(app.Listen(":3000"))
}
