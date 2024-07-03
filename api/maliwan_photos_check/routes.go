package maliwan_photos_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	maliwan_photos_checkController "github.com/puvadon-artmit/gofiber-template/api/maliwan_photos_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Photos_checkRoutes(router fiber.Router) {

	maliwan_photos_checkGroup := router.Group("/maliwan_photos_check")
	maliwan_photos_checkGroup.Get("/get-all", middleware.AuthorizationRequired(), maliwan_photos_checkController.GetAll)
	maliwan_photos_checkGroup.Get("/get-by-id/:maliwan_photos_check_id", middleware.AuthorizationRequired(), maliwan_photos_checkController.GetByIdHandler)
	maliwan_photos_checkGroup.Post("/create-picture", middleware.AuthorizationRequired(), maliwan_photos_checkController.Create)
	maliwan_photos_checkGroup.Get("/get-signed-url/:maliwan_photos_check_id", middleware.AuthorizationRequired(), maliwan_photos_checkController.GetByIdHandler)

	maliwan_photos_checkGroup.Delete("/delete-maliwan-image/:maliwan_photos_check_id", middleware.AuthorizationRequired(), maliwan_photos_checkController.DeletePhotos)

}
