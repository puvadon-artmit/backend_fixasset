package ServicesType_things_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Type_things_story, Error error) {
	db := database.DB
	types := model.Type_things_story{Type_things_storyID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Type_things_story, err error) {
	db := database.DB
	var records []*model.Type_things_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}
func CreateNewType_things_story(types model.Type_things_story) (value *model.Type_things_story, Error error) {
	db := database.DB
	types.Type_things_storyID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateType_things_story(id string, updatedType model.Type_things_story) (value *model.Type_things_story, Error error) {
	db := database.DB
	existingType := model.Type_things_story{Type_things_storyID: id}
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

func DeleteType_things_story(typeplan *model.Type_things_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
