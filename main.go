package main

import (
	"malai/database"
	"malai/database/migration"
	"malai/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseMysql()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}
