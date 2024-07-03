package Asset_count_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAsset_count_story "github.com/puvadon-artmit/gofiber-template/api/asset_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAsset_count_storyRoutes(router fiber.Router) {
	app := router.Group("asset-count-story")
	app.Get("/get-asset-count-story", middleware.AuthorizationRequired(), ControllerAsset_count_story.GetAllHandler)
	app.Get("/get-by-id/:asset_count_story_id", middleware.AuthorizationRequired(), ControllerAsset_count_story.GetById)
	app.Post("/create-asset_count_story", middleware.AuthorizationRequired(), ControllerAsset_count_story.Create)
	app.Patch("/update-asset_count_story/:asset_count_story_id", middleware.AuthorizationRequired(), ControllerAsset_count_story.UpdateAsset_count_story)
	app.Delete("/delete-asset_count_story/:asset_count_story_id", middleware.AuthorizationRequired(), ControllerAsset_count_story.DeleteAsset_count_story)
}
