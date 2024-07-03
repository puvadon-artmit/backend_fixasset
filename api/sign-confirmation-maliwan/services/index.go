package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Signature_Maliwan, Error error) {
	db := database.DB
	signature_maliwan := model.Signature_Maliwan{Signature_MaliwanID: id}
	tx := db.Preload(clause.Associations).Find(&signature_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &signature_maliwan, nil
}

func UpdateSignature_MaliwanByID(id string, updatedItem model.Signature_Maliwan) error {
	db := database.DB

	existingItem := model.Signature_Maliwan{Signature_MaliwanID: id}
	tx := db.First(&existingItem)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Model(&existingItem).Updates(updatedItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetAllRecordsSignature() (result []*model.Signature_Maliwan, err error) {
	db := database.DB
	var records []*model.Signature_Maliwan
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func GetBySignature_MaliwanID(Signature_MaliwanID string) ([]*model.Signature_Maliwan, error) {
	db := database.DB
	var round_count []*model.Signature_Maliwan
	tx := db.Preload(clause.Associations).Where("maliwan_count_id = ?", Signature_MaliwanID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if len(round_count) == 0 {
		return nil, errors.New("no records found")
	}
	return round_count, nil
}

func CreateNewSignature_Maliwan(asset_check model.Signature_Maliwan) (value *model.Signature_Maliwan, Error error) {
	db := database.DB
	asset_check.Signature_MaliwanID = uuid.New().String()
	tx := db.Create(&asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &asset_check, nil
}

func GetByAsset_countDB(Asset_countID string) ([]*model.Signature_Maliwan, error) {
	db := database.DB
	var asset_count []*model.Signature_Maliwan
	tx := db.Preload(clause.Associations).Where("autoclik_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}
