package Main_Category_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_Category_story "github.com/puvadon-artmit/gofiber-template/api/main_category_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMain_Category_storyRoutes(router fiber.Router) {
	app := router.Group("main_category_story")
	app.Get("/get-all", middleware.AuthorizationRequired(), ControllerMain_Category_story.GetAllHandler)
	app.Get("/get-main-id/:id", middleware.AuthorizationRequired(), ControllerMain_Category_story.GetById)
	app.Get("/get-photo-id/:main_category_story_id", middleware.AuthorizationRequired(), ControllerMain_Category_story.GetByIdHandler)
	app.Post("/create", middleware.AuthorizationRequired(), ControllerMain_Category_story.Create)
	app.Patch("/update/:main_category_story_id", middleware.AuthorizationRequired(), ControllerMain_Category_story.UpdateMain_Category_story)
	app.Delete("/delete/:main_category_story_id", middleware.AuthorizationRequired(), ControllerMain_Category_story.DeleteMainCategory)
}
