package Maliwan_counts_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_counts_story "github.com/puvadon-artmit/gofiber-template/api/maliwan_counts_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_counts_storyRoutes(router fiber.Router) {
	app := router.Group("maliwan_counts_story")
	app.Get("/get-maliwan-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_counts_story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_counts_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_story.GetById)
	app.Post("/create-maliwan-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_counts_story.Create)
	app.Patch("/update-maliwan-count-story/:maliwan_counts_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_story.UpdateMaliwan_counts_story)
	app.Delete("/delete-maliwan-count-story/:maliwan_counts_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_story.DeleteMaliwan_counts_story)
}
