package Status_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerStatus_Story "github.com/puvadon-artmit/gofiber-template/api/status_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupStatus_StoryRoutes(router fiber.Router) {
	app := router.Group("status-story")
	app.Get("/get-status-story", middleware.AuthorizationRequired(), ControllerStatus_Story.GetAllHandler)
	app.Get("/get-by-id/:status_story_id", middleware.AuthorizationRequired(), ControllerStatus_Story.GetById)
	app.Post("/create-status-story", middleware.AuthorizationRequired(), ControllerStatus_Story.Create)
	app.Patch("/update-status-story/:status_story_id", middleware.AuthorizationRequired(), ControllerStatus_Story.UpdateStatus_story)
	app.Delete("/delete-status-story/:status_story_id", middleware.AuthorizationRequired(), ControllerStatus_Story.DeleteStatus_story)
}
