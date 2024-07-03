package counting_rightsRoutes

import (
	"github.com/gofiber/fiber/v2"
	Controllercounting_rights "github.com/puvadon-artmit/gofiber-template/api/counting_rights/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func Setupcounting_rightsRoutes(router fiber.Router) {
	app := router.Group("counting_rights")
	app.Get("/get-counting-rights", middleware.AuthorizationRequired(), Controllercounting_rights.GetAllHandler)
	app.Get("/get-by-id/:counting_rights_id", middleware.AuthorizationRequired(), Controllercounting_rights.GetById)
	app.Get("/get-user-id/:user_id", middleware.AuthorizationRequired(), Controllercounting_rights.GetUSer_IDHandler)
	app.Post("/create-counting-rights", middleware.AuthorizationRequired(), Controllercounting_rights.Create)
	app.Patch("/update-counting-rights/:counting_rights_id", middleware.AuthorizationRequired(), Controllercounting_rights.UpdateCounting_rights)
	app.Delete("/delete-counting-rights/:counting_rights_id", middleware.AuthorizationRequired(), Controllercounting_rights.DeleteCountingRights)
}
