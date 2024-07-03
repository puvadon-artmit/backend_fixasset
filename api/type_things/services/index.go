package ServicesType

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Type_things, Error error) {
	db := database.DB
	types := model.Type_things{TypeID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Type_things, err error) {
	db := database.DB
	var records []*model.Type_things
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CountLatestRecords() (count int, err error) {
	latestRecords, err := GetAllRecords()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func CreateNewType_things(types model.Type_things) (value *model.Type_things, Error error) {
	db := database.DB
	types.TypeID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateType_things(id string, updatedType model.Type_things) (value *model.Type_things, Error error) {
	db := database.DB
	existingType := model.Type_things{TypeID: id}
	tx := db.Model(&existingType).Where("type_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteType_thingsID(type_things *model.Type_things) error {
	db := database.DB
	tx := db.Delete(type_things)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
