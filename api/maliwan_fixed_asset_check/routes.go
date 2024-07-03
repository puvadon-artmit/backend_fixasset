package Maliwan_Fixed_Asset_CheckRoutes

import (
	"github.com/gofiber/fiber/v2"
	maliwan_fixed_asset_checkController "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_check/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupMaliwan_Fixed_Asset_CheckRoutes(router fiber.Router) {
	asset_fixed_asset_checkGroup := router.Group("/maliwan-fixed-asset-check")
	asset_fixed_asset_checkGroup.Get("/get-maliwan-fixed-asset-check", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetAll)
	asset_fixed_asset_checkGroup.Get("/get-by-id/:maliwan_fixed_asset_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetById)
	asset_fixed_asset_checkGroup.Get("/get-signed-url/:maliwan_fixed_asset_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetById)
	// asset_fixed_asset_checkGroup.Get("/get-round-count/:round_count_id", asset_fixed_asset_checkController.GetByRound_CountHandler)
	asset_fixed_asset_checkGroup.Get("/get-round-id/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetByMaliwanRound_CountHandler)

	asset_fixed_asset_checkGroup.Get("/get-latest-round-id/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetByLatestroundId)

	asset_fixed_asset_checkGroup.Get("/get-all-round-id/:maliwan_fixed_asset_round_count_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetByMaliwanRound_CountAllHandler)
	// asset_fixed_asset_checkGroup.Get("/get-round-latest/:maliwan_round_count_id", middleware.AuthorizationRequired(), asset_fixed_asset_checkController.GetByMaliwanRound_CountHandlerlatest)
	// asset_fixed_asset_checkGroup.Get("/get-round-latest", asset_fixed_asset_checkController.GetAllHandler)
	asset_fixed_asset_checkGroup.Post("/create-maliwan-fixed-asset-check", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.CreateNew)
	asset_fixed_asset_checkGroup.Patch("/update/:asset_fixed_asset_check_id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.UpdateMaliwan_Fixed_Asset_CheckHandler)

	asset_fixed_asset_checkGroup.Get("/filter-no-and-count-id", middleware.AuthorizationRequired(), maliwan_fixed_asset_checkController.GetByItem_noAndMaliwan_Round_CountIDHandler)
}
