package ServicesMaliwan_Fixed_Asset_Store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_Store, Error error) {
	db := database.DB
	count_maliwan := model.Maliwan_Fixed_Asset_Store{Maliwan_Fixed_Asset_StoreId: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_Store, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_Store
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Fixed_Asset_Store(count_maliwan model.Maliwan_Fixed_Asset_Store) (value *model.Maliwan_Fixed_Asset_Store, Error error) {
	db := database.DB
	count_maliwan.Maliwan_Fixed_Asset_StoreId = uuid.New().String()
	tx := db.Create(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func UpdateMaliwan_Fixed_Asset_Store(id string, updatedType model.Maliwan_Fixed_Asset_Store) (value *model.Maliwan_Fixed_Asset_Store, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Store{Maliwan_Fixed_Asset_StoreId: id}
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

func DeleteMaliwan_Fixed_Asset_Store(maliwan_fixed_asset_store_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("maliwan_fixed_asset_count_id = ?", maliwan_fixed_asset_store_id).Delete(&model.Maliwan_Fixed_Asset_Store{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByMaliwan_Fixed_AssetC_CountID(Maliwan_Count_StoreID string) ([]model.Maliwan_Fixed_Asset_Store, error) {
	db := database.DB
	var maliwanStores []model.Maliwan_Fixed_Asset_Store
	tx := db.Preload(clause.Associations).Where("maliwan_fixed_asset_count_id = ?", Maliwan_Count_StoreID).Find(&maliwanStores)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwanStores, nil
}

func UpdateMaliwan_Fixed_AssetCountStatusPullData(id string) (*model.Maliwan_Fixed_Asset_Count, error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Count{Maliwan_Fixed_Asset_CountID: id}

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

func CreateMaliwan_Fixed_Asset(count_maliwan model.Maliwan_Fixed_Asset) (value *model.Maliwan_Fixed_Asset, Error error) {
	db := database.DB
	count_maliwan.Maliwan_Fixed_AssetId = uuid.New().String()
	tx := db.Create(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func GetAllMaliwan_Fixed_Asset() (result []*model.Maliwan_Fixed_Asset, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func GetByMaliwan_Fixed_Asset_Count(id string) (value *model.Maliwan_Fixed_Asset_Count, Error error) {
	db := database.DB
	count_maliwan := model.Maliwan_Fixed_Asset_Count{Maliwan_Fixed_Asset_CountID: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func UpdateMaliwanCountStatusPullData(id string) (*model.Maliwan_Fixed_Asset_Count, error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Count{Maliwan_Fixed_Asset_CountID: id}

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

func GetByNoDB(No string) ([]model.Maliwan_Fixed_Asset, error) {
	db := database.DB
	var autocliks []model.Maliwan_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("no = ?", No).Find(&autocliks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autocliks, nil
}

func GetByBranchDB(Branch string) ([]model.Maliwan_Fixed_Asset, error) {
	db := database.DB
	var autocliks []model.Maliwan_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("branch = ?", Branch).Find(&autocliks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autocliks, nil
}

func CreateMaliwan_Update_Fixed_Asset_Story(count_maliwan model.Maliwan_Update_Fixed_Asset_Story) (value *model.Maliwan_Update_Fixed_Asset_Story, Error error) {
	db := database.DB

	count_maliwan.Maliwan_Update_Fixed_Asset_StoryID = uuid.New().String()
	tx := db.Create(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &count_maliwan, nil
}

func GetAllMaliwan_Update_Fixed_Asset_Story() (result []*model.Maliwan_Update_Fixed_Asset_Story, err error) {
	db := database.DB
	var records []*model.Maliwan_Update_Fixed_Asset_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func DeleteAllDataMaliwan() error {
	db := database.DB
	result := db.Exec("DELETE FROM maliwan_fixed_assets")
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("ลบข้อมูลทั้งหมดเสร็จสิ้น")
	return nil
}

type MaliwanFixedAssetResult struct {
	No          *string `json:"no"`
	Description *string `json:"description"`
	Branch      string  `json:"branch"`
}

func GetMaliwan_Fixed_AssetDB(limit int, offset int) ([]*MaliwanFixedAssetResult, error) {
	db := database.DB

	var records []*model.Maliwan_Fixed_Asset
	tx := db.Select("no, description, branch").Limit(limit).Offset(offset).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var jsonData []*MaliwanFixedAssetResult
	for _, record := range records {
		jsonData = append(jsonData, &MaliwanFixedAssetResult{
			No:          record.No,
			Description: record.Description,
			Branch:      record.Branch,
		})
	}
	return jsonData, nil
}

func GetTotalMaliwan_Fixed_Assets() (int64, error) {
	var count int64
	db := database.DB
	tx := db.Model(&model.Maliwan_Fixed_Asset{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}

func UpdateRequest_Update_Maliwan_Fixed_Asset() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "c6fc7294-79de-4294-8137-c7c8e73550d7"
	existingType := model.Request_Update_Data{Request_Update_DataID: id}

	tx := db.Preload(clause.Associations).First(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	succeed := "succeed"
	tx = db.Model(&existingType).Update("Status", succeed)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}
