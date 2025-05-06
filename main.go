package main

import (
	"log"
	"railway-bac/initializers"
	"railway-bac/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Use(cors.New())
	app.Mount("/api", micro)
	app.Use(logger.New())

	micro.Route("/auth", func(router fiber.Router) {
			routes.SetupAuthRoutes(router)
		})


	micro.Route("/users", func(router fiber.Router) {
			routes.SetupUserRoutes(router)
		})
	// micro.Route("/auth", func(router fiber.Router) {
	// 		routes.AuthRoutes(router)
	// 	})
	// routes.AuthRoutes(api)
	// routes.RecordRoutes(api)

	// Start server
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
