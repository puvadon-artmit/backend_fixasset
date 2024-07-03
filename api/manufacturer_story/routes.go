package manufacturer_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerManufacturer_story "github.com/puvadon-artmit/gofiber-template/api/manufacturer_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupManufacturer_storyRoutes(router fiber.Router) {
	app := router.Group("manufacturer-story")
	app.Get("/get-manufacturer-story", middleware.AuthorizationRequired(), ControllerManufacturer_story.GetAllHandler)
	app.Get("/get-by-id/:manufacturer_story_id", middleware.AuthorizationRequired(), ControllerManufacturer_story.GetById)
	app.Post("/create-manufacturer-story", middleware.AuthorizationRequired(), ControllerManufacturer_story.Create)
}
