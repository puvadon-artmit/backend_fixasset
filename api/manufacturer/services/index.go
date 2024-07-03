package ServicesManufacturer

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Manufacturer, Error error) {
	db := database.DB
	manufacturer := model.Manufacturer{ManufacturerID: id}
	tx := db.Preload(clause.Associations).Find(&manufacturer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer, nil
}

func GetAllRecords() (result []*model.Manufacturer, err error) {
	db := database.DB
	var records []*model.Manufacturer
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CountAllRecords() (count int, err error) {
	latestRecords, err := GetAllRecords()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func UpdateManufacturer(id string, updatedType model.Manufacturer) (value *model.Manufacturer, Error error) {
	db := database.DB
	existingType := model.Manufacturer{ManufacturerID: id}
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

func CreateNewManufacturer(manufacturer model.Manufacturer) (value *model.Manufacturer, Error error) {
	db := database.DB
	manufacturer.ManufacturerID = uuid.New().String()
	tx := db.Create(&manufacturer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer, nil
}
