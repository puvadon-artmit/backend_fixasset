package item_modelController

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/item_model/services"

	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("item_model_id"))
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

func UpdateItemModelByID(c *fiber.Ctx) error {
	id := c.Params("item_model_id")

	updatedUser := new(model.Item_model)
	err := c.BodyParser(updatedUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := services.UpdatebyItemModelByID(id, *updatedUser)
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

func UpdateItemModelHandler(c *fiber.Ctx) error {
	itemModelID := c.Params("item_model_id")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var updatedItem model.Item_model
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel, err := services.GetById(itemModelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel.ItemModelName = updatedItem.ItemModelName
	existingItemModel.Comment = updatedItem.Comment
	existingItemModel.ProductNumber = updatedItem.ProductNumber
	existingItemModel.Weight = updatedItem.Weight
	existingItemModel.RequiredUnits = updatedItem.RequiredUnits

	if len(form.File["frontpicture"]) > 0 {
		file := form.File["frontpicture"][0]

		// Save the new image
		randomName := uuid.New().String()
		filename := randomName + "_" + file.Filename
		err := c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		existingItemModel.Frontpicture = &filename
	}

	if err := services.UpdateItemModelByID(itemModelID, *existingItemModel); err != nil {
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

func GetCountItemModel(c *fiber.Ctx) error {
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

type CreateItem_ModelInput struct {
	ItemModelName string `json:"item_model_name"`
	Comment       string `json:"comment"`
	ProductNumber int    `json:"product_number"`
	Weight        string `json:"weight"`
	RequiredUnits int    `json:"required_units"`
	Frontpicture  string `json:"frontpicture"`
	UserID        string `json:"user_id"`
	TypeID        string `json:"type_id"`
}

type UploadForm struct {
	UserID        string `form:"user_id"`
	TypeID        string `form:"type_id"`
	ItemModelName string `form:"item_model_name"`
	Frontpicture  string `form:"frontpicture"`
	Comment       string `form:"comment"`
	ProductNumber int    `form:"product_number"`
	Weight        string `form:"weight"`
	RequiredUnits int    `form:"required_units"`
}

// ---------------------------------------------------------------------------------------------

func UploadFile(c *fiber.Ctx) error {
	form := new(UploadForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("frontpicture")
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

	ItemModelID := uuid.New().String()
	frontpictureImageName := filename

	input := CreateItem_ModelInput{
		ItemModelName: form.ItemModelName,
		Frontpicture:  frontpictureImageName,
		UserID:        form.UserID,
		TypeID:        form.TypeID,
		Comment:       form.Comment,
		ProductNumber: form.ProductNumber,
		Weight:        form.Weight,
		RequiredUnits: form.RequiredUnits,
	}

	item_model := model.Item_model{
		ItemModelID:   ItemModelID,
		ItemModelName: &input.ItemModelName,
		Frontpicture:  &frontpictureImageName,
		UserID:        input.UserID,
		TypeID:        input.TypeID,
		Comment:       &input.Comment,

		Weight:        &input.Weight,
		RequiredUnits: &input.RequiredUnits,
	}

	if err := database.DB.Create(&item_model).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString(ItemModelID)
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func DeleteUnusedImagesHandler(c *fiber.Ctx) error {
	err := services.DeleteUnusedImages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Unused images deleted successfully")
}

func Create(c *fiber.Ctx) error {
	item_model := new(model.Item_model)
	err := c.BodyParser(item_model)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(item_model)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := services.CreateNewItem_model(*item_model)
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

func DeleteitemModel(c *fiber.Ctx) error {
	itemModelID := c.Params("item_model_id")

	itemModel, err := services.GetById(itemModelID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteItem_model(itemModel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": itemModel,
	})
}

func DeletePhotos(c *fiber.Ctx) error {
	itemModelID := c.Params("item_model_id")

	itemModel, err := services.GetById(itemModelID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteItem_model(itemModel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "ItemModel record deleted successfully",
	})
}

// func GetByIdHandler(c *fiber.Ctx) error {
// 	FrontpictureID := c.Params("item_model_id")

// 	FrontpictureURL, err := services.GetByItem_Photo(FrontpictureID)
// 	if err != nil {
// 		return c.Status(404).SendString("")
// 	}

// 	endpoint := os.Getenv("AWS_ENDPOINT")
// 	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
// 	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
// 	useSSL := true

// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	expiry := 1 * time.Minute

// 	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), os.Getenv("AWS_BUCKET_NAME"), FrontpictureURL, expiry, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
// 	}

// 	return c.SendString(presignedURL.String())
// }

func GetByIdHandler(c *fiber.Ctx) error {
	ItemModelID := c.Params("item_model_id")

	item, err := services.GetById(ItemModelID)
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

	if item.Frontpicture != nil && *item.Frontpicture != "" {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *item.Frontpicture, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		// แปลง item ให้เป็น map[string]interface{}
		itemMap := map[string]interface{}{
			"item_model_id":   item.ItemModelID,
			"item_model_name": item.ItemModelName,
			"comment":         item.Comment,
			"product_number":  item.ProductNumber,
			"weight":          item.Weight,
			"required_units":  item.RequiredUnits,
			"frontpicture":    item.Frontpicture,
			"created_at":      item.CreatedAt,
			"updated_at":      item.UpdatedAt,
			"DeletedAt":       item.DeletedAt,
			"user_id":         item.UserID,
			"type_id":         item.TypeID,
			"frontpictureurl": presignedURL.String(),
		}

		response = map[string]interface{}{
			"item": itemMap,
		}
	} else {
		// แปลง item ให้เป็น map[string]interface{}
		itemMap := map[string]interface{}{
			"item_model_id":   item.ItemModelID,
			"item_model_name": item.ItemModelName,
			"comment":         item.Comment,
			"product_number":  item.ProductNumber,
			"weight":          item.Weight,
			"required_units":  item.RequiredUnits,
			"frontpicture":    item.Frontpicture,
			"created_at":      item.CreatedAt,
			"updated_at":      item.UpdatedAt,
			"DeletedAt":       item.DeletedAt,
			"user_id":         item.UserID,
			"type_id":         item.TypeID,
		}

		response = map[string]interface{}{
			"item": itemMap,
		}
	}

	return c.JSON(response)
}
