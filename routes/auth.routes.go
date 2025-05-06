package routes

import (
	"github.com/gofiber/fiber/v2"
	"railway-bac/controllers"
)
// all auth routes including oauth
func SetupAuthRoutes(router fiber.Router) {
	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
}