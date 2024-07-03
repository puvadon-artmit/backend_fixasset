package Main_CategoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_Category "github.com/puvadon-artmit/gofiber-template/api/main_category/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMain_CategoryRoutes(router fiber.Router) {
	app := router.Group("main_category")
	app.Get("/get-main-category", middleware.AuthorizationRequired(), ControllerMain_Category.GetAllHandler)
	app.Get("/get-coun-main-category", middleware.AuthorizationRequired(), ControllerMain_Category.GetMain_Category)
	app.Get("/get-by-id/:main_category_id", middleware.AuthorizationRequired(), ControllerMain_Category.GetByIdMain_CategoryHandler)
	app.Get("/get-photo-id/:main_category_id", middleware.AuthorizationRequired(), ControllerMain_Category.GetByIdHandler)
	app.Post("/create", middleware.AuthorizationRequired(), ControllerMain_Category.Create)
	app.Patch("/update/:main_category_id", middleware.AuthorizationRequired(), ControllerMain_Category.UpdateMain_Category)
	app.Delete("/delete/:main_category_id", middleware.AuthorizationRequired(), ControllerMain_Category.DeleteMainCategory)
}
