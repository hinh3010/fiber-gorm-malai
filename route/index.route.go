package route

import (
	"malai/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserHandlerRead)
	r.Get("/users", handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Get("/users/lv2/:id", handler.UserHandlerGetByIdLv2)
	r.Post("/users", handler.UserHandlerCreate)
	r.Put("/users/:id", handler.UserHandlerUpdateById)
}
