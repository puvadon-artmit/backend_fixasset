package Branch_AutoclikRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerBranch_Autoclik "github.com/puvadon-artmit/gofiber-template/api/branch_autoclik/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupBranch_AutoclikRoutes(router fiber.Router) {
	app := router.Group("branch-autoclik")
	app.Get("/get-branch-autoclik", middleware.AuthorizationRequired(), ControllerBranch_Autoclik.GetAllHandler)
	app.Get("/get-by-id/:branchautoclik_id", middleware.AuthorizationRequired(), ControllerBranch_Autoclik.GetById)
	app.Post("/create-branch-autoclik", middleware.AuthorizationRequired(), ControllerBranch_Autoclik.Create)
	app.Patch("/update-branch-autoclik/:branchautoclik_id", middleware.AuthorizationRequired(), ControllerBranch_Autoclik.UpdateBranch_Autoclik)
	app.Delete("/delete-branch-autoclik/:branchautoclik_id", middleware.AuthorizationRequired(), ControllerBranch_Autoclik.DeleteBranch_Autoclik)
}
