package ServicesAutoclik_count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_count_Story, Error error) {
	db := database.DB
	count_autoclik_story := model.Autoclik_count_Story{Autoclik_count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik_story, nil
}

func GetAllRecords() (result []*model.Autoclik_count_Story, err error) {
	db := database.DB
	var records []*model.Autoclik_count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_count_Story(count_autoclik_story model.Autoclik_count_Story) (value *model.Autoclik_count_Story, Error error) {
	db := database.DB
	count_autoclik_story.Autoclik_count_StoryID = uuid.New().String()
	tx := db.Create(&count_autoclik_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik_story, nil
}

func UpdateAutoclik_count_Story(id string, updatedType model.Autoclik_count_Story) (value *model.Autoclik_count_Story, Error error) {
	db := database.DB
	existingType := model.Autoclik_count_Story{Autoclik_count_StoryID: id}
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

func DeleteAutoclik_count_Story(typeplan *model.Autoclik_count_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
