package maliwan_fixed_asset_counting_rightsRoutes

import (
	"github.com/gofiber/fiber/v2"
	Controllermaliwan_fixed_asset_counting_rights "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_counting_rights/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupaMaliwan_Fixed_Asset_Counting_RightsRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-counting-rights")
	app.Get("/get-maliwan-fixed-asset-counting-rights", middleware.AuthorizationRequired(), Controllermaliwan_fixed_asset_counting_rights.GetAllHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_counting_rights_id", middleware.AuthorizationRequired(), Controllermaliwan_fixed_asset_counting_rights.GetById)
	app.Get("/get-user-id/:user_id", middleware.AuthorizationRequired(), Controllermaliwan_fixed_asset_counting_rights.GetUSer_IDHandler)
	app.Post("/create-maliwan-fixed-asset-rights", middleware.AuthorizationRequired(), Controllermaliwan_fixed_asset_counting_rights.Create)
	// app.Patch("/update-counting-rights/:maliwan_fixed_asset_counting_rights_id", Controllermaliwan_fixed_asset_counting_rights.UpdateMaliwan_Fixed_Asset_Counting_Rights)
	app.Delete("/delete-counting-rights/:maliwan_fixed_asset_counting_rights_id", middleware.AuthorizationRequired(), Controllermaliwan_fixed_asset_counting_rights.DeleteCountingRights)
}

// -fixed-asset-
