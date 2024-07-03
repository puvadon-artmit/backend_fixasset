package Signature_AutoclikRoutes

import (
	"github.com/gofiber/fiber/v2"
	Signature_Autoclik_Fixed_AssetController "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik-fixed-asset/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupSignature_Autoclike_Fixed_AssetRoutes(router fiber.Router) {

	asset_checkGroup := router.Group("/signature-autocilk-fixed-asset")

	asset_checkGroup.Get("/get-signature", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.GetAllHandler)
	asset_checkGroup.Get("/get-by-id/:autoclik_signature_id", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.GetByIdHandlers)
	asset_checkGroup.Get("/get-autoclik-fixed-asset-count-id/:autoclik_fixed_asset_count_id", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.GetByAutoclik_countIDHandlers)
	asset_checkGroup.Get("/get-asset-count-id/:asset_count_id", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.GetByAsset_countHandler)
	// asset_checkGroup.Post("/upload-signature", SignatureController.UploadFile)
	asset_checkGroup.Post("/upload-signature", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.Create)
	asset_checkGroup.Get("/get-image/:filename", middleware.AuthorizationRequired(), Signature_Autoclik_Fixed_AssetController.GetImage)
	// asset_checkGroup.Patch("/update-signature-id/:signature_id", Signature_Autoclik_Fixed_AssetController.UpdateSignatureHandler)

}
