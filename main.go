package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./views", ".gohtml")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "public/")

	UseRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
