package autoclik_fixed_asset_photos_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	autoclik_fixed_asset_photos_checkController "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_photos/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Fixed_Asset_Photos_checkRoutes(router fiber.Router) {

	autoclik_fixed_asset_photos_checkGroup := router.Group("/autoclik-fixed-asset-photos-check")
	autoclik_fixed_asset_photos_checkGroup.Get("/get-all", middleware.AuthorizationRequired(), autoclik_fixed_asset_photos_checkController.GetAll)
	autoclik_fixed_asset_photos_checkGroup.Get("/get-by-id/:autoclik_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), autoclik_fixed_asset_photos_checkController.GetByIdHandler)
	autoclik_fixed_asset_photos_checkGroup.Post("/create-picture", middleware.AuthorizationRequired(), autoclik_fixed_asset_photos_checkController.Create)
	autoclik_fixed_asset_photos_checkGroup.Get("/get-signed-url/:autoclik_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), autoclik_fixed_asset_photos_checkController.GetByIdHandler)

	autoclik_fixed_asset_photos_checkGroup.Delete("/delete-autoclik-fixed-asset-image/:autoclik_fixed_asset_photos_check_id", middleware.AuthorizationRequired(), autoclik_fixed_asset_photos_checkController.DeletePhotos)

}
