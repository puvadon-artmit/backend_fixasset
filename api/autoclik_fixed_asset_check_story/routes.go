package Autoclik_Fixed_Asset_check_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Fixed_Asset_check_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Fixed_Asset_check_StoryRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset-check-story")
	app.Get("/get-autoclik-fixed-asset-check-story", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_check_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_check_Story.GetById)
	app.Post("/create-autoclik-fixed-asset-check-story", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_check_Story.Create)
	app.Patch("/update-autoclik-check-story/:autoclik_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_check_Story.UpdateAutoclik_Fixed_Asset_check_Story)
	app.Delete("/delete-autoclik-check-story/:autoclik_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_check_Story.DeleteAutoclik_Fixed_Asset_check_Story)
}
