package Autoclik_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	autoclik_checkController "github.com/puvadon-artmit/gofiber-template/api/autoclik_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_checkRoutes(router fiber.Router) {
	autoclik_checkGroup := router.Group("/autoclik_check")
	autoclik_checkGroup.Get("/get-autoclik-check", middleware.AuthorizationRequired(), autoclik_checkController.GetAll)
	autoclik_checkGroup.Get("/get-by-id/:autoclik_check_id", middleware.AuthorizationRequired(), autoclik_checkController.GetById)
	autoclik_checkGroup.Get("/get-signed-url/:autoclik_check_id", middleware.AuthorizationRequired(), autoclik_checkController.GetBySignedURLHandler)
	// autoclik_checkGroup.Get("/get-round-count/:round_count_id", autoclik_checkController.GetByRound_CountHandler)
	autoclik_checkGroup.Get("/get-round-id/:autoclik_round_count_id", middleware.AuthorizationRequired(), autoclik_checkController.GetByAutoclikRound_CountHandler)
	autoclik_checkGroup.Get("/get-round-latest/:autoclik_round_count_id", middleware.AuthorizationRequired(), autoclik_checkController.GetByAutoclikRound_CountHandlerlatest)
	// autoclik_checkGroup.Get("/get-round-latest", autoclik_checkController.GetAllHandler)
	autoclik_checkGroup.Post("/create-autoclik-check", middleware.AuthorizationRequired(), autoclik_checkController.CreateNew)
	autoclik_checkGroup.Patch("/update/:autoclik_check_id", middleware.AuthorizationRequired(), autoclik_checkController.UpdateAutoclik_checkHandler)

	autoclik_checkGroup.Get("/filter-item-no-and-count-id", middleware.AuthorizationRequired(), autoclik_checkController.GetByItem_noAndAutoclik_Round_CountIDHandler)
}
