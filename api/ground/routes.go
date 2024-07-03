package item_modelRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	GroundControllers "github.com/puvadon-artmit/gofiber-template/api/ground/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupGroundRoutes(router fiber.Router) {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Static("/uploads", "./uploads")

	categoryGroup := router.Group("/ground")

	categoryGroup.Get("/get-ground", middleware.AuthorizationRequired(), GroundControllers.GetAll)
	categoryGroup.Delete("/delete/:ground_id", middleware.AuthorizationRequired(), GroundControllers.DeleteGroundIDHandler)
	categoryGroup.Get("/get-count-ground", middleware.AuthorizationRequired(), GroundControllers.GetCountGround)
	categoryGroup.Get("/get-by-id/:ground_id", middleware.AuthorizationRequired(), GroundControllers.GetByIdgroundHandler)
	categoryGroup.Post("/upload", middleware.AuthorizationRequired(), GroundControllers.Create)
	categoryGroup.Get("/get-image/:filename", GroundControllers.GetImage)
	categoryGroup.Patch("/update/:ground_id", GroundControllers.UpdateGroundID)
}
