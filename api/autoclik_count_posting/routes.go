package Autoclik_Count_ProductRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_Count_Product "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_posting/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_Count_ProductRoutes(router fiber.Router) {
	app := router.Group("autoclik-count-product")
	app.Get("/get-autoclik-count-product", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.GetAllHandler)
	app.Get("/get-by-id/:autoclik_count_product_product_group_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.GetById)
	app.Get("/get-count-autoclik-id/:autoclik_count_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.GetAutoclik_countHandler)
	app.Post("/create-autoclik-count-product", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.Create)
	app.Patch("/update-autoclik-count-product/:autoclik_count_product_product_group_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.UpdateAutoclik_Count_Product_Group)
	app.Delete("/delete-autoclik-count-product/:autoclik_count_product_product_group_id", middleware.AuthorizationRequired(), ControllerAutoclik_Count_Product.DeleteAutoclik_Count_Product_Group)
}
