package typeRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerType "github.com/puvadon-artmit/gofiber-template/api/type_things/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupTypeRoutes(router fiber.Router) {
	app := router.Group("type")
	app.Get("/get-type", middleware.AuthorizationRequired(), ControllerType.GetAllHandler)
	app.Get("/get-by-id/:type_id", middleware.AuthorizationRequired(), ControllerType.GetById)
	app.Patch("/update/:type_id", ControllerType.UpdateTypeIDHandler)
	app.Delete("/delete/:type_id", middleware.AuthorizationRequired(), ControllerType.DeleteTypeIDHandler)
	app.Get("/get-count-type", middleware.AuthorizationRequired(), ControllerType.GetCountType)
	app.Post("/create-type", middleware.AuthorizationRequired(), ControllerType.Create)
}
