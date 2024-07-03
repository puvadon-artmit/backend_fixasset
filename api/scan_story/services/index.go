package ServicesScan_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Scan_story, Error error) {
	db := database.DB
	scan_story := model.Scan_story{Scan_storyID: id}
	tx := db.Preload(clause.Associations).Find(&scan_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &scan_story, nil
}

func GetAllRecords() (result []*model.Scan_story, err error) {
	db := database.DB
	var records []*model.Scan_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewScan_story(inspection model.Scan_story) (value *model.Scan_story, Error error) {
	db := database.DB

	// if inspection.Asset_count_details == "" {
	// 	inspection.Asset_count_details = ""
	// }

	inspection.Scan_storyID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateScan_story(id string, updatedType model.Scan_story) (value *model.Scan_story, Error error) {
	db := database.DB
	existingType := model.Scan_story{Scan_storyID: id}
	tx := db.Model(&existingType).Where("scan_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteScan_story(scan_story *model.Scan_story) error {
	db := database.DB
	tx := db.Delete(scan_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
