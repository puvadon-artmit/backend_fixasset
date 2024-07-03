package ServicesMaliwan_Fixed_Asset_count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_count_Story, Error error) {
	db := database.DB
	maliwan_fixed_asset_count_story := model.Maliwan_Fixed_Asset_count_Story{Maliwan_Fixed_Asset_count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&maliwan_fixed_asset_count_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_fixed_asset_count_story, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_count_Story, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Fixed_Asset_count_Story(inspection model.Maliwan_Fixed_Asset_count_Story) (value *model.Maliwan_Fixed_Asset_count_Story, Error error) {
	db := database.DB

	inspection.Maliwan_Fixed_Asset_count_StoryID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateMaliwan_Fixed_Asset_count_Story(id string, updatedType model.Maliwan_Fixed_Asset_count_Story) (value *model.Maliwan_Fixed_Asset_count_Story, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_count_Story{Maliwan_Fixed_Asset_count_StoryID: id}
	tx := db.Model(&existingType).Where("maliwan_fixed_asset_count_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteMaliwan_Fixed_Asset_count_Story(maliwan_fixed_asset_count_story *model.Maliwan_Fixed_Asset_count_Story) error {
	db := database.DB
	tx := db.Delete(maliwan_fixed_asset_count_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
