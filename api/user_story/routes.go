package User_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerType_things_story "github.com/puvadon-artmit/gofiber-template/api/user_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupUser_storyRoutes(router fiber.Router) {
	app := router.Group("user_story")
	app.Get("/get-user-story", middleware.AuthorizationRequired(), ControllerType_things_story.GetAllHandler)
	app.Get("/get-by-id/:user_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.GetById)
	app.Post("/create-user-story", middleware.AuthorizationRequired(), ControllerType_things_story.Create)
	app.Patch("/update-user-story/:user_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.UpdateUser_story)
	app.Delete("/delete-user-story/:user_story_id", middleware.AuthorizationRequired(), ControllerType_things_story.DeleteUser_story)
}
