package ServicesStatus_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Status_story, Error error) {
	db := database.DB
	types := model.Status_story{Status_story_ID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Status_story, err error) {
	db := database.DB
	var records []*model.Status_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}
func CreateNewStatus_story(types model.Status_story) (value *model.Status_story, Error error) {
	db := database.DB
	types.Status_story_ID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateStatus_story(id string, updatedType model.Status_story) (value *model.Status_story, Error error) {
	db := database.DB
	existingType := model.Status_story{Status_story_ID: id}
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

func DeleteStatus_story(typeplan *model.Status_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
