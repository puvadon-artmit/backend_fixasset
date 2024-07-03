package Autoclik_Fixed_AssetRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_data "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_fixed_assetRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset")

	app.Get("/get-no", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByItem_No_Autoclik_BinHandler)
	app.Get("/get-product-group", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByProductGroupHandler)
	app.Get("/filter-autoclik-fixed-asset", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllFA_Location_Code)

	app.Get("/get-count-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCountAutoclik_data)

	app.Get("/pull-autoclik-fixed-asset", middleware.AuthorizationRequired(), ControllerAutoclik_data.Pull_information)
	app.Get("/get-all-autoclik-fixed-asset", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllAutoclik_Fixed_Asset)

	app.Get("/get-count-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCountAllItemBin)
	app.Get("/clear-and-pull-item-bin", middleware.AuthorizationRequired(), ControllerAutoclik_data.Clearandsuckbinitem)

	app.Get("/filter-autoclik-fixed-data", ControllerAutoclik_data.GetFilterAutoclik_FixedHandler)

	app.Get("/filter-fixed-data", ControllerAutoclik_data.GetFilterFixedHandler)

	// app.Get("/filter-item-bin", ControllerAutoclik_data.GetChoosebytypeAutoclikbin)
}
