package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_branch "github.com/puvadon-artmit/gofiber-template/api/main-branch/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMain_branchRoutes(router fiber.Router) {
	app := router.Group("main-branch")
	app.Get("/get-main-branch", middleware.AuthorizationRequired(), ControllerMain_branch.GetAllHandler)
	app.Get("/get-coun-main-branch", middleware.AuthorizationRequired(), ControllerMain_branch.GetMain_branch)
	app.Get("/get-by-id/:main_branch_id", middleware.AuthorizationRequired(), ControllerMain_branch.GetById)
	app.Post("/create-main-branch", middleware.AuthorizationRequired(), ControllerMain_branch.Create)
	app.Patch("/update-main-branch/:main_branch_id", middleware.AuthorizationRequired(), ControllerMain_branch.UpdateMain_branch)
	app.Delete("/delete-main-branch/:main_branch_id", middleware.AuthorizationRequired(), ControllerMain_branch.DeleteMain_branch)
}
