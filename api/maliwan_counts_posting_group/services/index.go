package ServicesMaliwan_Counts_Gen_Posting_Group

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Counts_Item_Category_Code, Error error) {
	db := database.DB
	count_autoclik := model.Maliwan_Counts_Item_Category_Code{Maliwan_Counts_Item_Category_CodeID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

// func GetAllRecords() (result []*model.Maliwan_Counts_Gen_Posting_Group, err error) {
// 	db := database.DB
// 	var records []*model.Maliwan_Counts_Gen_Posting_Group
// 	tx := db.Preload(clause.Associations).Find(&records)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return records, nil
// }

// func GetById(id string) (value *model.Main_Branch_Story, Error error) {
// 	db := database.DB
// 	main_branch := model.Main_Branch_Story{MainBranchStoryID: id}
// 	tx := db.Preload(clause.Associations).Find(&main_branch)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &main_branch, nil
// }

func GetAllRecords() (result []*model.Maliwan_Counts_Item_Category_Code, err error) {
	db := database.DB
	var records []*model.Maliwan_Counts_Item_Category_Code
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func GetAllRecords() (result []*model.Maliwan_Counts_Gen_Posting_Group, err error) {
// 	db := database.DB
// 	var records []*model.Maliwan_Counts_Gen_Posting_Group
// 	tx := db.Preload("Autoclik_count").Find(&records)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return records, nil
// }

func CreateNewMaliwan_Counts_Gen_Posting_Group(count_autoclik model.Maliwan_Counts_Item_Category_Code) (value *model.Maliwan_Counts_Item_Category_Code, Error error) {
	db := database.DB
	count_autoclik.Maliwan_Counts_Item_Category_CodeID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateMaliwan_Counts_Gen_Posting_Group(id string, updatedType model.Maliwan_Counts_Item_Category_Code) (value *model.Maliwan_Counts_Item_Category_Code, Error error) {
	db := database.DB
	existingType := model.Maliwan_Counts_Item_Category_Code{Maliwan_Counts_Item_Category_CodeID: id}
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

func DeleteMaliwan_Counts_Gen_Posting_Group(typeplan *model.Maliwan_Counts_Item_Category_Code) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetBymaliwan_count_IDDB(Asset_countID string) ([]*model.Maliwan_Counts_Item_Category_Code, error) {
	db := database.DB
	var asset_count []*model.Maliwan_Counts_Item_Category_Code
	tx := db.Preload(clause.Associations).Where("maliwan_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}

func GetByItem_Autoclik_BinNoDB(No string) ([]*model.Item_Autoclik_Bin_Code, error) {
	db := database.DB
	var autoclik []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).Where("item_no = ?", No).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik, nil
}
