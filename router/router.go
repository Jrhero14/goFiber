package router

import (
	noteRoutes "github.com/Jrhero14/goFiber/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api)
}
