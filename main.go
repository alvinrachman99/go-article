package main

import (
	"article/config"
	"article/database"
	"article/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)

	app := fiber.New()

	router.SetupRoutes(app, db, cfg)

	log.Fatal(app.Listen(":3000"))
}
