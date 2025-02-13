package main

import (
	"article/config"
	"article/database"
	"article/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app, db, cfg)

	log.Fatal(app.Listen(":3000"))
}
