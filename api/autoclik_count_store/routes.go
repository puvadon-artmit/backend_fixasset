package Autoclik_count_StoreRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_count_Store "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_store/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_count_StoreRoutes(router fiber.Router) {
	app := router.Group("autoclik_count_store")
	app.Get("/get-autoclik-count-store", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.GetAllHandler)
	app.Get("/get-by-id/:autoclik_count_store_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.GetById)
	app.Get("/get-by-autoclik-count-id/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.GetByAutoclik_Count_StoreIDHandler)
	app.Post("/create-autoclik-count-store", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.Create)
	app.Patch("/update-autoclik-count-store/:autoclik_count_store_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.UpdateAutoclik_count_Store)
	app.Delete("/delete-autoclik-count-store/:autoclik_count_store_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.DeleteAutoclik_count_Store)
	app.Delete("/delete-all-autoclik-count-id/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_count_Store.DeleteAutoclik_Count_ID)
}
