package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/typeplan/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupTypeplanRoutes(router fiber.Router) {
	app := router.Group("typeplan")
	app.Get("/get-typeplan", middleware.AuthorizationRequired(), ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:typeplan_id", middleware.AuthorizationRequired(), ControllerTypeplan.GetById)
	app.Post("/create-typeplan", middleware.AuthorizationRequired(), ControllerTypeplan.Create)
	app.Patch("/update-typeplan/:typeplan_id", middleware.AuthorizationRequired(), ControllerTypeplan.UpdateTypeplan)
	app.Delete("/delete-typeplan/:controller_id", middleware.AuthorizationRequired(), ControllerTypeplan.DeleteTypeplan)
}
