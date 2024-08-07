package main

import (
	"github.com/gofiber/fiber/v2"
	"go-htmx/database"
	"strconv"
)

func UseRoutes(app *fiber.App) {
	app.Get("/", index)
	app.Post("/", postTodo)
	app.Delete("/", deleteTodo)
}

func index(ctx *fiber.Ctx) error {
	todos, err := database.GetTodos()
	if err != nil {
		return ctx.Render("index", fiber.Map{
			"Error": err.Error(),
			"Todos": []database.Todo{},
		})
	}
	return ctx.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func postTodo(ctx *fiber.Ctx) error {

	name := ctx.FormValue("name")
	if len(name) == 0 {
		return ctx.Render("error", fiber.Map{
			"Error": "Veuillez saisir un nom",
		})
	}

	res, err := database.PostTodo(database.Todo{Name: name})
	if err != nil {
		return ctx.Render("error", fiber.Map{
			"Error": err.Error(),
		})
	}
	return ctx.Render("post-todo", fiber.Map{
		"Todo": res,
	})

}

func deleteTodo(ctx *fiber.Ctx) error {
	id := ctx.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Render("error", fiber.Map{
			"Error": err.Error(),
		})
	}

	err = database.DeleteTodo(idInt)
	if err != nil {
		return ctx.Render("error", fiber.Map{
			"Error": err.Error(),
		})
	}

	return ctx.Render("delete-todo", fiber.Map{
		"Id": id,
	})
}
