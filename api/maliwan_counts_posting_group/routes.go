package Maliwan_counts_posting_groupRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMaliwan_counts_posting_group "github.com/puvadon-artmit/gofiber-template/api/maliwan_counts_posting_group/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_counts_posting_groupRoutes(router fiber.Router) {
	app := router.Group("maliwan-counts-item-category-code")
	app.Get("/get-maliwan-counts-item-category-code", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.GetAllHandler)
	app.Get("/get-by-id/:maliwan_counts_item_category_code_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.GetById)
	app.Get("/get-maliwan-count-id/:maliwan_count_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.GetMaliwan_count_IDHandler)
	app.Post("/create-maliwan-counts-item-category-code", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.Create)
	app.Patch("/update-maliwan-counts-item-category-code/:maliwan_counts_item_category_code_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.UpdateMaliwan_Counts_Gen_Posting_Group)
	app.Delete("/delete-maliwan-counts-item-category-code/:maliwan_counts_item_category_code_id", middleware.AuthorizationRequired(), ControllerMaliwan_counts_posting_group.DeleteMaliwan_Counts_Gen_Posting_Group)
}
