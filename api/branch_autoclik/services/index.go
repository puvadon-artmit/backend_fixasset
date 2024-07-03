package ServicesBranch_Autoclik

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Branch_Autoclik, Error error) {
	db := database.DB
	branch_autoclik := model.Branch_Autoclik{BranchAutoclik_ID: id}
	tx := db.Preload(clause.Associations).Find(&branch_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &branch_autoclik, nil
}

func GetAllRecords() (result []*model.Branch_Autoclik, err error) {
	db := database.DB
	var records []*model.Branch_Autoclik
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewBranch_Autoclik(inspection model.Branch_Autoclik) (value *model.Branch_Autoclik, Error error) {
	db := database.DB
	inspection.BranchAutoclik_ID = uuid.New().String()

	fields := []string{"Branch_Name", "Branch_Code"}
	for _, field := range fields {
		if fieldValue := reflect.ValueOf(&inspection).Elem().FieldByName(field); fieldValue.IsValid() && fieldValue.String() == "" {
			fieldValue.SetString("")
		}
	}

	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateBranch_Autoclik(id string, updatedType model.Branch_Autoclik) (value *model.Branch_Autoclik, Error error) {
	db := database.DB
	existingType := model.Branch_Autoclik{BranchAutoclik_ID: id}
	tx := db.Model(&existingType).Where("branchautoclik_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteBranch_Autoclik(branch_autoclik *model.Branch_Autoclik) error {
	db := database.DB
	tx := db.Delete(branch_autoclik)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
