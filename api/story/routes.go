package storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	StoryController "github.com/puvadon-artmit/gofiber-template/api/story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupStoryRoutes(router fiber.Router) {
	app := router.Group("story")
	app.Get("/get-story", middleware.AuthorizationRequired(), StoryController.GetAll)
	app.Get("/get-by-id/:story_id", middleware.AuthorizationRequired(), StoryController.GetById)
	app.Post("/create-story", middleware.AuthorizationRequired(), StoryController.Create)
}
