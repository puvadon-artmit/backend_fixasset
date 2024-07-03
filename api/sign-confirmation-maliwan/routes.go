package Signature_MaliwanRoutes

import (
	"github.com/gofiber/fiber/v2"
	Signature_MaliwanController "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-maliwan/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupSignature_MaliwanRoutes(router fiber.Router) {

	asset_checkGroup := router.Group("/signature-maliwan")

	asset_checkGroup.Get("/get-signature", middleware.AuthorizationRequired(), Signature_MaliwanController.GetAllHandler)
	asset_checkGroup.Get("/get-by-id/:maliwan_signature_id", middleware.AuthorizationRequired(), Signature_MaliwanController.GetByIdHandler)
	asset_checkGroup.Get("/get-asset-count-id/:asset_count_id", middleware.AuthorizationRequired(), Signature_MaliwanController.GetByAsset_countHandler)

	asset_checkGroup.Get("/get-maliwan-signature-id/:maliwan_count_id", middleware.AuthorizationRequired(), Signature_MaliwanController.GetBySignature_MaliwanIDHandlers)

	// asset_checkGroup.Post("/upload-signature", SignatureController.UploadFile)
	asset_checkGroup.Post("/upload-signature", middleware.AuthorizationRequired(), Signature_MaliwanController.Create)
	asset_checkGroup.Get("/get-image/:filename", middleware.AuthorizationRequired(), Signature_MaliwanController.GetImage)
	// asset_checkGroup.Patch("/update-signature-id/:signature_id", Signature_MaliwanController.UpdateSignatureHandler)

}
