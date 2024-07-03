package storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	Ground_StoryController "github.com/puvadon-artmit/gofiber-template/api/ground_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupGround_storyRoutes(router fiber.Router) {
	app := router.Group("ground_story")
	app.Get("/get-ground-story", middleware.AuthorizationRequired(), Ground_StoryController.GetAll)
	app.Get("/get-by-id/:ground_story_id", middleware.AuthorizationRequired(), Ground_StoryController.GetById)
	app.Post("/create-ground-story", middleware.AuthorizationRequired(), Ground_StoryController.Create)
}
