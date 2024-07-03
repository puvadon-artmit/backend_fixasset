package StatusServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Status, err error) {
	db := database.DB
	status := model.Status{StatusID: id}
	tx := db.Preload(clause.Associations).Find(&status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &status, nil
}

// func GetAll() (value *[]StatusResult.GetAllResult, err error) {
// 	db := database.DB
// 	var result []StatusResult.GetAllResult
// 	tx := db.Table("statuses").Select("status_id, status_name, comment, created_at, user_id").Scan(&result)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &result, nil
// }

func GetAll() (result []*model.Status, err error) {
	db := database.DB
	var records []*model.Status
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewStatus(status *model.Status) (value *model.Status, err error) {
	db := database.DB
	status.StatusID = uuid.New().String()
	tx := db.Create(status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return status, nil
}

func CountLatestRecords() (count int, err error) {
	latestRecords, err := GetAll()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func UpdateStatus(id string, updatedType model.Status) (value *model.Status, Error error) {
	db := database.DB
	existingType := model.Status{StatusID: id}
	tx := db.Model(&existingType).Where("status_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteStatusID(status *model.Status) error {
	db := database.DB
	tx := db.Delete(status)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
