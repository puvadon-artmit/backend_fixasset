package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Round_Count_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_round_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Round_Count_StoryRoutes(router fiber.Router) {
	app := router.Group("maliwan-round-count-story")
	app.Get("/get-maliwan-round-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count_Story.GetAllHandler)
	app.Get("/get-by-id/:maliwan_round_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count_Story.GetById)
	app.Post("/create-maliwan-round-count-story", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count_Story.Create)
	app.Patch("/update-maliwan-round-count-story/:maliwan_round_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count_Story.UpdateInspection_story)
	app.Delete("/delete-maliwan-round-count-story/:maliwan_round_count_story_id", middleware.AuthorizationRequired(), ControllerMaliwan_Round_Count_Story.DeleteInspection_story)
}
