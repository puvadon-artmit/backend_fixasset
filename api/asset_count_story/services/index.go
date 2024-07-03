package ServicesAsset_count_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Asset_count_story, Error error) {
	db := database.DB
	asset_count_story := model.Asset_count_story{Asset_count_storyID: id}
	tx := db.Preload(clause.Associations).Find(&asset_count_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &asset_count_story, nil
}

func GetAllRecords() (result []*model.Asset_count_story, err error) {
	db := database.DB
	var records []*model.Asset_count_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAsset_count_story(inspection model.Asset_count_story) (value *model.Asset_count_story, Error error) {
	db := database.DB

	if inspection.Asset_count_details == "" {
		inspection.Asset_count_details = ""
	}

	inspection.Asset_count_storyID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateAsset_count_story(id string, updatedType model.Asset_count_story) (value *model.Asset_count_story, Error error) {
	db := database.DB
	existingType := model.Asset_count_story{Asset_count_storyID: id}
	tx := db.Model(&existingType).Where("asset_count_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteAsset_count_story(asset_count_story *model.Asset_count_story) error {
	db := database.DB
	tx := db.Delete(asset_count_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
