package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (photosURL string, err error) {
	db := database.DB
	maliwan_photos_check := model.Maliwan_Photos_check{Maliwan_Photos_checkID: id}
	tx := db.Select("photos").First(&maliwan_photos_check)
	if tx.Error != nil {
		return "", tx.Error
	}
	return maliwan_photos_check.Photos, nil
}

func GetByPhotoId(id string) (value *model.Maliwan_Photos_check, Error error) {
	db := database.DB
	getphoto := model.Maliwan_Photos_check{Maliwan_Photos_checkID: id}
	tx := db.Preload(clause.Associations).Find(&getphoto)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getphoto, nil
}

func GetAll() (result []*model.Maliwan_Photos_check, err error) {
	db := database.DB
	var records []*model.Maliwan_Photos_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Photos_check(maliwan_photos_check model.Maliwan_Photos_check) (value *model.Maliwan_Photos_check, Error error) {
	db := database.DB
	maliwan_photos_check.Maliwan_Photos_checkID = uuid.New().String()
	tx := db.Create(&maliwan_photos_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_photos_check, nil
}

func DeleteMaliwan_Photos_check(maliwan_photos_check *model.Maliwan_Photos_check) error {
	db := database.DB
	tx := db.Unscoped().Delete(maliwan_photos_check)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
