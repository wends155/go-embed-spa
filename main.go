package main

import (
	"log"

	"go-embed-spa/frontend"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type ReturnData struct {
	Message string
}

func handleAPI(c *fiber.Ctx) error {
	data := ReturnData{
		Message: "hello from server",
	}

	log.Println(c.Request().URI())
	return c.JSON(data)
}

func main() {
	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Get("/api", handleAPI)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))

	log.Println("Server listening to :8080")
	log.Fatal(app.Listen(":8080"))

}
