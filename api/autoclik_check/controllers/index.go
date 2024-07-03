package autoclik_checkController

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/autoclik_check/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("autoclik_check_id"))
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

func GetBySignedURLHandler(c *fiber.Ctx) error {
	photosCheckID := c.Params("autoclik_check_id")

	// ดึงข้อมูล Maliwan_check จากฐานข้อมูล
	autoclikCheck, err := services.GetById(photosCheckID)
	if err != nil {
		return c.Status(404).SendString("")
	}

	// ตรวจสอบว่ามีภาพหรือไม่ และเลือกฟิลด์ที่เกี่ยวข้องกับภาพ
	if autoclikCheck != nil && len(autoclikCheck.Autoclik_Photos_check) > 0 {
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
		for i, photo := range autoclikCheck.Autoclik_Photos_check {
			// สร้าง pre-signed URL สำหรับภาพ
			presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), photo.Photos, expiry, nil)
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
			}

			// กำหนดค่า photos ใหม่เป็น signed URL
			autoclikCheck.Autoclik_Photos_check[i].Photos = presignedURL.String()
		}
	}

	// ส่ง JSON กลับไปยัง client
	return c.JSON(fiber.Map{
		"status": "success",
		"result": autoclikCheck,
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

func UpdateAutoclik_checkHandler(c *fiber.Ctx) error {
	id := c.Params("autoclik_check_id")
	var updatedItem model.Autoclik_check
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := services.UpdateAutoclik_checkByID(id, updatedItem); err != nil {
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
	Item_no := c.Query("item_no")
	Autoclik_Round_CountID := c.Query("autoclik_round_count_id")
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
// 	var autoclik_check model.Autoclik_check
// 	if err := c.BodyParser(&autoclik_check); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	createdAsset, err := services.CreateNewAutoclik_check(autoclik_check)
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
	var autoclik_check model.Autoclik_check
	if err := c.BodyParser(&autoclik_check); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdAsset, err := services.CreateNewAutoclik_check(autoclik_check, autoclik_check.Autoclik_Photos_check)
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
	value, err := services.GetByAutoclikRound_CountID(c.Params("autoclik_round_count_id"))
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
