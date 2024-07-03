package ServicesMaliwan_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_count, Error error) {
	db := database.DB
	count_autoclik := model.Maliwan_count{Maliwan_countID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func GetAllRecords() (result []*model.Maliwan_count, err error) {
	db := database.DB
	var records []*model.Maliwan_count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Count(count_autoclik model.Maliwan_count) (value *model.Maliwan_count, Error error) {
	db := database.DB
	count_autoclik.Maliwan_countID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateMaliwan_Count(id string, updatedType model.Maliwan_count) (value *model.Maliwan_count, Error error) {
	db := database.DB
	existingType := model.Maliwan_count{Maliwan_countID: id}
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

func DeleteMaliwan_Count(typeplan *model.Maliwan_count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
