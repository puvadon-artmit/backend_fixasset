package LocationRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerLocation_story "github.com/puvadon-artmit/gofiber-template/api/location_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupLocation_storyRoutes(router fiber.Router) {
	app := router.Group("location_story")
	app.Get("/get-location_story", middleware.AuthorizationRequired(), ControllerLocation_story.GetAllHandler)
	app.Get("/get-by-id/:location_story_id", middleware.AuthorizationRequired(), ControllerLocation_story.GetById)
	app.Post("/create-location_story", middleware.AuthorizationRequired(), ControllerLocation_story.Create)
	app.Patch("/update-location_story/:location_story_id", middleware.AuthorizationRequired(), ControllerLocation_story.UpdateLocation_story)
	app.Delete("/delete-location_story/:controller_id", middleware.AuthorizationRequired(), ControllerLocation_story.DeleteLocation_story)
}
