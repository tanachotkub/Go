// routes/routes.go
package routes

import (
	"Go/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, memberHdl *handlers.MemberHandler) {
	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API Group
	api := app.Group("/api")
	api.Get("/members", memberHdl.GetMembers)
	api.Get("/members/:id", memberHdl.GetMemberByID)
	api.Post("/members", memberHdl.CreateMember)
	api.Put("/members/:id", memberHdl.UpdateMember)
	api.Delete("/members/:id", memberHdl.DeleteMember)
}
