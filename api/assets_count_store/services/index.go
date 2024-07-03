package ServicesAssets_count_Store

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Assets_Count_Store, Error error) {
	db := database.DB
	count_assets_story := model.Assets_Count_Store{Assets_Count_StoreID: id}
	tx := db.Preload(clause.Associations).Find(&count_assets_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_assets_story, nil
}

func GetByAssets_Count_StoreIDDB(Assets_Count_StoreID string) ([]model.Assets_Count_Store, error) {
	db := database.DB
	var assetsStores []model.Assets_Count_Store
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Assets_Count_StoreID).Find(&assetsStores)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return assetsStores, nil
}

func GetAllRecords() (result []*model.Assets_Count_Store, err error) {
	db := database.DB
	var records []*model.Assets_Count_Store
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAssets_count_Store(count_assets_story model.Assets_Count_Store) (value *model.Assets_Count_Store, Error error) {
	db := database.DB
	count_assets_story.Assets_Count_StoreID = uuid.New().String()
	tx := db.Create(&count_assets_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_assets_story, nil
}

func UpdateAssets_count_Store(id string, updatedType model.Assets_Count_Store) (value *model.Assets_Count_Store, Error error) {
	db := database.DB
	existingType := model.Assets_Count_Store{Assets_Count_StoreID: id}
	tx := db.Preload(clause.Associations).Find(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Model(&existingType).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteAssets_count_Store(autoclik_count_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("asset_count_id = ?", autoclik_count_id).Delete(&model.Autoclik_Count_Store{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteAssets_count_ID(assets_count_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("asset_count_id = ?", assets_count_id).Delete(&model.Assets_Count_Store{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// นำหมวดมหู่หลักไปหาตารางหมวดหมู่เล็ก
func GetCategoryIDsByMainCategoryID(MainCategoryID string) ([]string, error) {
	db := database.DB
	var categories []*model.Category
	tx := db.Where("main_category_id = ?", MainCategoryID).Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// ดึงเฉพาะ category_id
	var categoryIDs []string
	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.CategoryID)
	}
	return categoryIDs, nil
}

// นำใบตรวจนับไปหาหมวดหมู่เล็ก
func GetByAsset_countID_Category(Asset_countID string) ([]map[string]string, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []map[string]string
	for _, category := range asset_count {
		result = append(result, map[string]string{
			"category_id":    *category.Category_ID,
			"asset_count_id": category.Asset_countID,
		})
	}
	return result, nil
}

func GetByAsset_countID_Main_Category(Asset_countID string) ([]map[string]interface{}, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Main_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create a slice to store category_ids
	categoryIDs := make([]string, 0)

	assetCountCategories, err := GetByAsset_countID_Category(Asset_countID)
	if err != nil {
		return nil, err
	}

	for _, mainCategory := range asset_count {
		categoryIDsMain, err := GetCategoryIDsByMainCategoryID(*mainCategory.Main_Category_ID)
		if err != nil {
			return nil, err
		}

		if len(assetCountCategories) == 0 {
			// ถ้า assetCountCategories ไม่มีค่า ให้เพิ่ม categoryIDsMain เข้าไปใน array
			for _, categoryIDMain := range categoryIDsMain {
				if categoryIDMain != "" {
					categoryIDs = append(categoryIDs, categoryIDMain)
				}
			}
		}

		// เพิ่มค่า assetCountCategories เข้าไปใน array
		for _, assetCountCategory := range assetCountCategories {
			if categoryID := assetCountCategory["category_id"]; categoryID != "" {
				categoryIDs = append(categoryIDs, categoryID)
			}
		}
	}
	// Prepare the final result
	result := make([]map[string]interface{}, 0)

	// Iterate over categoryIDs and call GetAssetbyCategoryID

	assets, err := GetAssetbyCategoryID(categoryIDs)
	if err != nil {
		return nil, err
	}

	result = append(result, map[string]interface{}{
		"category_id": categoryIDs,
		"assets":      assets,
	})

	for _, asset := range assets {
		assetCountStore := model.Assets_Count_Store{
			Model_name:     &asset.Model_name,
			Manufacturer:   &asset.Manufacturer,
			Serial_Code:    &asset.Serial_Code,
			Type:           &asset.Type,
			Model:          &asset.Model,
			Branch:         &asset.Branch,
			Username:       &asset.Username,
			Property_code:  &asset.Property_code,
			Status:         &asset.Status,
			Group_hardware: &asset.Group_hardware,
			Group:          &asset.Group,
			User_hardware:  &asset.User_hardware,
			Phone_number:   &asset.Phone_number,
			Posting_group:  &asset.Posting_group,
			Latest_time:    &asset.Latest_time,
			ResponsibleID:  &asset.ResponsibleID,
			Comment1:       &asset.Comment1,
			Comment2:       &asset.Comment2,
			Comment3:       &asset.Comment3,
			// CreatedAt:      asset.CreatedAt,
			// UpdatedAt:      asset.UpdatedAt,
			// DeletedAt:      asset.DeletedAt,
			UserID:        &asset.UserID,
			CategoryID:    &asset.CategoryID,
			ItemModelID:   asset.ItemModelID,
			GroundID:      &asset.GroundID,
			Asset_countID: &Asset_countID,
		}
		_, err := CreateNewAssets_count_Store(assetCountStore)
		if err != nil {
			return nil, err
		}

	}

	return result, nil
}

func GetAssetbyCategoryID(CategoryIDs []string) ([]*model.Assets, error) {
	db := database.DB
	var categories []*model.Assets
	tx := db.Preload(clause.Associations).Where("category_id IN ?", CategoryIDs).Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}

func UpdateAssetCountStatusPullData(id string) (*model.Asset_count, error) {
	db := database.DB
	existingType := model.Asset_count{Asset_countID: id}

	// Find the existing record
	tx := db.Preload(clause.Associations).First(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update the status_pull_data field to "succeed"
	succeed := "succeed"
	tx = db.Model(&existingType).Update("StatusPulldata", succeed)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

// func GetAssetbyCategoryID(CategoryID string) ([]*model.Assets, error) {
// 	db := database.DB
// 	var categories []*model.Assets
// 	tx := db.Preload(clause.Associations).Where("category_id = ?", CategoryID).Find(&categories)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	var result []map[string]string
// 	for _, category := range categories {
// 		result = append(result, map[string]string{
// 			"assets_id":      category.AssetsID,
// 			"model_name":     category.Model_name,
// 			"manufacturer":   category.Manufacturer,
// 			"serial_code":    category.Serial_Code,
// 			"type":           category.Type,
// 			"model":          category.Model,
// 			"branch":         category.Branch,
// 			"username":       category.Username,
// 			"property_code":  category.Property_code,
// 			"status":         category.Status,
// 			"group_hardware": category.Group_hardware,
// 			"group":          category.Group,
// 			"user_hardware":  category.User_hardware,
// 			"phone_number":   category.Phone_number,
// 			"posting_group":  category.Posting_group,
// 			"latest_time":    category.Latest_time,
// 			"responsible_id": category.ResponsibleID,
// 			"comment1":       category.Comment1,
// 			"comment2":       category.Comment2,
// 			"comment3":       category.Comment3,
// 			"user_id":        category.UserID,
// 			"category_id":    category.CategoryID,
// 			"item_model_id":  category.ItemModelID,
// 			"ground_id":      category.GroundID,
// 		})
// 	}
// 	return categories, nil
// }

func GetBytest(Asset_countID string) ([]map[string]interface{}, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Main_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create a slice to store category_ids
	categoryIDs := make([]string, 0)

	for _, mainCategory := range asset_count {
		categoryIDsMain, err := GetCategoryIDsByMainCategoryID(*mainCategory.Main_Category_ID)
		if err != nil {
			return nil, err
		}

		for _, categoryIDMain := range categoryIDsMain {
			assetCountCategories, err := GetByAsset_countID_Category(Asset_countID)
			if err != nil {
				return nil, err
			}

			for _, assetCountCategory := range assetCountCategories {
				categoryIDs = append(categoryIDs, assetCountCategory["category_id"])
				categoryIDs = append(categoryIDs, categoryIDMain)
			}
		}
	}

	// Prepare the final result
	result := []map[string]interface{}{
		{
			"category_ids": categoryIDs,
		},
	}

	return result, nil
}
