package manufacturerRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerManufacturer "github.com/puvadon-artmit/gofiber-template/api/manufacturer/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupManufacturerRoutes(router fiber.Router) {
	app := router.Group("manufacturer")
	app.Get("/get-manufacturer", middleware.AuthorizationRequired(), ControllerManufacturer.GetAllHandler)
	app.Get("/get-count-manufacturer", middleware.AuthorizationRequired(), ControllerManufacturer.GetCountManufacturer)
	app.Get("/get-by-id/:manufacturer_id", middleware.AuthorizationRequired(), ControllerManufacturer.GetById)
	app.Post("/create-manufacturer", middleware.AuthorizationRequired(), ControllerManufacturer.Create)
	app.Patch("/update-manufacturer/:manufacturer_id", middleware.AuthorizationRequired(), ControllerManufacturer.UpdateManufacturer)
}
