package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Fixed_Asset_Store "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_store/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Fixed_Asset_StoreRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset-store")
	app.Get("/get-autoclik-fixed-asset-store", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.GetAllHandler)
	app.Get("/pull-data-by-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.GetByIdHandler)
	app.Get("/get-by-id/:autoclik_fixed_asset_store_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.GetById)
	app.Get("/get-by-autoclik-fixed-asset-count-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.GetByAutoclik_Fixed_AssetC_CountIDHandler)
	app.Post("/create-fixed-asset-store", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.Create)
	app.Patch("/update-autoclik-fixed-asset-store/:autoclik_fixed_asset_store_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.UpdateAutoclik_Fixed_Asset_Store)
	app.Delete("/delete-all-by-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Store.DeleteAutoclik_Fixed_Asset_Store)
}
