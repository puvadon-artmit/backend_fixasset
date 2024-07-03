package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Update_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_update_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Update_StoryRoutes(router fiber.Router) {
	app := router.Group("maliwan-update-story")
	app.Get("/get-maliwan-update-story", middleware.AuthorizationRequired(), ControllerMaliwan_Update_Story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_update_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Update_Story.GetById)
	app.Post("/create-maliwan-update-story", middleware.AuthorizationRequired(), ControllerMaliwan_Update_Story.Create)
	app.Patch("/update-maliwan-update-story/:maliwan_update_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Update_Story.UpdateInspection_story)
	app.Delete("/delete-maliwan-update-story/:maliwan_update_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Update_Story.DeleteInspection_story)
}
