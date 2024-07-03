package Responsible_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerResponsible_Story "github.com/puvadon-artmit/gofiber-template/api/responsible-story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupResponsible_StoryRoutes(router fiber.Router) {
	app := router.Group("responsible-story")
	app.Get("/get-responsible-story", middleware.AuthorizationRequired(), ControllerResponsible_Story.GetAllHandler)
	app.Get("/get-by-id/:responsible_story_id", middleware.AuthorizationRequired(), ControllerResponsible_Story.GetById)
	app.Post("/create-responsible-story", middleware.AuthorizationRequired(), ControllerResponsible_Story.Create)
	app.Patch("/update-responsible-story/:responsible_story_id", middleware.AuthorizationRequired(), ControllerResponsible_Story.UpdateResponsible_Story)
	app.Delete("/delete-responsible-story/:responsible_story_id", middleware.AuthorizationRequired(), ControllerResponsible_Story.DeleteResponsible_Story)
}
