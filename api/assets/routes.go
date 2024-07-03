package AssetsRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAssets "github.com/puvadon-artmit/gofiber-template/api/assets/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAssetsRoutes(router fiber.Router) {
	app := router.Group("assets")
	app.Get("/get-assets", middleware.AuthorizationRequired(), ControllerAssets.GetAllHandler)
	app.Get("/get-property-code/:property_code", middleware.AuthorizationRequired(), ControllerAssets.GetByProperty_codeHandler)
	app.Delete("/get-clear", middleware.AuthorizationRequired(), ControllerAssets.GetDeletePropertyCodes)
	app.Get("/get-category-id/:category_id", middleware.AuthorizationRequired(), ControllerAssets.GetByCategoryHandler)
	app.Get("/get-by-id/:assets_id", middleware.AuthorizationRequired(), ControllerAssets.GetByIdHandler)
	app.Get("/get-by-ids/:assets_id", middleware.AuthorizationRequired(), ControllerAssets.GetById)
	app.Post("/create-assets", middleware.AuthorizationRequired(), ControllerAssets.Create)
	app.Patch("/update-assets/:assets_id", middleware.AuthorizationRequired(), ControllerAssets.UpdateAssetByID)
	app.Delete("/delete-assets/:assets_id", middleware.AuthorizationRequired(), ControllerAssets.DeleteAsset)
	app.Get("/getcount-assets", middleware.AuthorizationRequired(), ControllerAssets.GetCountAssets)
	app.Get("/getcount-allassets", middleware.AuthorizationRequired(), ControllerAssets.GetCountAllAssets)

}
