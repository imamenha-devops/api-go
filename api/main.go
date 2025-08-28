package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB connection string dari ENV Vercel (setting di Dashboard → Environment Variables)
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go + Fiber + Vercel!")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []map[string]interface{}
		db.Raw("SELECT id, name FROM users LIMIT 10").Scan(&users)
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":3000")) // Vercel akan redirect PORT ke 3000
}
