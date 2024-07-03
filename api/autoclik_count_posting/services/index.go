package ServicesAutoclik_Count_Product_Group

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Count_Product_Group, Error error) {
	db := database.DB
	count_autoclik := model.Autoclik_Count_Product_Group{Autoclik_Count_Product_GroupID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

// func GetAllRecords() (result []*model.Autoclik_Count_Product_Group, err error) {
// 	db := database.DB
// 	var records []*model.Autoclik_Count_Product_Group
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

func GetAllRecords() (result []*model.Autoclik_Count_Product_Group, err error) {
	db := database.DB
	var records []*model.Autoclik_Count_Product_Group
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func GetAllRecords() (result []*model.Autoclik_Count_Product_Group, err error) {
// 	db := database.DB
// 	var records []*model.Autoclik_Count_Product_Group
// 	tx := db.Preload("Autoclik_count").Find(&records)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return records, nil
// }

func CreateNewAutoclik_Count_Product_Group(count_autoclik model.Autoclik_Count_Product_Group) (value *model.Autoclik_Count_Product_Group, Error error) {
	db := database.DB
	count_autoclik.Autoclik_Count_Product_GroupID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateAutoclik_Count_Product_Group(id string, updatedType model.Autoclik_Count_Product_Group) (value *model.Autoclik_Count_Product_Group, Error error) {
	db := database.DB
	existingType := model.Autoclik_Count_Product_Group{Autoclik_Count_Product_GroupID: id}
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

func DeleteAutoclik_Count_Product_Group(typeplan *model.Autoclik_Count_Product_Group) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAutoclik_countIDDB(Asset_countID string) ([]*model.Autoclik_Count_Product_Group, error) {
	db := database.DB
	var asset_count []*model.Autoclik_Count_Product_Group
	tx := db.Preload(clause.Associations).Where("autoclik_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}
