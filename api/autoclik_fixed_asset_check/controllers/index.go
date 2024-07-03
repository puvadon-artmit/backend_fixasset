package autoclik_fixed_asset_checkController

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	// Fetch the asset check data from the database
	value, err := services.GetById(c.Params("autoclik_fixed_asset_check_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	// Fetch the photo data associated with the asset check from the database
	var faAllPhotos []model.Autoclik_Fixed_Asset_Photos_check
	tx := database.DB.Table("autoclik_fa_all_photos").
		Select("autoclik_fa_all_photos.autoclik_fixed_asset_check_id, autoclik_fixed_asset_photos_checks.*").
		Joins("left join autoclik_fixed_asset_photos_checks on autoclik_fa_all_photos.autoclik_fixed_asset_photos_check_id = autoclik_fixed_asset_photos_checks.autoclik_fixed_asset_photos_check_id").
		Where("autoclik_fa_all_photos.autoclik_fixed_asset_check_id = ?", value.Autoclik_Fixed_Asset_CheckID).
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
			"autoclik_fixed_asset_check": value,
			"autoclik_fa_all_photos":     faAllPhotos,
		},
	})
}

// func GetBySignedURLHandler(c *fiber.Ctx) error {
// 	photosCheckID := c.Params("autoclik_fixed_asset_check_id")

// 	autoclikCheck, err := services.GetById(photosCheckID)
// 	if err != nil {
// 		return c.Status(404).SendString("")
// 	}

// 	if autoclikCheck != nil && len(autoclikCheck.Autoclik_Fixed_Asset_Photos_check) > 0 {
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

// 		for i, photo := range autoclikCheck.Autoclik_Fixed_Asset_Photos_check {
// 			presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
// 			if err != nil {
// 				log.Println(err)
// 				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
// 			}

// 			autoclikCheck.Autoclik_Fixed_Asset_Photos_check[i].Photos = presignedURL.String()
// 		}
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": autoclikCheck,
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

func UpdateAutoclik_Fixed_Asset_CheckHandler(c *fiber.Ctx) error {
	id := c.Params("autoclik_fixed_asset_check_id")
	var updatedItem model.Autoclik_Fixed_Asset_Check
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := services.UpdateAutoclik_Fixed_Asset_CheckByID(id, updatedItem); err != nil {
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

func GetByAutoclikRound_CountHandlerlatest(c *fiber.Ctx) error {
	Round_CountID := c.Params("autoclik_fixed_asset_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "autoclik_fixed_asset_round_count_id is required",
		})
	}

	assetCheck, err := services.GetAutoclikCheckByautoclik_round_count_id(Round_CountID)
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

func GetByItem_noAndAutoclik_Round_CountIDHandler(c *fiber.Ctx) error {
	Item_no := c.Query("no")
	Autoclik_Round_CountID := c.Query("autoclik_fixed_asset_round_count_id")
	autoclik, err := services.GetByItem_noAndAutoclik_Round_CountID(Item_no, Autoclik_Round_CountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": autoclik,
	})
}

// func CreateNew(c *fiber.Ctx) error {
// 	var autoclik_fixed_asset_check model.Autoclik_Fixed_Asset_Check
// 	if err := c.BodyParser(&autoclik_fixed_asset_check); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	createdAsset, err := services.CreateNewAutoclik_Fixed_Asset_Check(autoclik_fixed_asset_check)
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
		Autoclik_Fixed_Asset_Check model.Autoclik_Fixed_Asset_Check
		Autoclik_Photos            []model.Autoclik_Fixed_Asset_Photos_check `json:"autoclik_fa_all_photos"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdAsset, err := services.CreateNewAutoclik_Fixed_Asset_Check(request.Autoclik_Fixed_Asset_Check, request.Autoclik_Photos)
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

func GetByAutoclikRound_CountHandler(c *fiber.Ctx) error {
	value, err := services.GetByAutoclikRound_CountID(c.Params("autoclik_fixed_asset_round_count_id"))
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
	Round_CountID := c.Params("autoclik_fixed_asset_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "round_count_id is required",
		})
	}

	assetCheck, err := services.GetAutoclikCheckByFixedAssetID(Round_CountID)
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

func GetByAutoclikRound_CountAllHandler(c *fiber.Ctx) error {
	value, err := services.GetByAutoclikRound_CountAllID(c.Params("autoclik_fixed_asset_round_count_id"))
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
