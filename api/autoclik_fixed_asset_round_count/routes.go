package Autoclik_fixed_asset_round_countRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_fixed_asset_round_count "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_round_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_fixed_asset_round_countRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset-round-count")
	app.Get("/get-autoclik-fixed-asset-round-count", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.GetAllHandler)
	app.Get("/get-by-autoclik-count-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.GetByAutoclik_Fixed_Asset_Round_CountID)
	app.Get("/get-by-id/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.GetById)
	app.Post("/create-autoclik-round-count", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.Create)
	app.Patch("/update-autoclik-round-count/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.UpdateAutoclik_Fixed_Asset_Round_Count)
	app.Delete("/delete-autoclik-fixed-asset-round-count/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_fixed_asset_round_count.DeleteAutoclik_Fixed_Asset_Round_Count)
}
