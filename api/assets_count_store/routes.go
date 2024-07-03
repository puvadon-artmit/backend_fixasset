package Assets_count_StoreRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAssets_count_Store "github.com/puvadon-artmit/gofiber-template/api/assets_count_store/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAssets_count_StoreRoutes(router fiber.Router) {
	app := router.Group("assets-count-store")
	app.Get("/get-assets-count-store", middleware.AuthorizationRequired(), ControllerAssets_count_Store.GetAllHandler)
	app.Get("/get-by-id/:assets_count_store_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.GetById)
	app.Get("/get-main-category/:main_category_id", ControllerAssets_count_Store.GetCategoryIDsByMainCategoryIDHandler)
	app.Get("/get-by-asset-count-ids/:asset_count_id", ControllerAssets_count_Store.GetAsset_Count_By_CategoryHandler)
	app.Get("/get-by-asset-category-id/:category_id", ControllerAssets_count_Store.GetGetAssetbyCategoryIDHandler)
	app.Get("/asset-count-data-store/:asset_count_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.GetAsset_Count_By_Main_CategoryHandler)

	app.Get("/get-by-asset-count-id/:asset_count_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.GetByAssets_Count_StoreIDHandler)
	app.Post("/create-assets-count-store", middleware.AuthorizationRequired(), ControllerAssets_count_Store.Create)
	app.Patch("/update-assets-count-store/:assets_count_store_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.UpdateAssets_count_Store)
	app.Delete("/delete-assets-count-store/:assets_count_store_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.DeleteAssets_count_Store)
	app.Delete("/delete-all-assets-count-id/:asset_count_id", middleware.AuthorizationRequired(), ControllerAssets_count_Store.DeleteAssets_Count_ID)
}
