package routes

import (
	"apibooking/controllers"
	"apibooking/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	booking := r.Group("/api")

	booking.Get("/", middlewares.AuthMiddleware, controllers.Index)
	booking.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	booking.Post("/", middlewares.AuthMiddleware, controllers.Create)
	booking.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
	booking.Delete("/:id", middlewares.AuthMiddleware, controllers.Delete)
}