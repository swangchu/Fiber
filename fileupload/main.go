package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./uploads")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index.html", fiber.Map{})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("upload")
		if err != nil {
			return err
		}
		c.SaveFile(file, "uploads/"+file.Filename)

		return c.Render("views/index.html", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
