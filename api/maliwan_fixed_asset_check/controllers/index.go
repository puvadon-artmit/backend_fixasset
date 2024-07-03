package maliwan_fixed_asset_checkController

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_check/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	// Fetch the asset check data from the database
	value, err := services.GetById(c.Params("maliwan_fixed_asset_check_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	// Fetch the photo data associated with the asset check from the database
	var faAllPhotos []model.Maliwan_Fixed_Asset_Photos_check
	tx := database.DB.Table("maliwan_fa_all_photos").
		Select("maliwan_fa_all_photos.maliwan_fixed_asset_check_id, maliwan_fixed_asset_photos_checks.*").
		Joins("left join maliwan_fixed_asset_photos_checks on maliwan_fa_all_photos.maliwan_fixed_asset_photos_check_id = maliwan_fixed_asset_photos_checks.maliwan_fixed_asset_photos_check_id").
		Where("maliwan_fa_all_photos.maliwan_fixed_asset_check_id = ?", value.Maliwan_Fixed_Asset_CheckID).
		Scan(&faAllPhotos)
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  tx.Error,
		})
	}

	// Generate signed URLs for the photos
	endpoint := viper.GetString("aws.endpoint")
	accessKeyID := viper.GetString("aws.accessKeyId")
	secretAccessKey := viper.GetString("aws.secretAccessKey")
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to initialize Minio client")
	}

	expiry := 1 * time.Minute // Set presigned URL expiry time

	for i, photo := range faAllPhotos {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		// Update the photo URL with the presigned URL
		faAllPhotos[i].Photos = presignedURL.String()
	}

	// Return the result with the updated photo URLs
	return c.JSON(fiber.Map{
		"status": "success",
		"result": map[string]interface{}{
			"maliwan_fixed_asset_check": value,
			"maliwan_fa_all_photos":     faAllPhotos,
		},
	})
}

// func GetBySignedURLHandler(c *fiber.Ctx) error {
// 	photosCheckID := c.Params("maliwan_fixed_asset_check_id")

// 	maliwanCheck, err := services.GetById(photosCheckID)
// 	if err != nil {
// 		return c.Status(404).SendString("")
// 	}

// 	if maliwanCheck != nil && len(maliwanCheck.Maliwan_Fixed_Asset_Photos_check) > 0 {
// 		endpoint := viper.GetString("aws.endpoint")
// 		accessKeyID := viper.GetString("aws.accessKeyId")
// 		secretAccessKey := viper.GetString("aws.secretAccessKey")
// 		useSSL := true

// 		minioClient, err := minio.New(endpoint, &minio.Options{
// 			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 			Secure: useSSL,
// 		})
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		expiry := 1 * time.Minute

// 		for i, photo := range maliwanCheck.Maliwan_Fixed_Asset_Photos_check {
// 			presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
// 			if err != nil {
// 				log.Println(err)
// 				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
// 			}

// 			maliwanCheck.Maliwan_Fixed_Asset_Photos_check[i].Photos = presignedURL.String()
// 		}
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": maliwanCheck,
// 	})
// }

// func GetByProperty_code_Id(c *fiber.Ctx) error {
// 	Round_CountID := c.Params("round_count_id")
// 	if Round_CountID == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "round_count_id is required",
// 		})
// 	}

// 	assetCheck, err := services.GetAssetCheckByAssetCheckID(Round_CountID)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": assetCheck,
// 	})
// }

func UpdateMaliwan_Fixed_Asset_CheckHandler(c *fiber.Ctx) error {
	id := c.Params("maliwan_fixed_asset_check_id")
	var updatedItem model.Maliwan_Fixed_Asset_Check
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := services.UpdateMaliwan_Fixed_Asset_CheckByID(id, updatedItem); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Asset check updated successfully",
	})
}

func GetAll(c *fiber.Ctx) error {
	value, err := services.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetByMaliwanRound_CountHandlerlatest(c *fiber.Ctx) error {
	Round_CountID := c.Params("maliwan_fixed_asset_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "maliwan_fixed_asset_round_count_id is required",
		})
	}

	assetCheck, err := services.GetMaliwanCheckBymaliwan_round_count_id(Round_CountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": assetCheck,
	})
}

func GetByItem_noAndMaliwan_Round_CountIDHandler(c *fiber.Ctx) error {
	Item_no := c.Query("no")
	Maliwan_Round_CountID := c.Query("maliwan_fixed_asset_round_count_id")
	maliwan, err := services.GetByItem_noAndMaliwan_Round_CountID(Item_no, Maliwan_Round_CountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwan,
	})
}

// func CreateNew(c *fiber.Ctx) error {
// 	var maliwan_fixed_asset_check model.Maliwan_Fixed_Asset_Check
// 	if err := c.BodyParser(&maliwan_fixed_asset_check); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	createdAsset, err := services.CreateNewMaliwan_Fixed_Asset_Check(maliwan_fixed_asset_check)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"data":   createdAsset,
// 	})
// }

func CreateNew(c *fiber.Ctx) error {
	var request struct {
		Maliwan_Fixed_Asset_Check model.Maliwan_Fixed_Asset_Check
		Maliwan_Photos            []model.Maliwan_Fixed_Asset_Photos_check `json:"maliwan_fa_all_photos"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdAsset, err := services.CreateNewMaliwan_Fixed_Asset_Check(request.Maliwan_Fixed_Asset_Check, request.Maliwan_Photos)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   createdAsset,
	})
}

func GetByMaliwanRound_CountHandler(c *fiber.Ctx) error {
	value, err := services.GetByMaliwanRound_CountID(c.Params("maliwan_fixed_asset_round_count_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetByLatestroundId(c *fiber.Ctx) error {
	Round_CountID := c.Params("maliwan_fixed_asset_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "round_count_id is required",
		})
	}

	assetCheck, err := services.GetMaliwanCheckByFixedAssetID(Round_CountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": assetCheck,
	})
}

func GetByMaliwanRound_CountAllHandler(c *fiber.Ctx) error {
	value, err := services.GetByMaliwanRound_CountAllID(c.Params("maliwan_fixed_asset_round_count_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}
