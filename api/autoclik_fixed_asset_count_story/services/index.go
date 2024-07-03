package ServicesAutoclik_Fixed_Asset_count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Fixed_Asset_count_Story, Error error) {
	db := database.DB
	autoclik_fixed_asset_count_story := model.Autoclik_Fixed_Asset_count_Story{Autoclik_Fixed_Asset_count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&autoclik_fixed_asset_count_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_fixed_asset_count_story, nil
}

func GetAllRecords() (result []*model.Autoclik_Fixed_Asset_count_Story, err error) {
	db := database.DB
	var records []*model.Autoclik_Fixed_Asset_count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_Fixed_Asset_count_Story(inspection model.Autoclik_Fixed_Asset_count_Story) (value *model.Autoclik_Fixed_Asset_count_Story, Error error) {
	db := database.DB

	inspection.Autoclik_Fixed_Asset_count_StoryID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateAutoclik_Fixed_Asset_count_Story(id string, updatedType model.Autoclik_Fixed_Asset_count_Story) (value *model.Autoclik_Fixed_Asset_count_Story, Error error) {
	db := database.DB
	existingType := model.Autoclik_Fixed_Asset_count_Story{Autoclik_Fixed_Asset_count_StoryID: id}
	tx := db.Model(&existingType).Where("autoclik_fixed_asset_count_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteAutoclik_Fixed_Asset_count_Story(autoclik_fixed_asset_count_story *model.Autoclik_Fixed_Asset_count_Story) error {
	db := database.DB
	tx := db.Delete(autoclik_fixed_asset_count_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
