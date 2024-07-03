package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Fixed_Asset_Count "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_CountRoutes(router fiber.Router) {
	app := router.Group("maliwan-fixed-asset-count")
	app.Get("/get-fixed-asset-count", ControllerMaliwan_Fixed_Asset_Count.GetAllHandler)
	// app.Get("/pull-data/:maliwan_fixed_asset_count_id", ControllerMaliwan_Fixed_Asset_Count.GetByIdHandler)
	app.Get("/get-by-id/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Count.GetById)

	app.Post("/create-fixed-asset-count", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Count.Create)
	app.Patch("/update-maliwan-fixed-asset-count/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Count.UpdateMaliwan_Fixed_Asset_Count)
	app.Delete("/delete-maliwan-fixed-asset-count/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Fixed_Asset_Count.DeleteMaliwan_Fixed_Asset_Count)
}
