package ServicesAutoclik_Fixed_Asset_Round_Count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Fixed_Asset_Round_Count_Story, Error error) {
	db := database.DB
	count_autoclik := model.Autoclik_Fixed_Asset_Round_Count_Story{Autoclik_Fixed_Asset_Round_Count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func GetAllRecords() (result []*model.Autoclik_Fixed_Asset_Round_Count_Story, err error) {
	db := database.DB
	var records []*model.Autoclik_Fixed_Asset_Round_Count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_Fixed_Asset_Round_Count_Story(count_autoclik model.Autoclik_Fixed_Asset_Round_Count_Story) (value *model.Autoclik_Fixed_Asset_Round_Count_Story, Error error) {
	db := database.DB
	count_autoclik.Autoclik_Fixed_Asset_Round_Count_StoryID = uuid.New().String()
	tx := db.Create(&count_autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik, nil
}

func UpdateAutoclik_Fixed_Asset_Round_Count_Story(id string, updatedType model.Autoclik_Fixed_Asset_Round_Count_Story) (value *model.Autoclik_Fixed_Asset_Round_Count_Story, Error error) {
	db := database.DB
	existingType := model.Autoclik_Fixed_Asset_Round_Count_Story{Autoclik_Fixed_Asset_Round_Count_StoryID: id}
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

func DeleteAutoclik_Fixed_Asset_Round_Count_Story(typeplan *model.Autoclik_Fixed_Asset_Round_Count_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAutoclik_countIDDB(Autoclik_countID string) ([]*model.Autoclik_Fixed_Asset_Round_Count_Story, error) {
	db := database.DB
	autoclik_fixed_asset_round_count_story := []*model.Autoclik_Fixed_Asset_Round_Count_Story{}
	tx := db.Preload(clause.Associations).Where("autoclik_fixed_asset_count_id = ?", Autoclik_countID).Find(&autoclik_fixed_asset_round_count_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik_fixed_asset_round_count_story, nil
}
