package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Signature_Autoclik_Fixed_Asset, Error error) {
	db := database.DB
	signature_autoclik_fixed_asset := model.Signature_Autoclik_Fixed_Asset{Signature_Autoclik_Fixed_AssetID: id}
	tx := db.Preload(clause.Associations).Find(&signature_autoclik_fixed_asset)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &signature_autoclik_fixed_asset, nil
}

func GetByAutoclik_countID(Autoclik_countID string) ([]*model.Signature_Autoclik_Fixed_Asset, error) {
	db := database.DB
	var round_count []*model.Signature_Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("autoclik_fixed_asset_count_id = ?", Autoclik_countID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if len(round_count) == 0 {
		return nil, errors.New("no records found")
	}
	return round_count, nil
}

func UpdateSignature_Autoclik_Fixed_AssetByID(id string, updatedItem model.Signature_Autoclik_Fixed_Asset) error {
	db := database.DB

	existingItem := model.Signature_Autoclik_Fixed_Asset{Signature_Autoclik_Fixed_AssetID: id}
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

func GetAllRecordsSignature() (result []*model.Signature_Autoclik_Fixed_Asset, err error) {
	db := database.DB
	var records []*model.Signature_Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewSignature_Autoclik_Fixed_Asset(asset_check model.Signature_Autoclik_Fixed_Asset) (value *model.Signature_Autoclik_Fixed_Asset, Error error) {
	db := database.DB
	asset_check.Signature_Autoclik_Fixed_AssetID = uuid.New().String()
	tx := db.Create(&asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &asset_check, nil
}

func GetByAsset_countDB(Asset_countID string) ([]*model.Signature_Autoclik_Fixed_Asset, error) {
	db := database.DB
	var asset_count []*model.Signature_Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("autoclik_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}
