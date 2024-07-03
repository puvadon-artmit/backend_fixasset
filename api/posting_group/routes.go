package Posting_groupsRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerPosting_groups "github.com/puvadon-artmit/gofiber-template/api/posting_group/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupPosting_groupsRoutes(router fiber.Router) {
	app := router.Group("gen-product-posting-groups")
	app.Get("/clear-and-pull-posting-groups", middleware.AuthorizationRequired(), ControllerPosting_groups.ClearAndPull)
	app.Get("/autoclik-gen-posting-groups", middleware.AuthorizationRequired(), ControllerPosting_groups.GetAllHandler)
	app.Get("/maliwan-gen-posting-groups", middleware.AuthorizationRequired(), ControllerPosting_groups.GetPosting_group_Maliwan)

}
