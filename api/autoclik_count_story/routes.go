package Autoclik_count_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_count_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_count_StoryRoutes(router fiber.Router) {
	app := router.Group("autoclik_count_story")
	app.Get("/get-autoclik-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_count_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Story.GetById)
	app.Post("/create-autoclik-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_count_Story.Create)
	app.Patch("/update-autoclik-count-story/:autoclik_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Story.UpdateAutoclik_count_Story)
	app.Delete("/delete-autoclik-count-story/:autoclik_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Story.DeleteAutoclik_count_Story)
}
