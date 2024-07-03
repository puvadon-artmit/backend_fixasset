package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Count "github.com/puvadon-artmit/gofiber-template/api/autoclik_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_CountRoutes(router fiber.Router) {
	app := router.Group("autoclik_count")
	app.Get("/get-autoclik-count", middleware.AuthorizationRequired(), ControllerAutoclik_Count.GetAllHandler)
	app.Get("/get-by-id/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count.GetById)
	app.Post("/create-autoclik-count", middleware.AuthorizationRequired(), ControllerAutoclik_Count.Create)
	app.Patch("/update-autoclik-count/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count.UpdateAutoclik_count)
	app.Delete("/delete-autoclik-count/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count.DeleteAutoclik_count)
}
