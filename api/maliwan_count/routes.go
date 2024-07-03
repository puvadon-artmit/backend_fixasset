package Maliwan_countRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_Count "github.com/puvadon-artmit/gofiber-template/api/maliwan_count/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_CountRoutes(router fiber.Router) {
	app := router.Group("maliwan_count")
	app.Get("/get-maliwan-count", middleware.AuthorizationRequired(), ControllerMaliwan_Count.GetAllHandler)
	app.Get("/get-by-id/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Count.GetById)
	app.Post("/create-maliwan-count", middleware.AuthorizationRequired(), ControllerMaliwan_Count.Create)
	app.Patch("/update-maliwan-count/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Count.UpdateMaliwan_count)
	app.Delete("/delete-maliwan-count/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_Count.DeleteMaliwan_count)
}
