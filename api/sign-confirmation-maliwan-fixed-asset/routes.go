package Signature_MaliwanRoutes

import (
	"github.com/gofiber/fiber/v2"
	Signature_Maliwan_Fixed_AssetController "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-maliwan-fixed-asset/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupSignature_Maliwan_Fixed_AssetRoutes(router fiber.Router) {

	asset_checkGroup := router.Group("/signature-maliwan-fixed-asset")

	asset_checkGroup.Get("/get-signature", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.GetAllHandler)
	asset_checkGroup.Get("/get-by-id/:maliwan_signature_id", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.GetByIdHandlers)
	asset_checkGroup.Get("/get-maliwan-fixed-asset-count-id/:maliwan_fixed_asset_count_id", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.GetByMaliwan_countIDHandlers)
	asset_checkGroup.Get("/get-asset-count-id/:asset_count_id", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.GetByAsset_countHandler)
	// asset_checkGroup.Post("/upload-signature", SignatureController.UploadFile)
	asset_checkGroup.Post("/upload-signature", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.Create)
	asset_checkGroup.Get("/get-image/:filename", middleware.AuthorizationRequired(), Signature_Maliwan_Fixed_AssetController.GetImage)
	// asset_checkGroup.Patch("/update-signature-id/:signature_id", Signature_Maliwan_Fixed_AssetController.UpdateSignatureHandler)

}
