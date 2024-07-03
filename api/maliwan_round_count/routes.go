package Maliwan_Round_CountRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Round_Count "github.com/puvadon-artmit/gofiber-template/api/maliwan_round_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Round_CountRoutes(router fiber.Router) {
	app := router.Group("maliwan_round_count")
	app.Get("/get-maliwan-round-count", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.GetAllHandler)
	app.Get("/get-by-maliwan-count-id/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.GetByNoHandler)
	app.Get("/get-by-id/:maliwan_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.GetById)
	app.Post("/create-maliwan-round-count", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.Create)
	app.Patch("/update-maliwan-round-count/:maliwan_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.UpdateMaliwan_Round_Count)
	app.Delete("/delete-maliwan-round-count/:maliwan_round_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count.DeleteMaliwan_Round_Count)
}
