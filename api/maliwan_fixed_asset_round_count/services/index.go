package ServicesMaliwan_Fixed_Asset_Round_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_Round_Count, Error error) {
	db := database.DB
	count_maliwan := model.Maliwan_Fixed_Asset_Round_Count{Maliwan_Fixed_Asset_Round_CountID: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_Round_Count, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_Round_Count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Fixed_Asset_Round_Count(count_maliwan model.Maliwan_Fixed_Asset_Round_Count) (value *model.Maliwan_Fixed_Asset_Round_Count, Error error) {
	db := database.DB
	count_maliwan.Maliwan_Fixed_Asset_Round_CountID = uuid.New().String()
	tx := db.Create(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func UpdateMaliwan_Fixed_Asset_Round_Count(id string, updatedType model.Maliwan_Fixed_Asset_Round_Count) (value *model.Maliwan_Fixed_Asset_Round_Count, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Round_Count{Maliwan_Fixed_Asset_Round_CountID: id}
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

func DeleteMaliwan_Fixed_Asset_Round_Count(typeplan *model.Maliwan_Fixed_Asset_Round_Count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByMaliwan_countIDDB(Maliwan_countID string) ([]*model.Maliwan_Fixed_Asset_Round_Count, error) {
	db := database.DB
	maliwan_fixed_asset_round_count := []*model.Maliwan_Fixed_Asset_Round_Count{}
	tx := db.Preload(clause.Associations).Where("maliwan_fixed_asset_count_id = ?", Maliwan_countID).Find(&maliwan_fixed_asset_round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan_fixed_asset_round_count, nil
}
