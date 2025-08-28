package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

var app = fiber.New()

func init() {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber on Vercel!")
	})
}
func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(app)(w, r)
}
