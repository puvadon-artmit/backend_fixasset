package ServicesMain_Category

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetAllRecords() (result []*model.Main_Category, err error) {
	db := database.DB
	var records []*model.Main_Category
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CountMain_Category() (count int, err error) {
	latestRecords, err := GetAllRecords()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func CreateNewMain_Category(inspection model.Main_Category) (value *model.Main_Category, Error error) {
	db := database.DB
	inspection.Main_Category_ID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateMain_Category(id string, updatedType model.Main_Category) (value *model.Main_Category, Error error) {
	db := database.DB
	existingType := model.Main_Category{Main_Category_ID: id}
	tx := db.Model(&existingType).Where("main_category_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &existingType, nil
}

func DeleteMain_Category(main_category *model.Main_Category) error {
	db := database.DB
	tx := db.Delete(main_category)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteMainCategoryByID(Main_Category_ID string) error {
	db := database.DB
	tx := db.Where("main_category_id = ?", Main_Category_ID).Delete(&model.Main_Category{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetById(id string) (value *model.Main_Category, Error error) {
	db := database.DB
	main_category := model.Main_Category{Main_Category_ID: id}
	tx := db.Preload(clause.Associations).Find(&main_category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &main_category, nil
}

func GetByMain_CategoryID(id string) (CategoryImageURL string, err error) {
	db := database.DB
	categoryimage := model.Main_Category{Main_Category_ID: id}
	tx := db.Select("main_category_photo").First(&categoryimage)
	if tx.Error != nil {
		return "", tx.Error
	}
	return categoryimage.Main_Category_Photo, nil
}
