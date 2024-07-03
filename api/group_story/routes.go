package Group_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerGroup_Story "github.com/puvadon-artmit/gofiber-template/api/group_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupGroup_StoryRoutes(router fiber.Router) {
	app := router.Group("group-story")
	app.Get("/get-group-story", middleware.AuthorizationRequired(), ControllerGroup_Story.GetAllHandler)
	app.Get("/get-by-id/:group_story_id", middleware.AuthorizationRequired(), ControllerGroup_Story.GetById)
	app.Post("/create-group-story", ControllerGroup_Story.Create)
	app.Patch("/update-group-story/:group_story_id", middleware.AuthorizationRequired(), ControllerGroup_Story.UpdateGroup_Story)
	app.Delete("/delete-group-story/:group_story_id", middleware.AuthorizationRequired(), ControllerGroup_Story.DeleteGroup_Story)
}
