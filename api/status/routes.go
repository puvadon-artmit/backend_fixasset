package statusRoutes

import (
	"github.com/gofiber/fiber/v2"
	StatusController "github.com/puvadon-artmit/gofiber-template/api/status/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupStatusRoutes(router fiber.Router) {
	app := router.Group("status")
	app.Get("/get-status", middleware.AuthorizationRequired(), StatusController.GetAll)
	app.Get("/get-count-status", middleware.AuthorizationRequired(), StatusController.GetCountstatus)
	app.Get("/get-by-id/:status_id", middleware.AuthorizationRequired(), StatusController.GetById)
	app.Delete("/delete/:status_id", middleware.AuthorizationRequired(), StatusController.DeleteStatusIDHandler)
	app.Patch("/update/:status_id", middleware.AuthorizationRequired(), StatusController.UpdateStatusHandler)
	app.Post("/create-status", middleware.AuthorizationRequired(), StatusController.Create)
}
