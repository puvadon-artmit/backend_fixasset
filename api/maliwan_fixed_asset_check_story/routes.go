package Maliwan_Fixed_Asset_check_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Fixed_Asset_check_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_check_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_check_StoryRoutes(router fiber.Router) {
	app := router.Group("maliwan-fa-check-story")
	app.Get("/get-maliwan-fa-check-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_check_Story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_check_Story.GetById)
	app.Post("/create-maliwan-fa-check-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_check_Story.Create)
	app.Patch("/update-maliwan-check-story/:maliwan_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_check_Story.UpdateMaliwan_Fixed_Asset_check_Story)
	app.Delete("/delete-maliwan-check-story/:maliwan_fixed_asset_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_check_Story.DeleteMaliwan_Fixed_Asset_check_Story)
}
