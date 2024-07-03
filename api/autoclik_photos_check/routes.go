package autoclik_photos_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	autoclik_photos_checkController "github.com/puvadon-artmit/gofiber-template/api/autoclik_photos_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Photos_checkRoutes(router fiber.Router) {

	autoclik_photos_checkGroup := router.Group("/autoclik_photos_check")
	autoclik_photos_checkGroup.Get("/get-all", middleware.AuthorizationRequired(), autoclik_photos_checkController.GetAll)
	autoclik_photos_checkGroup.Get("/get-by-id/:autoclik_photos_check_id", middleware.AuthorizationRequired(), autoclik_photos_checkController.GetByIdHandler)
	autoclik_photos_checkGroup.Post("/create", middleware.AuthorizationRequired(), autoclik_photos_checkController.Create)
	autoclik_photos_checkGroup.Get("/get-signed-url/:autoclik_photos_check_id", middleware.AuthorizationRequired(), autoclik_photos_checkController.GetByIdHandler)

	autoclik_photos_checkGroup.Delete("/delete-image/:autoclik_photos_check_id", middleware.AuthorizationRequired(), autoclik_photos_checkController.DeletePhotos)

}
