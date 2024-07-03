package ServicesMaliwan_Counting_Trigger

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Counting_Trigger, Error error) {
	db := database.DB
	types := model.Maliwan_Counting_Trigger{Maliwan_Counting_TriggerID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetBymaliwan_count_IDDB(Maliwan_countID string) ([]string, error) {
	db := database.DB
	var itemCategoryCodes []string
	var maliwanCountsItemCategoryCodes []*model.Maliwan_Counts_Item_Category_Code
	tx := db.Where("maliwan_count_id = ?", Maliwan_countID).Find(&maliwanCountsItemCategoryCodes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// Extract only the item category codes
	for _, code := range maliwanCountsItemCategoryCodes {
		if code.Item_Category_Code != nil {
			itemCategoryCodes = append(itemCategoryCodes, *code.Item_Category_Code)
		}
	}
	return itemCategoryCodes, nil
}

func CreateNewMaliwan_Count_StoreData(maliwan_count_store model.Maliwan_Count_Store) (value *model.Maliwan_Count_Store, Error error) {
	db := database.DB
	maliwan_count_store.Maliwan_Count_StoreID = uuid.New().String()
	tx := db.Create(&maliwan_count_store)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_count_store, nil
}

func GetByMaliwan_dataNoDB(Item_Category_Code []string, branch string, gen_prod_posting_group_maliwan string) ([]*model.Maliwan_data, error) {
	db := database.DB
	var maliwan_data []*model.Maliwan_data
	tx := db.Preload(clause.Associations).
		Where("item_category_code IN ? AND branch  = ? AND gen_prod_posting_group = ?", Item_Category_Code, branch, gen_prod_posting_group_maliwan).
		Find(&maliwan_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan_data, nil
}

func GetAllRecords() (result []*model.Maliwan_Counting_Trigger, err error) {
	db := database.DB
	var records []*model.Maliwan_Counting_Trigger
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateMaliwan_Counting_Trigger(types model.Maliwan_Counting_Trigger) (value *model.Maliwan_Counting_Trigger, Error error) {
	db := database.DB
	var existingTrigger model.Maliwan_Counting_Trigger
	if err := db.Where("maliwan_count_id = ?", types.Maliwan_countID).First(&existingTrigger).Error; err == nil {
		return nil, errors.New("Autoclik_countID already exists")
	}
	types.Maliwan_Counting_TriggerID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateMaliwan_Counting_Trigger(id string, updatedType model.Maliwan_Counting_Trigger) (value *model.Maliwan_Counting_Trigger, Error error) {
	db := database.DB
	existingType := model.Maliwan_Counting_Trigger{Maliwan_Counting_TriggerID: id}
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

func DeleteCountingTriggerByID(Maliwan_Counting_TriggerID string) error {
	db := database.DB
	tx := db.Unscoped().Where("maliwan_counting_trigger_id = ?", Maliwan_Counting_TriggerID).Delete(&model.Maliwan_Counting_Trigger{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteByBranchMaliwan(Maliwan_Branch string) error {
	db := database.DB
	// tx := db.Unscoped().Where("branch_maliwan = ?", Maliwan_Branch).Delete(&model.Maliwan_data{})
	tx := db.Unscoped().Where("branch = ?", Maliwan_Branch).Delete(&model.Maliwan_data{})
	fmt.Println("สาขาที่ทำการลบ", Maliwan_Branch)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAsset_countIDDB(Asset_countID string) ([]*model.Asset_count_Category, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}

func GetByMaliwan_countID_DB(id string) (value *model.Maliwan_count, Error error) {
	db := database.DB
	count_autoclik := model.Maliwan_count{Maliwan_countID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateMaliwanCountStatusPullData(id string) (*model.Maliwan_count, error) {
	db := database.DB
	existingType := model.Maliwan_count{Maliwan_countID: id}

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
