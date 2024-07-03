package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Update_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_update_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Update_StoryRoutes(router fiber.Router) {
	app := router.Group("autoclik-update-story")
	app.Get("/get-autoclik-update-story", middleware.AuthorizationRequired(), ControllerAutoclik_Update_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_update_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Update_Story.GetById)
	app.Post("/create-autoclik-update-story", middleware.AuthorizationRequired(), ControllerAutoclik_Update_Story.Create)
	app.Patch("/update-autoclik-update-story/:autoclik_update_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Update_Story.UpdateInspection_story)
	app.Delete("/delete-autoclik-update-story/:autoclik_update_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Update_Story.DeleteInspection_story)
}
