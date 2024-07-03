package Category_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerStoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupCategory_StoryRoutes(router fiber.Router) {
	app := router.Group("category-story")
	app.Get("/get-category-story", middleware.AuthorizationRequired(), ControllerStoryRoutes.GetAllHandler)
	app.Get("/get-by-id/:category_story_id", middleware.AuthorizationRequired(), ControllerStoryRoutes.GetById)
	app.Post("/create-category-story", middleware.AuthorizationRequired(), ControllerStoryRoutes.Create)
	app.Patch("/update-category-story/:category_story_id", middleware.AuthorizationRequired(), ControllerStoryRoutes.UpdateCategory_Story)
	app.Delete("/delete-category-story/:controller_id", middleware.AuthorizationRequired(), ControllerStoryRoutes.DeleteCategory_Story)
}
