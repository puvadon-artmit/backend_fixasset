package Signature_AutoclikRoutes

import (
	"github.com/gofiber/fiber/v2"
	Signature_AutoclikController "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupSignature_AutoclikeRoutes(router fiber.Router) {

	asset_checkGroup := router.Group("/signature-autocilk")

	asset_checkGroup.Get("/get-signature", middleware.AuthorizationRequired(), Signature_AutoclikController.GetAllHandler)
	asset_checkGroup.Get("/get-by-id/:autoclik_signature_id", middleware.AuthorizationRequired(), Signature_AutoclikController.GetByIdHandlers)
	asset_checkGroup.Get("/get-autoclik-count-id/:autoclik_count_id", middleware.AuthorizationRequired(), Signature_AutoclikController.GetByAutoclik_countIDHandlers)
	asset_checkGroup.Get("/get-asset-count-id/:asset_count_id", middleware.AuthorizationRequired(), Signature_AutoclikController.GetByAsset_countHandler)
	// asset_checkGroup.Post("/upload-signature", SignatureController.UploadFile)
	asset_checkGroup.Post("/upload-signature", middleware.AuthorizationRequired(), Signature_AutoclikController.Create)
	asset_checkGroup.Get("/get-image/:filename", middleware.AuthorizationRequired(), Signature_AutoclikController.GetImage)
	// asset_checkGroup.Patch("/update-signature-id/:signature_id", Signature_AutoclikController.UpdateSignatureHandler)

}
