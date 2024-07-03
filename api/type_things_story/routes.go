package Type_things_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerType_things_story "github.com/puvadon-artmit/gofiber-template/api/type_things_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupType_things_storyRoutes(router fiber.Router) {
	app := router.Group("type-things-story")
	app.Get("/get-type-things-story", middleware.AuthorizationRequired(), ControllerType_things_story.GetAllHandler)
	app.Get("/get-by-id/:type_things_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.GetById)
	app.Post("/create-type-things-story", middleware.AuthorizationRequired(), ControllerType_things_story.Create)
	app.Patch("/update-type-things-story/:type_things_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.UpdateType_things_story)
	app.Delete("/delete-type-things-story/:type_things_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.DeleteType_things_story)
}
