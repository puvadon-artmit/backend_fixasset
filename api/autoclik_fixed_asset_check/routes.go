package Asset_Fixed_Asset_CheckRoutes

import (
	"github.com/gofiber/fiber/v2"
	asset_fixed_asset_checkController "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAsset_Fixed_Asset_CheckRoutes(router fiber.Router) {
	asset_fixed_asset_checkGroup := router.Group("/autoclik-fixed-asset-check")
	asset_fixed_asset_checkGroup.Get("/get-autoclik-check", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetAll)
	asset_fixed_asset_checkGroup.Get("/get-by-id/:autoclik_fixed_asset_check_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetById)
	asset_fixed_asset_checkGroup.Get("/get-signed-url/:autoclik_fixed_asset_check_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetById)
	// asset_fixed_asset_checkGroup.Get("/get-round-count/:round_count_id", asset_fixed_asset_checkController.GetByRound_CountHandler)
	asset_fixed_asset_checkGroup.Get("/get-round-id/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByAutoclikRound_CountHandler)

	asset_fixed_asset_checkGroup.Get("/get-latest-round-id/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByLatestroundId)

	asset_fixed_asset_checkGroup.Get("/get-all-round-id/:autoclik_fixed_asset_round_count_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByAutoclikRound_CountAllHandler)
	// asset_fixed_asset_checkGroup.Get("/get-round-latest/:autoclik_round_count_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByAutoclikRound_CountHandlerlatest)
	// asset_fixed_asset_checkGroup.Get("/get-round-latest", asset_fixed_asset_checkController.GetAllHandler)
	asset_fixed_asset_checkGroup.Post("/create-autoclik-fixed-asset-check", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.CreateNew)
	asset_fixed_asset_checkGroup.Patch("/update/:asset_fixed_asset_check_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.UpdateAutoclik_Fixed_Asset_CheckHandler)

	asset_fixed_asset_checkGroup.Get("/filter-no-and-count-id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByItem_noAndAutoclik_Round_CountIDHandler)
}
