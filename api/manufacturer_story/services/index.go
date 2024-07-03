package ServicesManufacturer_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Manufacturer_Story, Error error) {
	db := database.DB
	manufacturer_story := model.Manufacturer_Story{Manufacturer_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&manufacturer_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer_story, nil
}

func GetAllRecords() (result []*model.Manufacturer_Story, err error) {
	db := database.DB
	var records []*model.Manufacturer_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewManufacturer_story(manufacturer_story model.Manufacturer_Story) (value *model.Manufacturer_Story, Error error) {
	db := database.DB
	manufacturer_story.Manufacturer_StoryID = uuid.New().String()
	tx := db.Create(&manufacturer_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer_story, nil
}
