package count_main_categoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/count_main_category/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func Setupcount_main_categoryRoutes(router fiber.Router) {
	app := router.Group("count-main-category")
	app.Get("/get-count-main-category", middleware.AuthorizationRequired(), ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:asset_count_main_category_id", middleware.AuthorizationRequired(), ControllerTypeplan.GetById)
	app.Get("/get-main-asset-count/:asset_count_id", middleware.AuthorizationRequired(), ControllerTypeplan.GetAsset_countHandler)
	app.Get("/get-main-asset-count-id/:asset_count_id", ControllerTypeplan.GetAsset_countHandler)
	app.Post("/create-count-main-category", middleware.AuthorizationRequired(), ControllerTypeplan.Create)
	app.Patch("/update-count-main-category/:asset_count_main_category_id", middleware.AuthorizationRequired(), ControllerTypeplan.Updatecount_main_category)
	app.Delete("/delete-count-main-category/:asset_count_main_category_id", middleware.AuthorizationRequired(), ControllerTypeplan.Deletecount_main_category)
}
