package asset_checkController

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/asset_check/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("asset_check_id"))
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

func GetByProperty_code_Id(c *fiber.Ctx) error {
	Round_CountID := c.Params("round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "round_count_id is required",
		})
	}

	assetCheck, err := services.GetAssetCheckByAssetCheckID(Round_CountID)
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

func UpdateAsset_checkHandler(c *fiber.Ctx) error {
	id := c.Params("asset_check_id")
	var updatedItem model.Asset_check
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := services.UpdateAsset_checkByID(id, updatedItem); err != nil {
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

// func CreateNew(c *fiber.Ctx) error {
// 	var asset_check model.Asset_check
// 	if err := c.BodyParser(&asset_check); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	createdAsset, err := services.CreateNewAsset_check(asset_check)
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

func GetBySignedURLHandler(c *fiber.Ctx) error {
	photosCheckID := c.Params("asset_check_id")

	// ดึงข้อมูล Maliwan_check จากฐานข้อมูล
	maliwanCheck, err := services.GetById(photosCheckID)
	if err != nil {
		return c.Status(404).SendString("")
	}

	// ตรวจสอบว่ามีภาพหรือไม่ และเลือกฟิลด์ที่เกี่ยวข้องกับภาพ
	if maliwanCheck != nil && len(maliwanCheck.Photos_check) > 0 {
		// เชื่อมต่อกับ Minio client
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

		// กำหนดเวลาในการใช้งาน pre-signed URL
		expiry := 1 * time.Minute

		// สร้าง signed URL สำหรับทุกภาพในรายการ Maliwan_Photos_check
		for i, photo := range maliwanCheck.Photos_check {
			// สร้าง pre-signed URL สำหรับภาพ
			presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
			}

			// กำหนดค่า photos ใหม่เป็น signed URL
			maliwanCheck.Photos_check[i].Photos = presignedURL.String()
		}
	}

	// ส่ง JSON กลับไปยัง client
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwanCheck,
	})
}

func GetByProperty_codeAndRound_CountIDHandler(c *fiber.Ctx) error {
	No := c.Query("property_code")
	Maliwan_Round_CountID := c.Query("round_count_id")
	maliwan, err := services.GetByProperty_codeAndRound_CountID(No, Maliwan_Round_CountID)
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

func CreateNew(c *fiber.Ctx) error {
	var asset_check model.Asset_check
	if err := c.BodyParser(&asset_check); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdAsset, err := services.CreateNewAsset_check(asset_check, asset_check.Photos_check)
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

func GetByRound_CountHandler(c *fiber.Ctx) error {
	value, err := services.GetByRound_CountID(c.Params("round_count_id"))
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
