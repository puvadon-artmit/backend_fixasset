package ServicesRequest_update_data

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Request_Update_Data, Error error) {
	db := database.DB
	request_update_data := model.Request_Update_Data{Request_Update_DataID: id}
	tx := db.Preload(clause.Associations).Find(&request_update_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &request_update_data, nil
}

func GetAllRecords() (result []*model.Request_Update_Data, err error) {
	db := database.DB
	var records []*model.Request_Update_Data
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewRequest_update_data(inspection model.Request_Update_Data) (value *model.Request_Update_Data, Error error) {
	db := database.DB
	inspection.Request_Update_DataID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateRequest_update_data(id string, updatedType model.Request_Update_Data) (value *model.Request_Update_Data, Error error) {
	db := database.DB
	existingType := model.Request_Update_Data{Request_Update_DataID: id}
	tx := db.Model(&existingType).Where("request_update_data_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteRequest_update_data(request_update_data *model.Request_Update_Data) error {
	db := database.DB
	tx := db.Delete(request_update_data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
