package Maliwan_Fixed_Asset_count_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Fixed_Asset_count_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_count_StoryRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-count-story")
	app.Get("/get-maliwan-fixed-asset-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_count_Story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_count_Story.GetById)
	app.Post("/create-maliwan-fixed-asset-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_count_Story.Create)
	app.Patch("/update-maliwan-fixed-asset-count-story/:maliwan_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_count_Story.UpdateMaliwan_Fixed_Asset_count_Story)
	app.Delete("/delete-maliwan_fixed_asset_count_story/:maliwan_fixed_asset_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_count_Story.DeleteMaliwan_Fixed_Asset_count_Story)
}
