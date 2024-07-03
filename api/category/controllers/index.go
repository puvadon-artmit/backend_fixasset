package categoryController

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/category/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("category_id"))
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

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("category_id")

	updatedcategory := new(model.Category)
	err := c.BodyParser(updatedcategory)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedcategory)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := services.UpdateCategory(id, *updatedcategory)
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

func GetCounCategory(c *fiber.Ctx) error {
	count, err := services.CountCategory()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

func GetByMain_CategoryHandler(c *fiber.Ctx) error {
	value, err := services.GetByMain_Category_IDDB(c.Params("main_category_id"))
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

func GetCountCategory(c *fiber.Ctx) error {
	count, err := services.CountLatestCategory()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)

}

type CreateCategoryInput struct {
	CategoryName  string `json:"category_name" validate:"required"`
	CategoryImage string `json:"category_image" validate:"required"`
	UserID        string `json:"user_id"`
}

type UploadForm struct {
	Name         string `form:"name"`
	UserID       string `form:"user_id"`
	CategoryName string `form:"category_name"`
}

func UploadFile(c *fiber.Ctx) error {
	form := new(UploadForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("category_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// สร้างชื่อสุ่มสำหรับไฟล์
	randomName := uuid.New().String()
	filename := randomName + "_" + file.Filename

	// บันทึกไฟล์ลงใน "./uploads/"
	err = c.SaveFile(file, "./uploads/"+filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	categoryID := uuid.New().String()
	categoryImageName := filename

	input := CreateCategoryInput{
		CategoryName:  form.CategoryName,
		CategoryImage: categoryImageName,
		UserID:        form.UserID,
	}

	category := model.Category{
		CategoryID:    categoryID,
		CategoryName:  input.CategoryName,
		CategoryImage: categoryImageName,
		UserID:        input.UserID,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("file upload and database record complete")
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	err := c.BodyParser(category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := services.CreateNewCategory(*category)
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

func DeleteCategory(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")

	err := services.DeleteCategory(categoryID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Category deleted successfully",
	})
}

func GetByIdHandler(c *fiber.Ctx) error {
	CategoryID := c.Params("category_id")

	CategoryImageURL, err := services.GetByCategory_Photo(CategoryID)
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

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), CategoryImageURL, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

// -------------------------------  เวอชั่นใหม่  -------------------------------------------------------------------------------------------------

func GetByIdCategoryHandler(c *fiber.Ctx) error {
	CategoryID := c.Params("category_id")

	item, err := services.GetById(CategoryID)
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

	if item.CategoryImage != "" {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), item.CategoryImage, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		// แปลง item ให้เป็น map[string]interface{}
		categoryMap := map[string]interface{}{
			"category_id":       item.CategoryID,
			"category_name":     item.CategoryName,
			"category_image":    item.CategoryImage,
			"created_at":        item.CreatedAt,
			"updated_at":        item.UpdatedAt,
			"DeletedAt":         item.DeletedAt,
			"main_category_id":  item.Main_Category_ID,
			"category_imageurl": presignedURL.String(),
		}

		response = map[string]interface{}{
			"category": categoryMap,
		}
	} else {
		// แปลง item ให้เป็น map[string]interface{}
		categoryMap := map[string]interface{}{
			"category_id":      item.CategoryID,
			"category_name":    item.CategoryName,
			"category_image":   item.CategoryImage,
			"created_at":       item.CreatedAt,
			"updated_at":       item.UpdatedAt,
			"DeletedAt":        item.DeletedAt,
			"main_category_id": item.Main_Category_ID,
		}

		response = map[string]interface{}{
			"category": categoryMap,
		}
	}

	return c.JSON(response)
}
