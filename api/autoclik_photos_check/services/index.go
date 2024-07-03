package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (photosURL string, err error) {
	db := database.DB
	autoclik_photos_check := model.Autoclik_Photos_check{Autoclik_Photos_checkID: id}
	tx := db.Select("photos").First(&autoclik_photos_check)
	if tx.Error != nil {
		return "", tx.Error
	}
	return autoclik_photos_check.Photos, nil
}

func GetByPhotoId(id string) (value *model.Autoclik_Photos_check, Error error) {
	db := database.DB
	getphoto := model.Autoclik_Photos_check{Autoclik_Photos_checkID: id}
	tx := db.Preload(clause.Associations).Find(&getphoto)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getphoto, nil
}

func GetAll() (result []*model.Autoclik_Photos_check, err error) {
	db := database.DB
	var records []*model.Autoclik_Photos_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAutoclik_Photos_check(autoclik_photos_check model.Autoclik_Photos_check) (value *model.Autoclik_Photos_check, Error error) {
	db := database.DB
	autoclik_photos_check.Autoclik_Photos_checkID = uuid.New().String()
	tx := db.Create(&autoclik_photos_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_photos_check, nil
}

func DeleteAutoclik_Photos_check(autoclik_photos_check *model.Autoclik_Photos_check) error {
	db := database.DB
	tx := db.Unscoped().Delete(autoclik_photos_check)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
