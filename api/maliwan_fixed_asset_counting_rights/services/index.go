package ServicesMaliwan_Fixed_Asset_Counting_Rights

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_Counting_Rights, Error error) {
	db := database.DB
	types := model.Maliwan_Fixed_Asset_Counting_Rights{Maliwan_Fixed_Asset_Counting_RightsID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_Counting_Rights, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_Counting_Rights
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateMaliwan_Fixed_Asset_Counting_Rights(types model.Maliwan_Fixed_Asset_Counting_Rights) (value *model.Maliwan_Fixed_Asset_Counting_Rights, Error error) {
	db := database.DB
	types.Maliwan_Fixed_Asset_Counting_RightsID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateMaliwan_Fixed_Asset_Counting_Rights(id string, updatedType model.Maliwan_Fixed_Asset_Counting_Rights) (value *model.Maliwan_Fixed_Asset_Counting_Rights, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Counting_Rights{Maliwan_Fixed_Asset_Counting_RightsID: id}
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

func DeleteCountingRightsByID(countingRightsID string) error {
	db := database.DB
	tx := db.Unscoped().Where("maliwan_fixed_asset_counting_rights_id = ?", countingRightsID).Delete(&model.Maliwan_Fixed_Asset_Counting_Rights{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAsset_countIDDB(Asset_countID string) ([]*model.Asset_count_Category, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Category
	tx := db.Preload(clause.Associations).Where("maliwan_fixed_asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}

func GetUserID(UserID string) ([]*model.Maliwan_Fixed_Asset_Counting_Rights, error) {
	db := database.DB

	var User []*model.Maliwan_Fixed_Asset_Counting_Rights
	tx := db.Preload(clause.Associations).Where("user_id = ?", UserID).Find(&User)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return User, nil
}
