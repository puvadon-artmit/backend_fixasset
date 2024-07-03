package ControllerMain_Category

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	ServicesMain_Category "github.com/puvadon-artmit/gofiber-template/api/main_category/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMain_Category.GetById(c.Params("id"))
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

func GetMain_Category(c *fiber.Ctx) error {
	count, err := ServicesMain_Category.CountMain_Category()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesMain_Category.GetAllRecords()
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
		if records[i].Main_Category_Photo == "" {
			// log.Println("Main_Category_Photo is empty for record:", records[i].Main_Category_ID)
			continue
		}
		imageURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), records[i].Main_Category_Photo, expiry, nil)
		if err != nil {
			log.Println("Error generating pre-signed URL for record:", records[i].Main_Category_ID, "Error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate pre-signed URL",
				"error":   err,
			})
		}
		records[i].Main_Category_Photo = imageURL.String()
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Create(c *fiber.Ctx) error {
	asset_count := new(model.Main_Category)
	err := c.BodyParser(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(asset_count)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMain_Category.CreateNewMain_Category(*asset_count)
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

func UpdateMain_Category(c *fiber.Ctx) error {
	id := c.Params("main_category_id")

	updatedType := new(model.Main_Category)
	err := c.BodyParser(updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedType)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMain_Category.UpdateMain_Category(id, *updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedStatus,
	})
}

func DeleteMainCategory(c *fiber.Ctx) error {
	Main_Category_ID := c.Params("main_category_id")

	err := ServicesMain_Category.DeleteMainCategoryByID(Main_Category_ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Main_Category deleted successfully",
	})
}

func GetByIdHandler(c *fiber.Ctx) error {
	Main_Category_ID := c.Params("main_category_id")

	mainCategory, err := ServicesMain_Category.GetById(Main_Category_ID)
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

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), mainCategory.Main_Category_Photo, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	mainCategory.Main_Category_Photo = presignedURL.String()

	return c.JSON(mainCategory)
}

// -------------------------------- เวอร์ชั่นใหม่ -------------------------------------------------------------

func GetByIdMain_CategoryHandler(c *fiber.Ctx) error {
	Main_Category_ID := c.Params("main_category_id")

	item, err := ServicesMain_Category.GetById(Main_Category_ID)
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

	var response map[string]interface{}

	if item.Main_Category_Photo != "" {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), item.Main_Category_Photo, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		// แปลง item ให้เป็น map[string]interface{}
		main_categoryMap := map[string]interface{}{
			"main_category_id":       item.Main_Category_ID,
			"code_id":                item.Code_ID,
			"main_name":              item.Main_name,
			"main_category_photo":    item.Main_Category_Photo,
			"created_at":             item.CreatedAt,
			"updated_at":             item.UpdatedAt,
			"DeletedAt":              item.DeletedAt,
			"main_category_imageurl": presignedURL.String(),
		}

		response = map[string]interface{}{
			"main_category": main_categoryMap,
		}
	} else {
		// แปลง item ให้เป็น map[string]interface{}
		main_categoryMap := map[string]interface{}{
			"main_category_id":    item.Main_Category_ID,
			"code_id":             item.Code_ID,
			"main_name":           item.Main_name,
			"main_category_photo": item.Main_Category_Photo,
			"created_at":          item.CreatedAt,
			"updated_at":          item.UpdatedAt,
			"DeletedAt":           item.DeletedAt,
		}

		response = map[string]interface{}{
			"main_category": main_categoryMap,
		}
	}

	return c.JSON(response)
}
