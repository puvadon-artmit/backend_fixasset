package groupRoutes

import (
	"github.com/gofiber/fiber/v2"
	GroupController "github.com/puvadon-artmit/gofiber-template/api/group/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupGroupRoutes(router fiber.Router) {
	app := router.Group("group")
	app.Get("/get-group", middleware.AuthorizationRequired(), GroupController.GetAll)
	app.Get("/get-by-id/:group_id", middleware.AuthorizationRequired(), GroupController.GetById)

	app.Patch("/update/:group_id", middleware.AuthorizationRequired(), GroupController.UpdateGroupByID)

	app.Delete("/delete/:group_id", middleware.AuthorizationRequired(), GroupController.Deletegroup)

	app.Get("/get-count-group", middleware.AuthorizationRequired(), GroupController.GetCountGroup)
	app.Post("/create-group", middleware.AuthorizationRequired(), GroupController.Create)
}
