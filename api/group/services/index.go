package GroupServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Group, err error) {
	db := database.DB
	group := model.Group{GroupID: id}
	tx := db.Preload(clause.Associations).Find(&group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group, nil
}

func GetAll() (result []*model.Group, err error) {
	db := database.DB
	var records []*model.Group
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func GetAll() (value *[]GroupResult.GetAllResult, err error) {
// 	db := database.DB
// 	var result []GroupResult.GetAllResult
// 	tx := db.Table("groups").Select("group_id, group_name, comment, created_at, user_id").Scan(&result)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &result, nil
// }

func CountLatestRecords() (count int, err error) {
	latestRecords, err := GetAll()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func UpdateGroup(id string, updatedType model.Group) (value *model.Group, Error error) {
	db := database.DB
	existingType := model.Group{GroupID: id}
	tx := db.Model(&existingType).Where("group_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func CreateNewGroup(group *model.Group) (value *model.Group, err error) {
	db := database.DB
	group.GroupID = uuid.New().String()
	tx := db.Create(group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return group, nil
}

func DeleteGroupbyID(group *model.Group) error {
	db := database.DB
	tx := db.Delete(group)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
