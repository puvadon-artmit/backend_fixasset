package Maliwan_check_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_check_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_check_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_check_StoryRoutes(router fiber.Router) {
	app := router.Group("maliwan_check_story")
	app.Get("/get-maliwan-check-story", middleware.AuthorizationRequired(), ControllerMaliwan_check_Story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_check_Story.GetById)
	app.Post("/create-maliwan-check-story", middleware.AuthorizationRequired(), ControllerMaliwan_check_Story.Create)
	app.Patch("/update-maliwan-check-story/:maliwan_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_check_Story.UpdateMaliwan_check_Story)
	app.Delete("/delete-maliwan-check-story/:maliwan_check_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_check_Story.DeleteMaliwan_check_Story)
}
