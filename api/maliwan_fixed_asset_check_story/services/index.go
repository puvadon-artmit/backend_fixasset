package ServicesMaliwan_Fixed_Asset_check_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_check_Story, Error error) {
	db := database.DB
	count_maliwan_story := model.Maliwan_Fixed_Asset_check_Story{Maliwan_Fixed_Asset_check_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan_story, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_check_Story, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_check_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Fixed_Asset_check_Story(count_maliwan_story model.Maliwan_Fixed_Asset_check_Story) (value *model.Maliwan_Fixed_Asset_check_Story, Error error) {
	db := database.DB
	count_maliwan_story.Maliwan_Fixed_Asset_check_StoryID = uuid.New().String()
	tx := db.Create(&count_maliwan_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan_story, nil
}

func UpdateMaliwan_Fixed_Asset_check_Story(id string, updatedType model.Maliwan_Fixed_Asset_check_Story) (value *model.Maliwan_Fixed_Asset_check_Story, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_check_Story{Maliwan_Fixed_Asset_check_StoryID: id}
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

func DeleteMaliwan_Fixed_Asset_check_Story(typeplan *model.Maliwan_Fixed_Asset_check_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
