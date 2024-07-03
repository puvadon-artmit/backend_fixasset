package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Fixed_Asset_Store "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_store/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_StoreRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-store")
	app.Get("/get-maliwan-fixed-asset", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetAllHandler)
	// app.Get("/pull-data", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.PullDataHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_store_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetById)
	app.Get("/get-by-maliwan-fixed-asset-count-id/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetByMaliwan_Fixed_AssetC_CountIDHandler)
	app.Post("/create-fixed-asset-store", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.Create)
	app.Patch("/update-maliwan-fixed-asset-store/:maliwan_fixed_asset_store_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.UpdateMaliwan_Fixed_Asset_Store)
	app.Delete("/delete-all-by-id/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.DeleteMaliwan_Fixed_Asset_Store)
	app.Get("/get-maliwan-fa-count-id-pull-store/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.Import_data_Store)

	app2 := router.Group("maliwan-fixed-asset")
	app2.Get("/get-maliwan-fixed-asset", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetAllMaliwan_Fixed_AssetHandler)
	app2.Get("/get-limit-page", ControllerMaliwan_Fixed_Asset_Store.GetMaliwan_Fixed_AssetsHandler)
	app2.Get("/get-no", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetByNoHandler)
	app2.Get("/get-branch", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetByBranchHandler)
	app2.Get("/pull-data", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.PullDataHandler)

	app3 := router.Group("maliwan-fa-update-story")
	app3.Post("/create-maliwan-fa-update-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.CreateMaliwan_Update_Fixed_Asset_StoryHandler)
	app3.Get("/get-maliwan-fa-update-story", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Store.GetAllMaliwan_Update_Fixed_Asset_StoryHandler)

}
