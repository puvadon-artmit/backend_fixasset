package Maliwan_fixed_asset_round_countRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_fixed_asset_round_count "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_round_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_fixed_asset_round_countRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-round-count")
	app.Get("/get-maliwan-round-count", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.GetAllHandler)
	app.Get("/get-by-maliwan-count-id/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.GetByMaliwan_Fixed_Asset_Round_CountID)
	app.Get("/get-by-id/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.GetById)
	app.Post("/create-maliwan-round-count", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.Create)
	app.Patch("/update-maliwan-round-count/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.UpdateMaliwan_Fixed_Asset_Round_Count)
	app.Delete("/delete-maliwan-fixed-asset-round-count/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_fixed_asset_round_count.DeleteMaliwan_Fixed_Asset_Round_Count)
}
