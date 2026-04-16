// routes/routes.go
package routes

import (
	"Go/handlers"
	"Go/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, memberHdl *handlers.MemberHandler) {
	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	//(Public)
	api.Post("/login", memberHdl.LoginMember)
	api.Post("/register", memberHdl.CreateMember)

	//(Protected)
	members := api.Group("/members", middlewares.JWTMiddleware)
	members.Get("/", memberHdl.GetMembers)
	members.Get("/:id", memberHdl.GetMemberByID)
	members.Put("/:id", memberHdl.UpdateMember)
	members.Delete("/:id", memberHdl.DeleteMember)
}
