package ServicesMaliwan_Round_Count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Round_Count_Story, Error error) {
	db := database.DB
	inspection_story := model.Maliwan_Round_Count_Story{Maliwan_Round_Count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func GetAllRecords() (result []*model.Maliwan_Round_Count_Story, err error) {
	db := database.DB
	var records []*model.Maliwan_Round_Count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Round_Count_Story(inspection_story model.Maliwan_Round_Count_Story) (value *model.Maliwan_Round_Count_Story, Error error) {
	db := database.DB
	inspection_story.Maliwan_Round_Count_StoryID = uuid.New().String()
	tx := db.Create(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func UpdateMaliwan_Round_Count_Story(id string, updatedType model.Maliwan_Round_Count_Story) (value *model.Maliwan_Round_Count_Story, Error error) {
	db := database.DB
	existingType := model.Maliwan_Round_Count_Story{Maliwan_Round_Count_StoryID: id}
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

func DeleteMaliwan_Round_Count_Story(typeplan *model.Maliwan_Round_Count_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
