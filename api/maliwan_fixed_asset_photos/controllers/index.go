package maliwan_fixed_asset_photos_checkController

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_photos/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

// func GetPhotoById(c *fiber.Ctx) error {
// 	value, err := services.GetById(c.Params("maliwan_fixed_asset_photos_check_id"))
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err,
// 		})
// 	}

// 	responseData := fiber.Map{
// 		"status": "success",
// 		"result": value,
// 	}
// 	if value.Photos != "" {
// 		imagePath := "./uploads/" + value.Photos
// 		return c.SendFile(imagePath)
// 	}

// 	return c.JSON(responseData)
// }

func GetAll(c *fiber.Ctx) error {
	value, err := services.GetAll()
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

func SaveFileDataToDB(filename string) error {

	maliwan_fixed_asset_photos_checkID := uuid.New().String()
	maliwan_fixed_asset_photos_check := model.Maliwan_Fixed_Asset_Photos_check{
		Maliwan_Fixed_Asset_Photos_checkID: maliwan_fixed_asset_photos_checkID,
		Photos:                             filename,
	}

	if err := database.DB.Create(&maliwan_fixed_asset_photos_check).Error; err != nil {
		return err
	}

	return nil
}

func Create(c *fiber.Ctx) error {
	maliwan_fixed_asset_photos_check := new(model.Maliwan_Fixed_Asset_Photos_check)
	err := c.BodyParser(maliwan_fixed_asset_photos_check)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_fixed_asset_photos_check)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := services.CreateNewMaliwan_Fixed_Asset_Photos_check(*maliwan_fixed_asset_photos_check)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdStatus,
	})
}

func GetAllHandler(c *fiber.Ctx) error {
	records, err := services.GetAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func GetByIdHandler(c *fiber.Ctx) error {
	photosCheckID := c.Params("maliwan_fixed_asset_photos_check_id")

	photosURL, err := services.GetById(photosCheckID)
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

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photosURL, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

func DeletePhotos(c *fiber.Ctx) error {
	photosCheckID := c.Params("maliwan_fixed_asset_photos_check_id")

	photosCheck, err := services.GetByPhotoId(photosCheckID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteMaliwan_Fixed_Asset_Photos_check(photosCheck)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Photos record deleted successfully",
	})
}
