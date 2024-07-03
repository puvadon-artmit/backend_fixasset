package ServicesMaliwan_Fixed_Asset_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_Count, Error error) {
	db := database.DB
	count_maliwan := model.Maliwan_Fixed_Asset_Count{Maliwan_Fixed_Asset_CountID: id}
	tx := db.Preload(clause.Associations).Find(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_Count, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_Count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMaliwan_Fixed_Asset_Count(count_maliwan model.Maliwan_Fixed_Asset_Count) (value *model.Maliwan_Fixed_Asset_Count, Error error) {
	db := database.DB
	count_maliwan.Maliwan_Fixed_Asset_CountID = uuid.New().String()
	tx := db.Create(&count_maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &count_maliwan, nil
}

func UpdateMaliwan_Fixed_Asset_Count(id string, updatedType model.Maliwan_Fixed_Asset_Count) (value *model.Maliwan_Fixed_Asset_Count, Error error) {
	db := database.DB
	existingType := model.Maliwan_Fixed_Asset_Count{Maliwan_Fixed_Asset_CountID: id}
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

func DeleteMaliwan_Fixed_Asset_Count(typeplan *model.Maliwan_Fixed_Asset_Count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// func CreateNewMaliwan_Fixed_Asset_Count_Store(count_maliwan model.Maliwan_Fixed_Asset_Store) (value *model.Maliwan_Fixed_Asset_Store, Error error) {
// 	db := database.DB
// 	count_maliwan.Maliwan_Fixed_Asset_StoreId = uuid.New().String()
// 	tx := db.Create(&count_maliwan)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &count_maliwan, nil
// }
