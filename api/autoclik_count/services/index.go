package ServicesAutoclik_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_count, Error error) {
	db := database.DB
	count_autoclik := model.Autoclik_count{Autoclik_countID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func GetAllRecords() (result []*model.Autoclik_count, err error) {
	db := database.DB
	var records []*model.Autoclik_count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_Count(count_autoclik model.Autoclik_count) (value *model.Autoclik_count, Error error) {
	db := database.DB
	count_autoclik.Autoclik_countID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateAutoclik_Count(id string, updatedType model.Autoclik_count) (value *model.Autoclik_count, Error error) {
	db := database.DB
	existingType := model.Autoclik_count{Autoclik_countID: id}
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

func DeleteAutoclik_Count(typeplan *model.Autoclik_count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
