package Autoclik_check_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_check_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_check_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_check_StoryRoutes(router fiber.Router) {
	app := router.Group("autoclik_check_story")
	app.Get("/get-autoclik-check-story", middleware.AuthorizationRequired(), ControllerAutoclik_check_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_check_Story.GetById)
	app.Post("/create-autoclik-check-story", middleware.AuthorizationRequired(), ControllerAutoclik_check_Story.Create)
	app.Patch("/update-autoclik-check-story/:autoclik_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_check_Story.UpdateAutoclik_check_Story)
	app.Delete("/delete-autoclik-check-story/:autoclik_check_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_check_Story.DeleteAutoclik_check_Story)
}
