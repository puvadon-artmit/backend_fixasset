package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerInspection "github.com/puvadon-artmit/gofiber-template/api/inspection-assets/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAsset_countRoutes(router fiber.Router) {
	app := router.Group("inspection")
	app.Get("/get-inspection", middleware.AuthorizationRequired(), ControllerInspection.GetAllHandler)
	app.Get("/get-by-id/:asset_count_id", middleware.AuthorizationRequired(), ControllerInspection.GetAssetCountByIdHandler)
	app.Get("/get-by-count-id/:asset_count_id", ControllerInspection.GetById)
	app.Post("/create-inspection", middleware.AuthorizationRequired(), ControllerInspection.Create)
	app.Patch("/update-inspection/:asset_count_id", middleware.AuthorizationRequired(), ControllerInspection.UpdateAsset_count)
	app.Delete("/delete-inspection/:controller_id", middleware.AuthorizationRequired(), ControllerInspection.DeleteAsset_count)
	app.Get("/getcount-asset-count", middleware.AuthorizationRequired(), ControllerInspection.GetCountAsset_count)
}
