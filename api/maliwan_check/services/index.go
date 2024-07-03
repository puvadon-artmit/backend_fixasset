package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_check, Error error) {
	db := database.DB
	maliwan_check := model.Maliwan_check{Maliwan_checkID: id}
	tx := db.Preload(clause.Associations).Find(&maliwan_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_check, nil
}

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetCheck *model.Maliwan_check, err error) {
// 	db := database.DB
// 	var assetCheckModel model.Maliwan_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).First(&assetCheckModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &assetCheckModel, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Maliwan_check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Maliwan_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return assetCheckModels, nil
// }

func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Maliwan_check, err error) {
	db := database.DB
	var assetCheckModels []*model.Maliwan_check

	tx := db.Where("maliwan_round_count_id = ?", Round_CountID).Order("no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Maliwan_check

	for _, assetCheck := range assetCheckModels {
		if !propertyCodes[*assetCheck.No] {
			propertyCodes[*assetCheck.No] = true
			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
		}
	}

	return filteredAssetChecks, nil
}

func UpdateMaliwan_checkByID(id string, updatedItem model.Maliwan_check) error {
	db := database.DB

	existingItem := model.Maliwan_check{Maliwan_checkID: id}
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

func GetByNoAndMaliwan_Round_CountID(No string, Maliwan_Round_CountID string) ([]*model.Maliwan_check, error) {
	db := database.DB
	var maliwan []*model.Maliwan_check
	tx := db.Preload(clause.Associations).Where("no = ?", No).Where("maliwan_round_count_id = ?", Maliwan_Round_CountID).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}

func GetAllRecords() (result []*model.Maliwan_check, err error) {
	db := database.DB
	var records []*model.Maliwan_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CreateNewMaliwan_check(maliwan_check model.Maliwan_check) (value *model.Maliwan_check, Error error) {
// 	db := database.DB
// 	maliwan_check.Maliwan_checkID = uuid.New().String()
// 	tx := db.Create(&maliwan_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &maliwan_check, nil
// }

func CreateNewMaliwan_check(maliwan_check model.Maliwan_check, maliwan_photos_check []model.Maliwan_Photos_check) (value *model.Maliwan_check, Error error) {
	db := database.DB
	maliwan_check.Maliwan_checkID = uuid.New().String()

	tx := db.Create(&maliwan_check)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for i := range maliwan_photos_check {
		photomaliwancheck := model.PhotoMaliwanCheck{
			Maliwan_checkID:        maliwan_check.Maliwan_checkID,
			Maliwan_Photos_checkID: maliwan_photos_check[i].Maliwan_Photos_checkID,
		}
		tx := db.Create(&photomaliwancheck)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &maliwan_check, nil
}

func GetByMaliwan_Round_CountID(Maliwan_Round_CountID string) ([]*model.Maliwan_check, error) {
	db := database.DB
	var round_count []*model.Maliwan_check
	tx := db.Preload(clause.Associations).Where("maliwan_round_count_id = ?", Maliwan_Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetMaliwan_checkByautoclik_round_count_id(Maliwan_Round_CountID string) (assetChecks []*model.Maliwan_check, err error) {
	db := database.DB
	var assetCheckModels []*model.Maliwan_check

	tx := db.Where("maliwan_round_count_id = ?", Maliwan_Round_CountID).Order("no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Maliwan_check

	for _, maliwan_check := range assetCheckModels {
		if !propertyCodes[*maliwan_check.No] {
			propertyCodes[*maliwan_check.No] = true
			filteredAssetChecks = append(filteredAssetChecks, maliwan_check)
		}
	}

	return filteredAssetChecks, nil
}

// func CreateNewMaliwan_check(maliwan_check model.Maliwan_check) (*model.Maliwan_check, error) {
// 	db := database.DB

// 	assetCheckID := uuid.New()
// 	maliwan_check.Maliwan_checkID = assetCheckID.String()

// 	for i := range maliwan_check.Photos_check {
// 		photosCheckID := uuid.New()
// 		maliwan_check.Photos_check[i].Photos_checkID = photosCheckID.String()
// 	}

// 	currentTime := time.Now()
// 	maliwan_check.CreatedAt = &currentTime
// 	maliwan_check.UpdatedAt = &currentTime

// 	tx := db.Create(&maliwan_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &maliwan_check, nil
// }
