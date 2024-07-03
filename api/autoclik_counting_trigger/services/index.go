package ServicesAutoclik_Counting_Trigger

import (
	"errors"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Counting_Trigger, Error error) {
	db := database.DB
	types := model.Autoclik_Counting_Trigger{Autoclik_Counting_TriggerID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetByIdAutoclik_CountingArray(id string) (value *model.Autoclik_Counting_Trigger, Error error) {
	db := database.DB
	types := model.Autoclik_Counting_Trigger{Autoclik_Counting_TriggerID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetByIdBranch(id string) (value *model.Branch_Autoclik, Error error) {
	db := database.DB
	branch_autoclik := model.Branch_Autoclik{BranchAutoclik_ID: id}
	tx := db.Preload(clause.Associations).Find(&branch_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &branch_autoclik, nil
}

func GetByBranchIDDB(BranchAutoclik_ID string) ([]*model.Autoclik_count, error) {
	db := database.DB
	var autoclik_count_product_group []*model.Autoclik_count
	tx := db.Preload(clause.Associations).Where("branch_autoclik_id = ?", BranchAutoclik_ID).Find(&autoclik_count_product_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik_count_product_group, nil
}

func GetByAutoclik_countIDDB(Autoclik_countID string) ([]*model.Autoclik_Count_Product_Group, error) {
	db := database.DB
	var autoclik_count_product_group []*model.Autoclik_Count_Product_Group
	tx := db.Preload(clause.Associations).Where("autoclik_count_id = ?", Autoclik_countID).Find(&autoclik_count_product_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik_count_product_group, nil
}

func GetAllRecords() (result []*model.Autoclik_Counting_Trigger, err error) {
	db := database.DB
	var records []*model.Autoclik_Counting_Trigger
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateAutoclik_Counting_Trigger(types model.Autoclik_Counting_Trigger) (value *model.Autoclik_Counting_Trigger, Error error) {
	db := database.DB

	var existingTrigger model.Autoclik_Counting_Trigger
	if err := db.Where("autoclik_count_id = ?", types.Autoclik_countID).First(&existingTrigger).Error; err == nil {
		return nil, errors.New("Autoclik_countID already exists")
	}
	types.Autoclik_Counting_TriggerID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateAutoclik_Counting_Trigger(id string, updatedType model.Autoclik_Counting_Trigger) (value *model.Autoclik_Counting_Trigger, Error error) {
	db := database.DB
	existingType := model.Autoclik_Counting_Trigger{Autoclik_Counting_TriggerID: id}
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

func GetByAutoclik_countID(id string) (value *model.Autoclik_count, Error error) {
	db := database.DB
	count_autoclik := model.Autoclik_count{Autoclik_countID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func DeleteCountingTriggerByID(autoclik_counting_trigger_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("autoclik_counting_trigger_id = ?", autoclik_counting_trigger_id).Delete(&model.Autoclik_Counting_Trigger{})
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

func GetByGen_Prod_Posting_GroupDB(No string) ([]*model.Item_Autoclik, error) {
	db := database.DB
	var gen_prod_posting_group []*model.Item_Autoclik
	tx := db.Preload(clause.Associations).Where("gen_prod_posting_group = ?", No).Find(&gen_prod_posting_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return gen_prod_posting_group, nil
}

func GetBytriggerIDDB(No string) ([]*model.Autoclik_Counting_Trigger, error) {
	db := database.DB
	var trigger []*model.Autoclik_Counting_Trigger
	tx := db.Preload(clause.Associations).Where("autoclik_counting_trigger_id = ?", No).Find(&trigger)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return trigger, nil
}

// func GetByProductGroupNoDB(productGroups []string) ([]*model.Item_Autoclik_Bin_Code, error) {
// 	db := database.DB
// 	var product_group []*model.Item_Autoclik_Bin_Code
// 	tx := db.Preload(clause.Associations).Where("product_group IN ?", productGroups).Find(&product_group)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return product_group, nil
// }

func GetByProductGroupNoDB(productGroups []string, locationCode string) ([]*model.Item_Autoclik_Bin_Code, error) {
	db := database.DB
	var product_group []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).
		Where("product_group IN ? AND location_code = ?", productGroups, locationCode).
		Find(&product_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product_group, nil
}

func CreateNewAutoclik_Storedata(autoclik_data model.Autoclik_Count_Store) (value *model.Autoclik_Count_Store, Error error) {
	db := database.DB
	autoclik_data.Autoclik_Count_StoreID = uuid.New().String()
	tx := db.Create(&autoclik_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_data, nil
}

func ReaddataAndCreate(productGroups []string, locationCode string) error {
	values, err := GetByProductGroupNoDB(productGroups, locationCode)
	if err != nil {
		return err
	}

	for _, value := range values {
		autoclik_data := model.Autoclik_Count_Store{
			Odata_etag:           value.Odata_etag,
			Location_Code:        value.Location_Code,
			Bin_Code:             value.Bin_Code,
			Item_No:              value.Item_No,
			Variant_Code:         value.Variant_Code,
			Unit_of_Measure_Code: value.Unit_of_Measure_Code,
			ItemDesc:             value.ItemDesc,
			ItemDesc2:            value.ItemDesc2,
			ItemCategory:         value.ItemCategory,
			ProductGroup:         value.ProductGroup,
			Fixed:                value.Fixed,
			Default:              value.Default,
			Dedicated:            value.Dedicated,
			CalcQtyUOM:           value.CalcQtyUOM,
			Quantity_Base:        value.Quantity_Base,
			Bin_Type_Code:        value.Bin_Type_Code,
			Zone_Code:            value.Zone_Code,
			Lot_No_Filter:        value.Lot_No_Filter,
			Serial_No_Filter:     value.Serial_No_Filter,
		}

		_, err := CreateNewAutoclik_Storedata(autoclik_data)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateAutoclikCountStatusPullData(id string) (*model.Autoclik_count, error) {
	db := database.DB
	existingType := model.Autoclik_count{Autoclik_countID: id}

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
