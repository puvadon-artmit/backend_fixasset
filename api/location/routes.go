package LocationRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/location/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupLocationRoutes(router fiber.Router) {
	app := router.Group("location")
	app.Get("/get-location", middleware.AuthorizationRequired(), ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:location_id", middleware.AuthorizationRequired(), ControllerTypeplan.GetById)
	app.Post("/create-location", middleware.AuthorizationRequired(), ControllerTypeplan.Create)
	app.Patch("/update-location/:location_id", middleware.AuthorizationRequired(), ControllerTypeplan.UpdateLocation)
	app.Delete("/delete-location/:controller_id", middleware.AuthorizationRequired(), ControllerTypeplan.DeleteLocation)
}
