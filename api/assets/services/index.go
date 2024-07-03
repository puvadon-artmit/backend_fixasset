package ServicesAssets

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/api/redis"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Assets, Error error) {
	db := database.DB
	assets := model.Assets{AssetsID: id}
	tx := db.Preload(clause.Associations).Find(&assets)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assets, nil
}

// func GetLatestRecords() (result []*model.Assets, err error) {
// 	db := database.DB
// 	var records []*model.Assets
// 	tx := db.Preload(clause.Associations).Find(&records)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	latestTimeMap := make(map[string]time.Time)

// 	for _, record := range records {
// 		if latestTime, ok := latestTimeMap[record.Property_code]; !ok || record.Latest_time.After(latestTime) {
// 			latestTimeMap[record.Property_code] = record.Latest_time
// 		}
// 	}

// 	for _, record := range records {
// 		if record.Latest_time.Equal(latestTimeMap[record.Property_code]) {
// 			result = append(result, record)
// 		}
// 	}

// 	return result, nil
// }

func GetLatestRecords() (result []*model.Assets, err error) {
	db := database.DB
	var records []*model.Assets

	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	uniqueAssets := make(map[string]*model.Assets)
	for _, asset := range records {
		uniqueAssets[asset.Property_code] = asset
	}

	result = make([]*model.Assets, 0, len(uniqueAssets))
	for _, asset := range uniqueAssets {
		result = append(result, asset)
	}

	return result, nil
}

func CountLatestRecords() (count int, err error) {
	latestRecords, err := GetLatestRecords()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func CountLatestAllAssets() (count int, err error) {
	latestAssets, err := GetLatestRecords()
	if err != nil {
		return 0, err
	}
	return len(latestAssets), nil
}

func CountAssetsByCategory() (map[string]int, error) {
	latestAssets, err := GetLatestRecords()
	if err != nil {
		return nil, err
	}

	// สร้างแผนที่เพื่อเก็บจำนวน Assets ของแต่ละ CategoryID
	categoryCounts := make(map[string]int)

	// นับจำนวน Assets ของแต่ละ CategoryID
	for _, asset := range latestAssets {
		categoryID := asset.CategoryID
		categoryCounts[categoryID]++
	}

	return categoryCounts, nil
}

func DeleteAsset(asset *model.Assets) error {
	db := database.DB
	tx := db.Delete(asset)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteUnusedPropertyCodes() error {
	latestRecords, err := GetLatestRecords()
	if err != nil {
		return err
	}

	latestPropertyCodes := make(map[string]bool)
	for _, record := range latestRecords {
		latestPropertyCodes[record.Property_code] = true
	}

	db := database.DB
	var allRecords []*model.Assets
	tx := db.Find(&allRecords)
	if tx.Error != nil {
		return tx.Error
	}

	for _, record := range allRecords {
		if _, isLatest := latestPropertyCodes[record.Property_code]; !isLatest {
			if err := DeleteAsset(record); err != nil {
				return err
			}
		}
	}

	return nil
}

// func CreateNewAssets(assets []model.Assets) ([]*model.Assets, error) {
// 	db := database.DB
// 	createdAssets := make([]*model.Assets, 0)
// 	tx := db.Begin()
// 	for _, asset := range assets {
// 		newAsset := asset
// 		newAsset.AssetsID = uuid.New().String()
// 		if err := tx.Create(&newAsset).Error; err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}
// 		createdAssets = append(createdAssets, &newAsset)
// 	}
// 	tx.Commit()
// 	return createdAssets, nil
// }

//เพิ่มสินทรัพย์เวอร์ชั่นเก่า
// func CreateNewAssets(assets model.Assets) (value *model.Assets, Error error) {
// 	db := database.DB
// 	assets.AssetsID = uuid.New().String()

// 	tx := db.Create(&assets)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &assets, nil
// }

func CreateNewAssets(assets model.Assets) (*model.Assets, error) {
	db := database.DB
	assets.AssetsID = uuid.New().String()

	tx := db.Create(&assets)
	if tx.Error != nil {
		return nil, tx.Error
	}

	err := redis.SetcacheAssets("latest_assets", []*model.Assets{&assets})
	if err != nil {
		fmt.Println("Successfully retrieved data from cache.")
		return nil, err
	}

	return &assets, nil
}

func SyncRedisWithDatabase() (result []*model.Assets, err error) {
	// ดึงข้อมูลจากแคช (ถ้ามี)
	cachedAssets, err := redis.GetcacheAssets("latest_assets")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		result = append(result, cachedAssets...)
	}

	// ดึงข้อมูลทั้งหมดจากฐานข้อมูล
	db := database.DB
	var records []*model.Assets
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	uniqueAssets := make(map[string]*model.Assets)
	for _, asset := range records {
		uniqueAssets[asset.Property_code] = asset
	}

	for _, asset := range records {
		found := false
		for _, cachedAsset := range cachedAssets {
			if cachedAsset.AssetsID == asset.AssetsID {
				found = true
				break
			}
		}
		if !found {
			result = append(result, asset)
		}
	}

	// อัปเดตแคช "latest_assets" ด้วยข้อมูลใหม่ (ไม่รวมข้อมูลที่มีอยู่ในแคชแล้ว)
	err = redis.SetcacheAssets("latest_assets", result)
	if err != nil {
		fmt.Println("Failed to cache latest assets:", err)
	}

	return result, nil
}

// func SyncRedisWithDatabase() (result []*model.Assets, err error) {

// 	cachedAssets, err := redis.GetcacheAssets("latest_assets")
// 	if err == nil {
// 		fmt.Println("Successfully retrieved data from cache.")
// 		return cachedAssets, nil
// 	}

// 	db := database.DB
// 	var records []*model.Assets

// tx := db.Preload(clause.Associations).Find(&records)
// if tx.Error != nil {
// 	return nil, tx.Error
// }

// uniqueAssets := make(map[string]*model.Assets)
// for _, asset := range records {
// 	uniqueAssets[asset.Property_code] = asset
// }

// 	result = make([]*model.Assets, 0, len(uniqueAssets))
// 	for _, asset := range uniqueAssets {
// 		result = append(result, asset)
// 	}

// 	err = redis.SetcacheAssets("latest_assets", result)
// 	if err != nil {
// 		fmt.Println("Failed to cache latest assets:", err)
// 	}

// 	return result, nil
// }

// func SyncRedisWithDatabase() error {
// 	latestRecords, err := GetLatestRecords()
// 	if err != nil {
// 		return err
// 	}

// 	for _, record := range latestRecords {
// 		existingData, err := redis.GetcacheAssets(record.Property_code)
// 		if err != nil {
// 			return err
// 		}

// 		if len(existingData) == 0 {
// 			err := redis.SetcacheAssets(record.Property_code, []*model.Assets{record})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func CreateNewAssets(assets []model.Assets) ([]*model.Assets, error) {
// 	db := database.DB
// 	createdAssets := make([]*model.Assets, 0)

// 	for _, asset := range assets {
// 		var existingAsset model.Assets
// 		db.Where("property_code = ?", asset.Property_code).Order("latest_time desc").First(&existingAsset)

// 		if existingAsset.AssetsID != "" {
// 			asset.Latest_time = existingAsset.Latest_time
// 		} else {
// 			newAsset := asset
// 			newAsset.AssetsID = uuid.New().String()
// 			tx := db.Create(&newAsset)
// 			if tx.Error != nil {
// 				return nil, tx.Error
// 			}
// 			createdAssets = append(createdAssets, &newAsset)
// 		}
// 	}

// 	return createdAssets, nil
// }

// func CreateNewAssets(assets []model.Assets) ([]*model.Assets, error) {
// 	db := database.DB
// 	createdAssets := make([]*model.Assets, 0)

// 	for _, asset := range assets {
// 		var existingAsset model.Assets
// 		db.Where("property_code = ?", asset.Property_code).Order("latest_time desc").First(&existingAsset)

// 		if existingAsset.AssetsID != "" {
// 			asset.Latest_time = existingAsset.Latest_time
// 		} else {
// 			newAsset := asset
// 			newAsset.AssetsID = uuid.New().String()
// 			tx := db.Create(&newAsset)
// 			if tx.Error != nil {
// 				return nil, tx.Error
// 			}
// 			createdAssets = append(createdAssets, &newAsset)
// 		}
// 	}

// 	return createdAssets, nil
// }

func UpdateAssetByID(id string, updatedItem model.Assets) error {
	db := database.DB

	existingItem := model.Assets{AssetsID: id}
	tx := db.First(&existingItem)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Model(&existingItem).Updates(updatedItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetByCategoryIDDB(CategoryID string) ([]*model.Assets, error) {
	db := database.DB
	var category []*model.Assets
	tx := db.Preload(clause.Associations).Where("category_id = ?", CategoryID).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

// func GetByProperty_codeDB(Property_code string) ([]*model.Assets, error) {
// 	db := database.DB
// 	var assets []*model.Assets
// 	tx := db.Preload(clause.Associations).Where("property_code = ?", Property_code).Find(&assets)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	latestTimeMap := make(map[string]time.Time)

// 	for _, record := range assets {
// 		if latestTime, ok := latestTimeMap[record.Property_code]; !ok || record.Latest_time.After(latestTime) {
// 			latestTimeMap[record.Property_code] = record.Latest_time
// 		}
// 	}

// 	var result []*model.Assets
// 	for _, record := range assets {
// 		if record.Latest_time.Equal(latestTimeMap[record.Property_code]) {
// 			result = append(result, record)
// 		}
// 	}

// 	return assets, nil
// }

func GetByProperty_codeDB(Property_code string) ([]*model.Assets, error) {
	db := database.DB
	var assets []*model.Assets
	tx := db.Preload(clause.Associations).Where("property_code = ?", Property_code).Find(&assets)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// สร้างแผนที่เพื่อเก็บข้อมูลที่ unique โดยใช้ map[string]*model.Assets
	uniqueAssets := make(map[string]*model.Assets)
	for _, asset := range assets {
		uniqueAssets[asset.Property_code] = asset
	}

	// แปลง map กลับเป็น slice ของ assets
	uniqueAssetList := make([]*model.Assets, 0, len(uniqueAssets))
	for _, asset := range uniqueAssets {
		uniqueAssetList = append(uniqueAssetList, asset)
	}

	return uniqueAssetList, nil
}
