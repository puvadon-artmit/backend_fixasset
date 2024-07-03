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
	"github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("autoclik_signature_id"))
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
	Autoclik_countID := c.Params("autoclik_count_id")

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
		if signature.Autoclik_Signature == nil {
			continue
		}

		filename := *signature.Autoclik_Signature

		presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), filename, expiry, nil)
		if err != nil {
			log.Println(err)
			continue
		}

		response = append(response, fiber.Map{
			"autoclik_signature_id": signature.Signature_AutoclikID,
			"autoclik_signature":    signature.Autoclik_Signature,
			"user_id":               signature.UserID,
			"user":                  signature.User,
			"created_at":            signature.CreatedAt,
			"updated_at":            signature.UpdatedAt,
			"deleted_at":            signature.DeletedAt,
			"autoclik_count_id":     signature.Autoclik_countID,
			"autoclik_count":        signature.Autoclik_count,
			"presigned_url":         presignedURL.String(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": response,
	})
}

// func UpdateSignatureHandler(c *fiber.Ctx) error {
// 	Autoclik_SignatureID := c.Params("signature_id")

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	var updatedItem model.Signature
// 	if err := c.BodyParser(&updatedItem); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	existingSignature, err := services.GetById(Autoclik_SignatureID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	// existingSignature.SignatureName = updatedItem.SignatureName

// 	if len(form.File["signature"]) > 0 {
// 		file := form.File["signature"][0]

// 		randomName := uuid.New().String()
// 		filename := randomName + "_" + file.Filename
// 		err := c.SaveFile(file, "./uploads/"+filename)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 		}
// 		existingSignature.Signature = &filename
// 	}

// 	if err := services.UpdateSignatureByID(Autoclik_SignatureID, *existingSignature); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": existingSignature,
// 	})

// }

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
	SignatureID := c.Params("autoclik_signature")

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

	filename := *signature.Autoclik_Signature

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), filename, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

// func GetByIdHandlers(c *fiber.Ctx) error {
// 	Autoclik_SignatureID := c.Params("autoclik_signature_id")

// 	signature, err := services.GetById(Autoclik_SignatureID)
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

// 	filename := *signature.Autoclik_Signature

// 	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), os.Getenv("AWS_BUCKET_NAME"), filename, expiry, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
// 	}

// 	return c.SendString(presignedURL.String())
// }

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
		if records[i].Autoclik_Signature == nil || *records[i].Autoclik_Signature == "" {
			continue
		}
		imageURL, err := minioClient.PresignedGetObject(context.TODO(), viper.GetString("aws.bucket"), *records[i].Autoclik_Signature, expiry, nil)
		if err != nil {
			log.Println("Error generating pre-signed URL for record:", records[i].Signature_AutoclikID, "Error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to generate pre-signed URL",
				"error":   err,
			})
		}
		imageURLStr := imageURL.String()
		records[i].Autoclik_Signature = &imageURLStr
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

// type CreateSignatureInput struct {
// 	Autoclik_SignatureID   string `json:"signature_id"`
// 	Signature     string `json:"signature"`
// 	Round_countID string `json:"asset_count"`
// 	SignatureName string `form:"signature_name"`
// 	UserID        string `json:"user_id"`
// }

// type UploadForm struct {
// 	Round_countID string `form:"asset_count_id"`
// 	SignatureName string `form:"signature_name"`
// 	Signature     string `form:"signature"`
// 	UserID        string `form:"user_id"`
// }

// func UploadFile(c *fiber.Ctx) error {
// 	form := new(UploadForm)
// 	if err := c.BodyParser(form); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	// ตรวจสอบว่า Round_countID ไม่เป็นค่าว่างเปล่า
// 	if form.Round_countID == "" {
// 		return c.Status(fiber.StatusBadRequest).SendString("Round_countID is required")
// 	}

// 	file, err := c.FormFile("signature")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	randomName := uuid.New().String()
// 	filename := randomName + "_" + file.Filename

// 	err = c.SaveFile(file, "./uploads/"+filename)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	Autoclik_SignatureID := uuid.New().String()
// 	signatureImageName := filename

// 	input := CreateSignatureInput{
// 		SignatureName: form.SignatureName,
// 		Signature:     signatureImageName,
// 		Round_countID: form.Round_countID,
// 		UserID:        form.UserID,
// 	}

// 	item_model := model.Signature{
// 		Autoclik_SignatureID:   Autoclik_SignatureID,
// 		SignatureName: &input.SignatureName,
// 		Signature:     &signatureImageName,
// 		UserID:        input.UserID,
// 	}

// 	if err := database.DB.Create(&item_model).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	return c.SendString(Autoclik_SignatureID)
// }

func Create(c *fiber.Ctx) error {
	signature := new(model.Signature_Autoclik)
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

	createdSignature, err := services.CreateNewSignature_Autoclik(*signature)
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
