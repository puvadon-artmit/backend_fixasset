package ControllerAssets_count_Store

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAssets_count_Store "github.com/puvadon-artmit/gofiber-template/api/assets_count_store/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAssets_count_Store.GetById(c.Params("assets_count_store_id"))
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

func GetByAssets_Count_StoreIDHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets_count_Store.GetByAssets_Count_StoreIDDB(c.Params("asset_count_id"))
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
	records, err := ServicesAssets_count_Store.GetAllRecords()
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

func Create(c *fiber.Ctx) error {
	assets_count_store := new(model.Assets_Count_Store)
	err := c.BodyParser(assets_count_store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(assets_count_store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAssets_count_Store.CreateNewAssets_count_Store(*assets_count_store)
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

func UpdateAssets_count_Store(c *fiber.Ctx) error {
	id := c.Params("assets_count_store_id")

	updatedAssets_count_Store := new(model.Assets_Count_Store)
	err := c.BodyParser(updatedAssets_count_Store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedAssets_count_Store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesAssets_count_Store.UpdateAssets_count_Store(id, *updatedAssets_count_Store)
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

func DeleteAssets_count_Store(c *fiber.Ctx) error {
	BranchStoryID := c.Params("asset_count_id")

	autoclik_count_store, err := ServicesAssets_count_Store.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAssets_count_Store.DeleteAssets_count_Store(autoclik_count_store.Assets_Count_StoreID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "asset_count_Store deleted successfully",
	})
}

func DeleteAssets_Count_ID(c *fiber.Ctx) error {
	countingTriggerID := c.Params("asset_count_id")

	err := ServicesAssets_count_Store.DeleteAssets_count_ID(countingTriggerID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted successfully",
	})
}

func GetCategoryIDsByMainCategoryIDHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets_count_Store.GetCategoryIDsByMainCategoryID(c.Params("main_category_id"))
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

func GetAsset_Count_By_CategoryHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets_count_Store.GetByAsset_countID_Category(c.Params("asset_count_id"))
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

func GetAsset_Count_By_Main_CategoryHandler(c *fiber.Ctx) error {
	// Get the asset count ID from the parameters
	assetCountID := c.Params("asset_count_id")

	// Retrieve the asset count data by main category
	value, err := ServicesAssets_count_Store.GetByAsset_countID_Main_Category(assetCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Update the asset count status
	_, err = ServicesAssets_count_Store.UpdateAssetCountStatusPullData(assetCountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Return the successful response
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetGetAssetbyCategoryIDHandler(c *fiber.Ctx) error {
	categoryID := c.Params("category_id")
	value, err := ServicesAssets_count_Store.GetAssetbyCategoryID([]string{categoryID})
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
