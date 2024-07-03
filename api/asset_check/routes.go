package Asset_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	asset_checkController "github.com/puvadon-artmit/gofiber-template/api/asset_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAsset_checkRoutes(router fiber.Router) {
	asset_checkGroup := router.Group("/asset_check")
	asset_checkGroup.Get("/get-asset-check", middleware.AuthorizationRequired(), asset_checkController.GetAll)
	asset_checkGroup.Get("/get-by-id/:asset_check_id", middleware.AuthorizationRequired(), asset_checkController.GetById)
	asset_checkGroup.Get("/get-signed-url/:asset_check_id", middleware.AuthorizationRequired(), asset_checkController.GetBySignedURLHandler)
	asset_checkGroup.Get("/get-round-count/:round_count_id", middleware.AuthorizationRequired(), asset_checkController.GetByRound_CountHandler)
	asset_checkGroup.Get("/get-round-latest/:round_count_id", middleware.AuthorizationRequired(), asset_checkController.GetByProperty_code_Id)
	// asset_checkGroup.Get("/get-round-latest", asset_checkController.GetAllHandler)
	asset_checkGroup.Post("/upload-asset", middleware.AuthorizationRequired(), asset_checkController.CreateNew)
	asset_checkGroup.Patch("/update/:asset_check_id", middleware.AuthorizationRequired(), asset_checkController.UpdateAsset_checkHandler)

	asset_checkGroup.Get("/filter-property-and-count-id", middleware.AuthorizationRequired(), asset_checkController.GetByProperty_codeAndRound_CountIDHandler)

}
