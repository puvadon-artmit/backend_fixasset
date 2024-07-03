package ServicesAutoclik_count_Store

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Count_Store, Error error) {
	db := database.DB
	count_autoclik_story := model.Autoclik_Count_Store{Autoclik_Count_StoreID: id}
	tx := db.Preload(clause.Associations).Find(&count_autoclik_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik_story, nil
}

func GetByAutoclik_Count_StoreIDDB(Autoclik_Count_StoreID string) ([]model.Autoclik_Count_Store, error) {
	db := database.DB
	var autoclikStores []model.Autoclik_Count_Store
	tx := db.Preload(clause.Associations).Where("autoclik_count_id = ?", Autoclik_Count_StoreID).Find(&autoclikStores)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclikStores, nil
}

func GetAllRecords() (result []*model.Autoclik_Count_Store, err error) {
	db := database.DB
	var records []*model.Autoclik_Count_Store
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_count_Store(count_autoclik_story model.Autoclik_Count_Store) (value *model.Autoclik_Count_Store, Error error) {
	db := database.DB
	count_autoclik_story.Autoclik_Count_StoreID = uuid.New().String()
	tx := db.Create(&count_autoclik_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_autoclik_story, nil
}

func UpdateAutoclik_count_Store(id string, updatedType model.Autoclik_Count_Store) (value *model.Autoclik_Count_Store, Error error) {
	db := database.DB
	existingType := model.Autoclik_Count_Store{Autoclik_Count_StoreID: id}
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

func DeleteAutoclik_count_Store(typeplan *model.Autoclik_Count_Store) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteAutoclik_count_ID(autoclik_count_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("autoclik_count_id = ?", autoclik_count_id).Delete(&model.Autoclik_Count_Store{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
