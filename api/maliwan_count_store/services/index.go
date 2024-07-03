package ServicesMaliwan_count_Store

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Count_Store, Error error) {
	db := database.DB
	count_maliwan_story := model.Maliwan_Count_Store{Maliwan_Count_StoreID: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan_story, nil
}

func GetByMaliwan_Count_StoreIDDB(Maliwan_Count_StoreID string) ([]model.Maliwan_Count_Store, error) {
	db := database.DB
	var maliwanStores []model.Maliwan_Count_Store
	tx := db.Preload(clause.Associations).Where("maliwan_count_id = ?", Maliwan_Count_StoreID).Find(&maliwanStores)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwanStores, nil
}

func GetAllRecords() (result []*model.Maliwan_Count_Store, err error) {
	db := database.DB
	var records []*model.Maliwan_Count_Store
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_count_Store(count_maliwan_story model.Maliwan_Count_Store) (value *model.Maliwan_Count_Store, Error error) {
	db := database.DB
	count_maliwan_story.Maliwan_Count_StoreID = uuid.New().String()
	tx := db.Create(&count_maliwan_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan_story, nil
}

func UpdateMaliwan_count_Store(id string, updatedType model.Maliwan_Count_Store) (value *model.Maliwan_Count_Store, Error error) {
	db := database.DB
	existingType := model.Maliwan_Count_Store{Maliwan_Count_StoreID: id}
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

func DeleteMaliwan_count_Store(typeplan *model.Maliwan_Count_Store) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByItem_Maliwan_BinNoDB(No string) ([]*model.Item_bin_maliwan, error) {
	db := database.DB
	var maliwan []*model.Item_bin_maliwan
	tx := db.Preload(clause.Associations).Where("item_no = ?", No).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}

func DeleteMaliwan_count_ID(maliwan_count_id string) error {
	db := database.DB
	tx := db.Unscoped().Where("maliwan_count_id = ?", maliwan_count_id).Delete(&model.Maliwan_Count_Store{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
