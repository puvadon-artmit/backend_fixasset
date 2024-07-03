package Autoclik_Round_CountRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Round_Count "github.com/puvadon-artmit/gofiber-template/api/autoclik_round_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Round_CountRoutes(router fiber.Router) {
	app := router.Group("autoclik_round_count")
	app.Get("/get-autoclik-round-count", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.GetAllHandler)
	app.Get("/get-by-autoclik-count-id/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.GetByNoHandler)
	app.Get("/get-by-id/:autoclik_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.GetById)
	app.Post("/create-autoclik-round-count", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.Create)
	app.Patch("/update-autoclik-round-count/:autoclik_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.UpdateAutoclik_Round_Count)
	app.Delete("/delete-autoclik-round-count/:autoclik_round_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count.DeleteAutoclik_Round_Count)
}
