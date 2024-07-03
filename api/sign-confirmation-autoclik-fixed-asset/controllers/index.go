package asset_checkController

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik-fixed-asset/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("autoclik_fixed_asset_signature_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetByAutoclik_countIDHandlers(c *fiber.Ctx) error {
	Autoclik_countID := c.Params("autoclik_fixed_asset_count_id")

	// Validate UUID format
	if _, err := uuid.Parse(Autoclik_countID); err != nil {
		return c.Status(400).SendString("Invalid UUID format")
	}

	signatures, err := services.GetByAutoclik_countID(Autoclik_countID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

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
	}

	expiry := 1 * time.Minute

	var response []fiber.Map

	for _, signature := range signatures {
		if signature.Autoclik_Fixed_Asset_Signature == nil {
			continue
		}

		filename := *signature.Autoclik_Fixed_Asset_Signature

		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), filename, expiry, nil)
		if err != nil {
			log.Println(err)
			continue
		}

		response = append(response, fiber.Map{
			"autoclik_fixed_asset_signature_id": signature.Signature_Autoclik_Fixed_AssetID,
			"autoclik_fixed_asset_signature":    signature.Autoclik_Fixed_Asset_Signature,
			"user_id":                           signature.UserID,
			"user":                              signature.User,
			"created_at":                        signature.CreatedAt,
			"updated_at":                        signature.UpdatedAt,
			"deleted_at":                        signature.DeletedAt,
			"autoclik_fixed_asset_count_id":     signature.Autoclik_Fixed_Asset_CountID,
			"autoclik_fixed_asset_count":        signature.Autoclik_Fixed_Asset_Count,
			"presigned_url":                     presignedURL.String(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": response,
	})
}

func GetAll(c *fiber.Ctx) error {
	value, err := services.GetAllRecordsSignature()
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

func GetByIdHandlers(c *fiber.Ctx) error {
	SignatureID := c.Params("autoclik_fixed_asset_signature")

	signature, err := services.GetById(SignatureID)
	if err != nil {
		return c.Status(404).SendString("")
	}

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
	}

	expiry := 1 * time.Minute

	filename := *signature.Autoclik_Fixed_Asset_Signature

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), filename, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

func GetAllHandler(c *fiber.Ctx) error {
	records, err := services.GetAllRecordsSignature()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

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
	}

	expiry := 1 * time.Minute

	for i := range records {
		if records[i].Autoclik_Fixed_Asset_Signature == nil || *records[i].Autoclik_Fixed_Asset_Signature == "" {
			continue
		}
		imageURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *records[i].Autoclik_Fixed_Asset_Signature, expiry, nil)
		if err != nil {
			log.Println("Error generating pre-signed URL for record:", records[i].Signature_Autoclik_Fixed_AssetID, "Error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate pre-signed URL",
				"error":   err,
			})
		}
		imageURLStr := imageURL.String()
		records[i].Autoclik_Fixed_Asset_Signature = &imageURLStr
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Create(c *fiber.Ctx) error {
	signature := new(model.Signature_Autoclik_Fixed_Asset)
	err := c.BodyParser(signature)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(signature)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdSignature, err := services.CreateNewSignature_Autoclik_Fixed_Asset(*signature)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdSignature,
	})
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func GetByAsset_countHandler(c *fiber.Ctx) error {
	value, err := services.GetByAsset_countDB(c.Params("asset_count_id"))
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
