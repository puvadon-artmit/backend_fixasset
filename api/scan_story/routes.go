package Scan_storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerScan_story "github.com/puvadon-artmit/gofiber-template/api/scan_story/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupScan_storyRoutes(router fiber.Router) {
	app := router.Group("scan-story")
	app.Get("/get-scan_story", middleware.AuthorizationRequired(), ControllerScan_story.GetAllHandler)
	app.Get("/get-by-id/:scan_story_id", middleware.AuthorizationRequired(), ControllerScan_story.GetById)
	app.Post("/create-scan-story", middleware.AuthorizationRequired(), ControllerScan_story.Create)
	app.Patch("/update-scan_story/:scan_story_id", middleware.AuthorizationRequired(), ControllerScan_story.UpdateScan_story)
	app.Delete("/delete-scan_story/:scan_story_id", middleware.AuthorizationRequired(), ControllerScan_story.DeleteScan_story)
}
