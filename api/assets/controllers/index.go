package ControllerAssets

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	ServicesAssets "github.com/puvadon-artmit/gofiber-template/api/assets/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetById(c.Params("assets_id"))
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

func GetByIdHandler(c *fiber.Ctx) error {
	AssetsID := c.Params("assets_id")

	item, err := ServicesAssets.GetById(AssetsID)
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
	var frontpictureURL *string

	if item.Item_model.Frontpicture != nil && *item.Item_model.Frontpicture != "" {
		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *item.Item_model.Frontpicture, expiry, nil)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
		}

		presignedURLStr := presignedURL.String()
		frontpictureURL = &presignedURLStr
	}

	itemMap := map[string]interface{}{
		"model_name":      item.Model_name,
		"manufacturer":    item.Manufacturer,
		"serial_code":     item.Serial_Code,
		"type":            item.Type,
		"model":           item.Model,
		"branch":          item.Branch,
		"username":        item.Username,
		"property_code":   item.Property_code,
		"status":          item.Status,
		"group_hardware":  item.Group_hardware,
		"group":           item.Group,
		"user_hardware":   item.User_hardware,
		"phone_number":    item.Phone_number,
		"posting_group":   item.Posting_group,
		"comment1":        item.Comment1,
		"comment2":        item.Comment2,
		"comment3":        item.Comment3,
		"responsible_id":  item.ResponsibleID,
		"user_id":         item.UserID,
		"category_id":     item.CategoryID,
		"item_model_id":   item.ItemModelID,
		"ground_id":       item.GroundID,
		"frontpictureurl": frontpictureURL,
		"Item_model": map[string]interface{}{
			"item_model_id":   item.Item_model.ItemModelID,
			"item_model_name": item.Item_model.ItemModelName,
			"comment":         item.Item_model.Comment,
			"product_number":  item.Item_model.ProductNumber,
			"weight":          item.Item_model.Weight,
			"required_units":  item.Item_model.RequiredUnits,
			"frontpicture":    item.Item_model.Frontpicture,
			"created_at":      item.Item_model.CreatedAt,
			"updated_at":      item.Item_model.UpdatedAt,
			"DeletedAt":       item.Item_model.DeletedAt,
		},
		"Responsible": map[string]interface{}{
			"responsible_id":   item.Responsible.ResponsibleID,
			"responsible_name": item.Responsible.ResponsibleName,
			"employee_code":    item.Responsible.EmployeeCode,
			"comment":          item.Responsible.Comment,
			"created_at":       item.Responsible.CreatedAt,
			"updated_at":       item.Responsible.UpdatedAt,
			"DeletedAt":        item.Responsible.DeletedAt,
		},
		"User": map[string]interface{}{
			"user_id":       item.User.UserID,
			"firstname":     item.User.Firstname,
			"lastname":      item.User.Lastname,
			"username":      item.User.Username,
			"password":      item.User.Password,
			"status":        item.User.Status,
			"employee_code": item.User.EmployeeCode,
			"role_active":   item.User.RoleActive,
			"created_at":    item.User.CreatedAt,
			"updated_at":    item.User.UpdatedAt,
			"DeletedAt":     item.User.DeletedAt,
		},
		"Category": map[string]interface{}{
			"category_id":    item.Category.CategoryID,
			"category_name":  item.Category.CategoryName,
			"category_image": item.Category.CategoryImage,
			"created_at":     item.Category.CreatedAt,
			"updated_at":     item.Category.UpdatedAt,
			"DeletedAt":      item.Category.DeletedAt,
		},
		"Ground": map[string]interface{}{
			"ground_id":     item.Ground.GroundID,
			"ground_name":   item.Ground.GroundName,
			"location_code": item.Ground.Location_code,
			"building":      item.Ground.Building,
			"floor":         item.Ground.Floor,
			"room":          item.Ground.Room,
			"ground_image":  item.Ground.GroundImage,
			"comment":       item.Ground.Comment,
			"created_at":    item.Ground.CreatedAt,
			"updated_at":    item.Ground.UpdatedAt,
			"DeletedAt":     item.Ground.DeletedAt,
		},
	}

	response = map[string]interface{}{
		"result": itemMap,
	}

	return c.JSON(response)
}

func GetByCategoryHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetByCategoryIDDB(c.Params("category_id"))
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

func GetCountAssets(c *fiber.Ctx) error {
	count, err := ServicesAssets.CountLatestRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)

}

func GetCountAllAssets(c *fiber.Ctx) error {
	allcount, err := ServicesAssets.CountAssetsByCategory()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(allcount)

}

func GetByProperty_codeHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetByProperty_codeDB(c.Params("property_code"))
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

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesAssets.SyncRedisWithDatabase()
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

func GetDeletePropertyCodes(c *fiber.Ctx) error {
	err := ServicesAssets.DeleteUnusedPropertyCodes()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Unused property codes deleted successfully",
	})
}

func Create(c *fiber.Ctx) error {
	assets := new(model.Assets)
	err := c.BodyParser(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(assets)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAssets.CreateNewAssets(*assets)
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

// func Create(c *fiber.Ctx) error {
// 	assets := make([]model.Assets, 0)
// 	err := c.BodyParser(&assets)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	for _, asset := range assets {
// 		err = validator.New().Struct(asset)
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"status": "error",
// 				"error":  err.Error(),
// 			})
// 		}
// 	}

// 	createdAssets, err := ServicesAssets.CreateNewAssets(assets)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": createdAssets,
// 	})
// }

func UpdateAssetByID(c *fiber.Ctx) error {
	id := c.Params("assets_id")

	updatedLocation := new(model.Assets)
	err := c.BodyParser(updatedLocation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedLocation)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = ServicesAssets.UpdateAssetByID(id, *updatedLocation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedLocation,
	})
}

func DeleteAsset(c *fiber.Ctx) error {
	AssetsID := c.Params("assets_id")

	assets, err := ServicesAssets.GetById(AssetsID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAssets.DeleteAsset(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "DeleteAsset deleted successfully",
	})
}

func DeletePropertyCodes(c *fiber.Ctx) error {
	AssetsID := c.Params("assets_id")

	assets, err := ServicesAssets.GetById(AssetsID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAssets.DeleteAsset(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "DeleteAsset deleted successfully",
	})
}
