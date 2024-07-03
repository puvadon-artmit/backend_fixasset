package ServicesAutoclik_Fixed_Asset_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Fixed_Asset_Count, Error error) {
	db := database.DB
	count_autoclik := model.Autoclik_Fixed_Asset_Count{Autoclik_Fixed_Asset_CountID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func GetAllRecords() (result []*model.Autoclik_Fixed_Asset_Count, err error) {
	db := database.DB
	var records []*model.Autoclik_Fixed_Asset_Count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_Fixed_Asset_Count(count_autoclik model.Autoclik_Fixed_Asset_Count) (value *model.Autoclik_Fixed_Asset_Count, Error error) {
	db := database.DB
	count_autoclik.Autoclik_Fixed_Asset_CountID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateAutoclik_Fixed_Asset_Count(id string, updatedType model.Autoclik_Fixed_Asset_Count) (value *model.Autoclik_Fixed_Asset_Count, Error error) {
	db := database.DB
	existingType := model.Autoclik_Fixed_Asset_Count{Autoclik_Fixed_Asset_CountID: id}
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

func DeleteAutoclik_Fixed_Asset_Count(typeplan *model.Autoclik_Fixed_Asset_Count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func CreateNewAutoclik_Fixed_Asset_Count_Store(count_autoclik model.Autoclik_Fixed_Asset_Store) (value *model.Autoclik_Fixed_Asset_Store, Error error) {
	db := database.DB
	count_autoclik.Autoclik_Fixed_Asset_StoreId = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateAutoclik_Fixed_AssetCountStatusPullData(id string) (*model.Autoclik_Fixed_Asset_Count, error) {
	db := database.DB
	existingType := model.Autoclik_Fixed_Asset_Count{Autoclik_Fixed_Asset_CountID: id}

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
