package maliwan_fixed_asset_photos_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	maliwan_fixed_asset_photos_checkController "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_photos/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_Photos_checkRoutes(router fiber.Router) {

	maliwan_fixed_asset_photos_checkGroup := router.Group("/maliwan-fixed-asset-photos-check")
	maliwan_fixed_asset_photos_checkGroup.Get("/get-all", middleware.AuthorizationRequired(), maliwan_fixed_asset_photos_checkController.GetAll)
	maliwan_fixed_asset_photos_checkGroup.Get("/get-by-id/:maliwan_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_photos_checkController.GetByIdHandler)
	maliwan_fixed_asset_photos_checkGroup.Post("/create-picture", middleware.AuthorizationRequired(), maliwan_fixed_asset_photos_checkController.Create)
	maliwan_fixed_asset_photos_checkGroup.Get("/get-signed-url/:maliwan_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_photos_checkController.GetByIdHandler)

	maliwan_fixed_asset_photos_checkGroup.Delete("/delete-maliwan-fixed-asset-image/:maliwan_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_photos_checkController.DeletePhotos)

}
