package Maliwan_fixed_asset_round_count_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_fixed_asset_round_count "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_round_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_fixed_asset_round_count_storyRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-round-count-story")
	app.Get("/get-maliwan-fixed-asset-round-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.GetAllHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.GetById)
	app.Post("/create-maliwan-fixed-asset-round-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.Create)
	app.Patch("/update-maliwan-round-count/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.UpdateMaliwan_Fixed_Asset_Round_Count_Story)
	app.Delete("/delete-maliwan-round-count/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.DeleteMaliwan_Fixed_Asset_Round_Count_Story)
}
