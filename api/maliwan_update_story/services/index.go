package ServicesMaliwan_Update_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Update_Story, Error error) {
	db := database.DB
	inspection_story := model.Maliwan_Update_Story{Maliwan_Update_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func GetAllRecords() (result []*model.Maliwan_Update_Story, err error) {
	db := database.DB
	var records []*model.Maliwan_Update_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Update_Story(inspection_story model.Maliwan_Update_Story) (value *model.Maliwan_Update_Story, Error error) {
	db := database.DB
	inspection_story.Maliwan_Update_StoryID = uuid.New().String()
	tx := db.Create(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func UpdateMaliwan_Update_Story(id string, updatedType model.Maliwan_Update_Story) (value *model.Maliwan_Update_Story, Error error) {
	db := database.DB
	existingType := model.Maliwan_Update_Story{Maliwan_Update_StoryID: id}
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

func DeleteMaliwan_Update_Story(typeplan *model.Maliwan_Update_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
