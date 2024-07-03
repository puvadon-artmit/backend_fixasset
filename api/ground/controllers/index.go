package groundController

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/ground/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("ground_id"))
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

func UpdateItemModelHandler(c *fiber.Ctx) error {
	groundID := c.Params("ground_id")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var updatedItem model.Ground
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel, err := services.GetById(groundID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel.GroundName = updatedItem.GroundName
	existingItemModel.Comment = updatedItem.Comment
	existingItemModel.Location_code = updatedItem.Location_code
	existingItemModel.Building = updatedItem.Building
	existingItemModel.Floor = updatedItem.Floor

	if len(form.File["GroundImage"]) > 0 {
		file := form.File["GroundImage"][0]

		// Save the new image
		randomName := uuid.New().String()
		filename := randomName + "_" + file.Filename
		err := c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		existingItemModel.GroundImage = &filename
	}

	if err := services.UpdateGroundByID(groundID, *existingItemModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": existingItemModel,
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

func GetCountGround(c *fiber.Ctx) error {
	count, err := services.CountLatestRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

type GroundInput struct {
	GroundName    string `json:"ground_name"`
	Comment       string `json:"comment"`
	Location_code string `json:"location_code"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
	Room          string `json:"room"`
	GroundImage   string `json:"ground_image"`
	UserID        string `json:"user_id"`
}

type UploadForm struct {
	UserID        string `form:"user_id"`
	GroundName    string `form:"ground_name"`
	GroundImage   string `form:"ground_image"`
	Comment       string `form:"comment"`
	Location_code string `form:"product_number"`
	Building      string `form:"building"`
	Floor         string `form:"floor"`
	Room          string `form:"room"`
}

// ---------------------------------------------------------------------------------------------

func UploadFile(c *fiber.Ctx) error {
	form := new(UploadForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("ground_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	randomName := uuid.New().String()
	filename := randomName + "_" + file.Filename

	// บันทึกไฟล์ลงใน "./uploads/"
	err = c.SaveFile(file, "./uploads/"+filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	GroundID := uuid.New().String()
	GroundImageImageName := filename

	input := GroundInput{
		GroundName:    form.GroundName,
		GroundImage:   GroundImageImageName,
		UserID:        form.UserID,
		Comment:       form.Comment,
		Location_code: form.Location_code,
		Building:      form.Building,
		Room:          form.Room,
		Floor:         form.Floor,
	}

	ground := model.Ground{
		GroundID:      GroundID,
		GroundName:    &input.GroundName,
		GroundImage:   &GroundImageImageName,
		Comment:       &input.Comment,
		Location_code: &input.Location_code,
		Room:          &input.Room,
		Building:      &input.Building,
		Floor:         &input.Floor,
	}

	if err := database.DB.Create(&ground).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString(GroundID)
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func DeleteGroundById(c *fiber.Ctx) error {
	GroundID := c.Params("ground_id")

	ground, err := services.GetById(GroundID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteGroundPhoto(ground)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Ground record deleted successfully",
	})
}

func GetByIdHandler(c *fiber.Ctx) error {
	GroundID := c.Params("ground_id")

	GroundImageURL, err := services.GetById(GroundID)
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

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *GroundImageURL.GroundImage, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	*GroundImageURL.GroundImage = presignedURL.String()

	return c.JSON(GroundImageURL)
}

func Create(c *fiber.Ctx) error {
	ground := new(model.Ground)
	err := c.BodyParser(ground)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(ground)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := services.CreateNewGround(*ground)
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

func GetByIdgroundHandler(c *fiber.Ctx) error {
	GroundID := c.Params("ground_id")

	item, err := services.GetById(GroundID)
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

	if item.GroundImage != nil && *item.GroundImage != "" {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *item.GroundImage, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		// แปลง item ให้เป็น map[string]interface{}
		groundMap := map[string]interface{}{
			"ground_id":       item.GroundID,
			"ground_name":     item.GroundName,
			"location_code":   item.Location_code,
			"building":        item.Building,
			"floor":           item.Floor,
			"room":            item.Room,
			"comment":         item.Comment,
			"ground_image":    item.GroundImage,
			"created_at":      item.CreatedAt,
			"updated_at":      item.UpdatedAt,
			"DeletedAt":       item.DeletedAt,
			"ground_imageurl": presignedURL.String(),
		}

		response = map[string]interface{}{
			"ground": groundMap,
		}
	} else {
		// แปลง item ให้เป็น map[string]interface{}
		groundMap := map[string]interface{}{
			"ground_id":     item.GroundID,
			"ground_name":   item.GroundName,
			"location_code": item.Location_code,
			"building":      item.Building,
			"floor":         item.Floor,
			"room":          item.Room,
			"comment":       item.Comment,
			"ground_image":  item.GroundImage,
			"created_at":    item.CreatedAt,
			"updated_at":    item.UpdatedAt,
			"DeletedAt":     item.DeletedAt,
		}

		response = map[string]interface{}{
			"ground": groundMap,
		}
	}

	return c.JSON(response)
}

func UpdateGroundID(c *fiber.Ctx) error {
	id := c.Params("ground_id")

	updatedType := new(model.Ground)
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

	_, err = services.UpdateGround(id, *updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedBranch, err := services.GetById(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedBranch,
	})
}

func DeleteGroundIDHandler(c *fiber.Ctx) error {
	GroundID := c.Params("ground_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	asset_count, err := services.GetById(GroundID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteGroundID(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": asset_count,
	})
}
