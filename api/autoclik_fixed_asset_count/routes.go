package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Fixed_Asset_Count "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Fixed_Asset_CountRoutes(router fiber.Router) {
	app := router.Group("autoclik-fixed-asset-count")
	app.Get("/get-autoclik-fixed-asset-count", ControllerAutoclik_Fixed_Asset_Count.GetAllHandler)
	app.Get("/pull-data/:autoclik_fixed_asset_count_id", ControllerAutoclik_Fixed_Asset_Count.GetByIdHandler)
	app.Get("/get-by-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Count.GetById)
	app.Post("/create-fixed-asset-count", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Count.Create)
	app.Patch("/update-autoclik-fixed-asset-count/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Count.UpdateAutoclik_Fixed_Asset_Count)
	app.Delete("/delete-autoclik-fixed-asset-count/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Fixed_Asset_Count.DeleteAutoclik_Fixed_Asset_Count)
}
