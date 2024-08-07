package autoclik_counting_rightsRoutes

import (
	"github.com/gofiber/fiber/v2"
	Controllerautoclik_counting_rights "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_rights/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupaMaliwan_counting_rightsRoutes(router fiber.Router) {
	app := router.Group("maliwan-counting-rights")
	app.Get("/get-maliwan-counting-rights", middleware.AuthorizationRequired(), Controllerautoclik_counting_rights.GetAllHandler)
	app.Get("/get-by-id/:maliwan_counting_rights_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_rights.GetById)
	app.Get("/get-user-id/:user_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_rights.GetUSer_IDHandler)
	app.Post("/create-maliwan-counting-rights", middleware.AuthorizationRequired(), Controllerautoclik_counting_rights.Create)
	// app.Patch("/update-counting-rights/:autoclik_counting_rights_id", Controllerautoclik_counting_rights.UpdateAutoclik_Counting_Rights)
	app.Delete("/delete-counting-rights/:maliwan_counting_rights_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_rights.DeleteCountingRights)
}
