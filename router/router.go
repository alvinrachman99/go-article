package router

import (
	"article/config"
	"article/handler"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB, cfg config.Config) {
	postsHandler := handler.NewPostsHandler(db, cfg)

	article := app.Group("/article")
	article.Post("/", postsHandler.Create)
	article.Get("/:limit/:offset", postsHandler.GetPosts)
	article.Get("/:id", postsHandler.GetPostsById)
	article.Put("/:id", postsHandler.Update)
	article.Delete("/:id", postsHandler.Delete)
}
