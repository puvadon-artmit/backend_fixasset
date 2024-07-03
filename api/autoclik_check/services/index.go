package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_check, Error error) {
	db := database.DB
	autoclik_check := model.Autoclik_check{Autoclik_checkID: id}
	tx := db.Preload(clause.Associations).Find(&autoclik_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_check, nil
}

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetCheck *model.Autoclik_check, err error) {
// 	db := database.DB
// 	var assetCheckModel model.Autoclik_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).First(&assetCheckModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &assetCheckModel, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Autoclik_check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Autoclik_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return assetCheckModels, nil
// }

func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Autoclik_check, err error) {
	db := database.DB
	var assetCheckModels []*model.Autoclik_check

	tx := db.Where("round_count_id = ?", Round_CountID).Order("item_no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Autoclik_check

	for _, assetCheck := range assetCheckModels {
		if !propertyCodes[*assetCheck.Item_no] {
			propertyCodes[*assetCheck.Item_no] = true
			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
		}
	}

	return filteredAssetChecks, nil
}

func UpdateAutoclik_checkByID(id string, updatedItem model.Autoclik_check) error {
	db := database.DB

	existingItem := model.Autoclik_check{Autoclik_checkID: id}
	tx := db.First(&existingItem)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Model(&existingItem).Updates(updatedItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetAllRecords() (result []*model.Autoclik_check, err error) {
	db := database.DB
	var records []*model.Autoclik_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CreateNewAutoclik_check(autoclik_check model.Autoclik_check) (value *model.Autoclik_check, Error error) {
// 	db := database.DB
// 	autoclik_check.Autoclik_checkID = uuid.New().String()
// 	tx := db.Create(&autoclik_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &autoclik_check, nil
// }

func CreateNewAutoclik_check(autoclik_check model.Autoclik_check, autoclik_photos_check []model.Autoclik_Photos_check) (value *model.Autoclik_check, Error error) {
	db := database.DB
	autoclik_check.Autoclik_checkID = uuid.New().String()

	tx := db.Create(&autoclik_check)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for i := range autoclik_photos_check {
		autoclik_allphoto := model.Autoclik_AllPhoto{
			Autoclik_checkID:        autoclik_check.Autoclik_checkID,
			Autoclik_Photos_checkID: autoclik_photos_check[i].Autoclik_Photos_checkID,
		}
		tx := db.Create(&autoclik_allphoto)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &autoclik_check, nil
}

func GetByAutoclikRound_CountID(Round_CountID string) ([]*model.Autoclik_check, error) {
	db := database.DB
	var round_count []*model.Autoclik_check
	tx := db.Preload(clause.Associations).Where("autoclik_round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetByItem_noAndAutoclik_Round_CountID(Item_no string, Autoclik_Round_CountID string) ([]*model.Autoclik_check, error) {
	db := database.DB
	var autoclik []*model.Autoclik_check
	tx := db.Preload(clause.Associations).Where("item_no = ?", Item_no).Where("autoclik_round_count_id = ?", Autoclik_Round_CountID).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik, nil
}

func GetAutoclikCheckByautoclik_round_count_id(Round_CountID string) (assetChecks []*model.Autoclik_check, err error) {
	db := database.DB
	var assetCheckModels []*model.Autoclik_check

	tx := db.Where("autoclik_round_count_id = ?", Round_CountID).Order("item_no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Autoclik_check

	for _, autoclikCheck := range assetCheckModels {
		if !propertyCodes[*autoclikCheck.Item_no] {
			propertyCodes[*autoclikCheck.Item_no] = true
			filteredAssetChecks = append(filteredAssetChecks, autoclikCheck)
		}
	}

	return filteredAssetChecks, nil
}

// func CreateNewAutoclik_check(autoclik_check model.Autoclik_check) (*model.Autoclik_check, error) {
// 	db := database.DB

// 	assetCheckID := uuid.New()
// 	autoclik_check.Autoclik_checkID = assetCheckID.String()

// 	for i := range autoclik_check.Photos_check {
// 		photosCheckID := uuid.New()
// 		autoclik_check.Photos_check[i].Photos_checkID = photosCheckID.String()
// 	}

// 	currentTime := time.Now()
// 	autoclik_check.CreatedAt = &currentTime
// 	autoclik_check.UpdatedAt = &currentTime

// 	tx := db.Create(&autoclik_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &autoclik_check, nil
// }
