package Maliwan_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	maliwan_checkController "github.com/puvadon-artmit/gofiber-template/api/maliwan_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_checkRoutes(router fiber.Router) {
	maliwan_checkGroup := router.Group("/maliwan_check")
	maliwan_checkGroup.Get("/get-maliwan-check", middleware.AuthorizationRequired(), maliwan_checkController.GetAll)
	maliwan_checkGroup.Get("/get-by-id/:maliwan_check_id", middleware.AuthorizationRequired(), maliwan_checkController.GetById)

	maliwan_checkGroup.Get("/get-signed-url/:maliwan_check_id", middleware.AuthorizationRequired(), maliwan_checkController.GetByIdHandler)

	// maliwan_checkGroup.Get("/get-round-count/:round_count_id", maliwan_checkController.GetByRound_CountHandler)
	maliwan_checkGroup.Get("/get-round-id/:maliwan_round_count_id", middleware.AuthorizationRequired(), maliwan_checkController.GetByMaliwanRound_CountHandler)
	maliwan_checkGroup.Get("/get-round-latest/:maliwan_round_count_id", middleware.AuthorizationRequired(), maliwan_checkController.GetByProperty_code_Id)
	// maliwan_checkGroup.Get("/get-round-latest", maliwan_checkController.GetAllHandler)
	maliwan_checkGroup.Post("/create-maliwan-check", middleware.AuthorizationRequired(), maliwan_checkController.CreateNew)
	maliwan_checkGroup.Patch("/update/:maliwan_check_id", middleware.AuthorizationRequired(), maliwan_checkController.UpdateMaliwan_checkHandler)

	maliwan_checkGroup.Get("/filter-no-and-count-id", middleware.AuthorizationRequired(), maliwan_checkController.GetByNoAndMaliwan_Round_CountIDHandler)
}
