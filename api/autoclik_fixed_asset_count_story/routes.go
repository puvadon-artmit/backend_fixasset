package Autoclik_Fixed_Asset_count_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Fixed_Asset_count_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Fixed_Asset_count_StoryRoutes(router fiber.Router) {
	app := router.Group("asset-fixed-asset-count-story")
	app.Get("/get-autoclik-fixed-asset-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_count_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_count_Story.GetById)
	app.Post("/create-autoclik-fixed-asset-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_count_Story.Create)
	app.Patch("/update-autoclik-fixed-asset-count-story/:autoclik_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_count_Story.UpdateAutoclik_Fixed_Asset_count_Story)
	app.Delete("/delete-autoclik_fixed_asset_count_story/:autoclik_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_count_Story.DeleteAutoclik_Fixed_Asset_count_Story)
}
