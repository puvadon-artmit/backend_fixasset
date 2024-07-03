package ServicesMaliwan_counts_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_counts_story, Error error) {
	db := database.DB
	count_autoclik := model.Maliwan_counts_story{Maliwan_counts_storyID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func GetAllRecords() (result []*model.Maliwan_counts_story, err error) {
	db := database.DB
	var records []*model.Maliwan_counts_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_counts_story(count_autoclik model.Maliwan_counts_story) (value *model.Maliwan_counts_story, Error error) {
	db := database.DB
	count_autoclik.Maliwan_counts_storyID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateMaliwan_counts_story(id string, updatedType model.Maliwan_counts_story) (value *model.Maliwan_counts_story, Error error) {
	db := database.DB
	existingType := model.Maliwan_counts_story{Maliwan_counts_storyID: id}
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

func DeleteMaliwan_counts_story(typeplan *model.Maliwan_counts_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
