package Maliwan_count_StoreRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_count_Store "github.com/puvadon-artmit/gofiber-template/api/maliwan_count_store/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_count_StoreRoutes(router fiber.Router) {
	app := router.Group("maliwan_count_store")
	app.Get("/get-maliwan-count-store", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.GetAllHandler)
	app.Get("/get-by-id/:maliwan_count_store_id", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.GetById)
	app.Get("/get-by-maliwan-count-id/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.GetByMaliwan_Count_StoreIDHandler)
	app.Get("/get-item-no", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.GetByItem_No_Maliwan_BinHandler)
	app.Post("/create-maliwan-count-store", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.Create)
	app.Patch("/update-maliwan-count-store/:maliwan_count_store_id", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.UpdateMaliwan_count_Store)
	app.Delete("/delete-maliwan-count-store/:maliwan_count_store_id", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.DeleteMaliwan_count_Store)
	app.Delete("/delete-all-maliwan-count-id/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.DeleteMaliwan_Count_ID)
}

func SetupMaliwan_bin_dataRoutes(router fiber.Router) {
	app := router.Group("maliwan-bin-data")
	app.Get("/get-item-no", middleware.AuthorizationRequired(), ControllerMaliwan_count_Store.GetByItem_No_Maliwan_BinHandler)
}
