package Autoclik_Round_Count_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Round_Count_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_round_count_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Round_Count_StoryRoutes(router fiber.Router) {
	app := router.Group("autoclik_round_count_story")
	app.Get("/get-autoclik-round-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count_Story.GetAllHandler)
	app.Get("/get-by-id/:autoclik_round_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count_Story.GetById)
	app.Post("/create-autoclik-round-count-story", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count_Story.Create)
	app.Patch("/update-autoclik-round-count-story/:autoclik_round_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count_Story.UpdateAutoclik_Round_Count_Story)
	app.Delete("/delete-autoclik-round-count-story/:autoclik_round_count_story_id", middleware.AuthorizationRequired(), ControllerAutoclik_Round_Count_Story.DeleteAutoclik_Round_Count_Story)
}
