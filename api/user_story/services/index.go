package ServicesUser_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.User_story, Error error) {
	db := database.DB
	user := model.User_story{User_storyID: id}
	tx := db.Preload(clause.Associations).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func GetAllRecords() (result []*model.User_story, err error) {
	db := database.DB
	var records []*model.User_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}
func CreateNewUser_story(user model.User_story) (value *model.User_story, Error error) {
	db := database.DB
	user.User_storyID = uuid.New().String()
	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func UpdateUser_story(id string, updatedType model.User_story) (value *model.User_story, Error error) {
	db := database.DB
	existingType := model.User_story{User_storyID: id}
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

func DeleteUser_story(user *model.User_story) error {
	db := database.DB
	tx := db.Delete(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
