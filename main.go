package main

import (
	"example.com/go-blog-api/post"
	"example.com/go-blog-api/tag"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {

	//posts api
	app.Get("/api/post", post.GetPosts)
	app.Get("/api/post/:id", post.GetPost)
	app.Post("/api/post", post.NewPost)
	app.Delete("/api/post/:id", post.DeletePost)
	app.Put("/api/post/:id", post.UpdatePost)

	//tags api
	app.Get("/api/tag", tag.GetTags)
	//app.Get("/api/tag/:id", tag.GetTag)
	app.Post("/api/tag", tag.NewTag)
	//app.Delete("/api/tag/:id", tag.DeleteTag)
	app.Put("/api/tag/:id", tag.UpdateTag)
}

func main() {
	post.InitialMigration()
	tag.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}
