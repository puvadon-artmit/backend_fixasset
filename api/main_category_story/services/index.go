package ServicesMain_Category_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetAllRecords() (result []*model.Main_Category_story, err error) {
	db := database.DB
	var records []*model.Main_Category_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMain_Category_story(inspection model.Main_Category_story) (value *model.Main_Category_story, Error error) {
	db := database.DB
	inspection.Main_Category_storyID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateMain_Category_story(id string, updatedType model.Main_Category_story) (value *model.Main_Category_story, Error error) {
	db := database.DB
	existingType := model.Main_Category_story{Main_Category_storyID: id}
	tx := db.Model(&existingType).Where("main_category_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &existingType, nil
}

func DeleteMain_Category_story(main_category_story *model.Main_Category_story) error {
	db := database.DB
	tx := db.Delete(main_category_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteMainCategoryByID(Main_Category_story_ID string) error {
	db := database.DB
	tx := db.Where("main_category_story_id = ?", Main_Category_story_ID).Delete(&model.Main_Category_story{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetById(id string) (value *model.Main_Category_story, Error error) {
	db := database.DB
	main_category_story := model.Main_Category_story{Main_Category_storyID: id}
	tx := db.Preload(clause.Associations).Find(&main_category_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &main_category_story, nil
}

// func GetByMain_Category_storyID(id string) (CategoryImageURL string, err error) {
// 	db := database.DB
// 	categoryimage := model.Main_Category_story{Main_Category_storyID: id}
// 	tx := db.Select("main_category_story_photo").First(&categoryimage)
// 	if tx.Error != nil {
// 		return "", tx.Error
// 	}
// 	return categoryimage.Main_Category_story_Photo, nil
// }
