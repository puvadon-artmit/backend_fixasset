package maliwan_checkController

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/maliwan_check/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("maliwan_check_id"))
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
	Round_CountID := c.Params("maliwan_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "maliwan_round_count_id is required",
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

func GetByNoAndMaliwan_Round_CountIDHandler(c *fiber.Ctx) error {
	No := c.Query("no")
	Maliwan_Round_CountID := c.Query("maliwan_round_count_id")
	maliwan, err := services.GetByNoAndMaliwan_Round_CountID(No, Maliwan_Round_CountID)
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

func UpdateMaliwan_checkHandler(c *fiber.Ctx) error {
	id := c.Params("maliwan_check_id")
	var updatedItem model.Maliwan_check
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := services.UpdateMaliwan_checkByID(id, updatedItem); err != nil {
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
	Round_CountID := c.Params("autoclik_round_count_id")
	if Round_CountID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "autoclik_round_count_id is required",
		})
	}

	assetCheck, err := services.GetMaliwan_checkByautoclik_round_count_id(Round_CountID)
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

// func CreateNew(c *fiber.Ctx) error {
// 	var maliwan_check model.Maliwan_check
// 	if err := c.BodyParser(&maliwan_check); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	createdAsset, err := services.CreateNewMaliwan_check(maliwan_check)
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
	var maliwan_check model.Maliwan_check
	if err := c.BodyParser(&maliwan_check); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdAsset, err := services.CreateNewMaliwan_check(maliwan_check, maliwan_check.Maliwan_Photos_check)
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
	value, err := services.GetByMaliwan_Round_CountID(c.Params("maliwan_round_count_id"))
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

func GetByIdHandler(c *fiber.Ctx) error {
	photosCheckID := c.Params("maliwan_check_id")

	// ดึงข้อมูล Maliwan_check จากฐานข้อมูล
	maliwanCheck, err := services.GetById(photosCheckID)
	if err != nil {
		return c.Status(404).SendString("")
	}

	// ตรวจสอบว่ามีภาพหรือไม่ และเลือกฟิลด์ที่เกี่ยวข้องกับภาพ
	if maliwanCheck != nil && len(maliwanCheck.Maliwan_Photos_check) > 0 {
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
		for i, photo := range maliwanCheck.Maliwan_Photos_check {
			// สร้าง pre-signed URL สำหรับภาพ
			presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
			}

			// กำหนดค่า photos ใหม่เป็น signed URL
			maliwanCheck.Maliwan_Photos_check[i].Photos = presignedURL.String()
		}
	}

	// ส่ง JSON กลับไปยัง client
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwanCheck,
	})
}
