package route

import (
	"malai/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserHandlerRead)
	r.Get("/users", handler.UserHandlerGetAll)
	r.Post("/users", handler.UserHandlerCreate)
}
