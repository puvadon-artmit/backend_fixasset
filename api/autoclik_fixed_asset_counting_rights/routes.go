package autoclik_fixed_asset_counting_rightsRoutes

import (
	"github.com/gofiber/fiber/v2"
	Controllerautoclik_fixed_asset_counting_rights "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_counting_rights/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupaAutoclik_Fixed_Asset_Counting_RightsRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset-counting-rights")
	app.Get("/get-autoclik-fixed-asset-counting-rights", middleware.AuthorizationRequired(), Controllerautoclik_fixed_asset_counting_rights.GetAllHandler)
	app.Get("/get-by-id/:autoclik_fixed_asset_counting_rights_id", middleware.AuthorizationRequired(), Controllerautoclik_fixed_asset_counting_rights.GetById)
	app.Get("/get-user-id/:user_id", middleware.AuthorizationRequired(), Controllerautoclik_fixed_asset_counting_rights.GetUSer_IDHandler)
	app.Post("/create-autoclik-fixed-asset-rights", middleware.AuthorizationRequired(), Controllerautoclik_fixed_asset_counting_rights.Create)
	// app.Patch("/update-counting-rights/:autoclik_fixed_asset_counting_rights_id", Controllerautoclik_fixed_asset_counting_rights.UpdateAutoclik_Fixed_Asset_Counting_Rights)
	app.Delete("/delete-counting-rights/:autoclik_fixed_asset_counting_rights_id", middleware.AuthorizationRequired(), Controllerautoclik_fixed_asset_counting_rights.DeleteCountingRights)
}

// -fixed-asset-
