// user.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"railway-bac/controllers"
	"railway-bac/middleware"
)

func SetupUserRoutes(router fiber.Router) {
	router.Post("/records", middleware.DeserializeUser, controllers.CreateRecord)
	router.Get("/records", middleware.DeserializeUser, controllers.GetRecords)
}