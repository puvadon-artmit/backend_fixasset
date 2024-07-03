package Assets_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	Assets_StoryController "github.com/puvadon-artmit/gofiber-template/api/assets_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAssets_StoryRoutes(router fiber.Router) {
	app := router.Group("assets_story")
	app.Get("/get-assets-story", middleware.AuthorizationRequired(), Assets_StoryController.GetAll)
	app.Get("/get-by-id/:assets_story_id", middleware.AuthorizationRequired(), Assets_StoryController.GetById)
	app.Post("/create-assets-story", middleware.AuthorizationRequired(), Assets_StoryController.Create)
}
