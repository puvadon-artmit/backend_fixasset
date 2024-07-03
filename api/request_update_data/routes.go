package Request_update_dataRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerRequest_update_data "github.com/puvadon-artmit/gofiber-template/api/request_update_data/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupRequest_update_dataRoutes(router fiber.Router) {
	app := router.Group("request-update-data")
	app.Get("/get-request-update-data", middleware.AuthorizationRequired(), ControllerRequest_update_data.GetAllHandler)
	app.Get("/get-by-id/:request_update_data_id", middleware.AuthorizationRequired(), ControllerRequest_update_data.GetById)
	app.Post("/create-request-update-data", middleware.AuthorizationRequired(), ControllerRequest_update_data.Create)
	app.Patch("/update-request-update-data/:request_update_data_id", middleware.AuthorizationRequired(), ControllerRequest_update_data.UpdateRequest_update_data)
	app.Delete("/delete-request-update-data/:request_update_data_id", middleware.AuthorizationRequired(), ControllerRequest_update_data.DeleteRequest_update_data)
}
