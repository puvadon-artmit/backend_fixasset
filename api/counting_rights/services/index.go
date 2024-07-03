package ServicesCounting_rights

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Counting_rights, Error error) {
	db := database.DB
	types := model.Counting_rights{Counting_rightsID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Counting_rights, err error) {
	db := database.DB
	var records []*model.Counting_rights
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateCounting_rights(types model.Counting_rights) (value *model.Counting_rights, Error error) {
	db := database.DB
	types.Counting_rightsID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateCounting_rights(id string, updatedType model.Counting_rights) (value *model.Counting_rights, Error error) {
	db := database.DB
	existingType := model.Counting_rights{Counting_rightsID: id}
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
	tx := db.Unscoped().Where("counting_rights_id = ?", countingRightsID).Delete(&model.Counting_rights{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAsset_countIDDB(Asset_countID string) ([]*model.Asset_count_Category, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}

func GetUserID(UserID string) ([]*model.Counting_rights, error) {
	db := database.DB

	var User []*model.Counting_rights
	tx := db.Preload(clause.Associations).Where("user_id = ?", UserID).Find(&User)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return User, nil
}
