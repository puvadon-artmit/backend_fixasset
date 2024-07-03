package typeRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerResponsible "github.com/puvadon-artmit/gofiber-template/api/responsible/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupResponsibleRoutes(router fiber.Router) {
	app := router.Group("responsible")
	app.Get("/get-responsible", middleware.AuthorizationRequired(), ControllerResponsible.GetAllHandler)
	app.Get("/get-count-responsible", middleware.AuthorizationRequired(), ControllerResponsible.GetCountResponsible)
	app.Get("/get-by-id/:responsible_id", middleware.AuthorizationRequired(), ControllerResponsible.GetById)
	app.Patch("/update/:responsible_id", middleware.AuthorizationRequired(), ControllerResponsible.UpdateResponsibleHandler)
	app.Delete("/delete/:responsible_id", middleware.AuthorizationRequired(), ControllerResponsible.DeleteResponsibleIDHandler)
	app.Post("/create-responsible", middleware.AuthorizationRequired(), ControllerResponsible.Create)
}
