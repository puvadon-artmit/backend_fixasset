package branchRoutes

import (
	"github.com/gofiber/fiber/v2"
	BranchController "github.com/puvadon-artmit/gofiber-template/api/branch/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupBranchRoutes(router fiber.Router) {
	app := router.Group("branch")
	app.Get("/get-all", middleware.AuthorizationRequired(), BranchController.GetAll)
	app.Get("/get-coun-branch", middleware.AuthorizationRequired(), BranchController.GetBranch)
	app.Get("/get-by-id/:branch_id", middleware.AuthorizationRequired(), BranchController.GetBybranchId)
	app.Post("/create", middleware.AuthorizationRequired(), BranchController.Create)

	app.Patch("/update/:branch_id", middleware.AuthorizationRequired(), BranchController.Updatebranch)
	app.Delete("/delete/:branch_id", middleware.AuthorizationRequired(), BranchController.Deletebranch)

}
