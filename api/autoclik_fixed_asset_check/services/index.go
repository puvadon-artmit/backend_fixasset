package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Autoclik_Fixed_Asset_Check, Error error) {
	db := database.DB
	autoclik_fixed_asset_check := model.Autoclik_Fixed_Asset_Check{Autoclik_Fixed_Asset_CheckID: id}
	tx := db.Preload(clause.Associations).Find(&autoclik_fixed_asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_fixed_asset_check, nil
}

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetCheck *model.Autoclik_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModel model.Autoclik_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).First(&assetCheckModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &assetCheckModel, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Autoclik_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Autoclik_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return assetCheckModels, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Autoclik_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Autoclik_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Order("item_no desc").Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	propertyCodes := make(map[string]bool)
// 	var filteredAssetChecks []*model.Autoclik_Fixed_Asset_Check

// 	for _, assetCheck := range assetCheckModels {
// 		if !propertyCodes[*assetCheck.Item_no] {
// 			propertyCodes[*assetCheck.Item_no] = true
// 			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
// 		}
// 	}

// 	return filteredAssetChecks, nil
// }

func UpdateAutoclik_Fixed_Asset_CheckByID(id string, updatedItem model.Autoclik_Fixed_Asset_Check) error {
	db := database.DB

	existingItem := model.Autoclik_Fixed_Asset_Check{Autoclik_Fixed_Asset_CheckID: id}
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

func GetAllRecords() (result []*model.Autoclik_Fixed_Asset_Check, err error) {
	db := database.DB
	var records []*model.Autoclik_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CreateNewAutoclik_Fixed_Asset_Check(autoclik_fixed_asset_check model.Autoclik_Fixed_Asset_Check) (value *model.Autoclik_Fixed_Asset_Check, Error error) {
// 	db := database.DB
// 	autoclik_fixed_asset_check.Autoclik_Fixed_Asset_CheckID = uuid.New().String()
// 	tx := db.Create(&autoclik_fixed_asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &autoclik_fixed_asset_check, nil
// }

func CreateNewAutoclik_Fixed_Asset_Check(autoclik_check model.Autoclik_Fixed_Asset_Check, autoclik_photos_check []model.Autoclik_Fixed_Asset_Photos_check) (value *model.Autoclik_Fixed_Asset_Check, Error error) {
	db := database.DB
	autoclik_check.Autoclik_Fixed_Asset_CheckID = uuid.New().String()

	tx := db.Create(&autoclik_check)
	if tx.Error != nil {
		return nil, tx.Error
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	stmt, err := sqlDB.Prepare("INSERT INTO autoclik_fa_all_photos (autoclik_fixed_asset_check_id, autoclik_fixed_asset_photos_check_id) VALUES ($1, $2)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := range autoclik_photos_check {
		_, err := stmt.Exec(autoclik_check.Autoclik_Fixed_Asset_CheckID, autoclik_photos_check[i].Autoclik_Fixed_Asset_Photos_checkID)
		if err != nil {
			return nil, err
		}
	}

	return &autoclik_check, nil
}

func GetByAutoclikRound_CountID(Round_CountID string) ([]*model.Autoclik_Fixed_Asset_Check, error) {
	db := database.DB
	var round_count []*model.Autoclik_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("autoclik_fixed_asset_round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetByItem_noAndAutoclik_Round_CountID(Item_no string, Autoclik_Round_CountID string) ([]*model.Autoclik_Fixed_Asset_Check, error) {
	db := database.DB
	var autoclik []*model.Autoclik_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("no = ?", Item_no).Where("autoclik_fixed_asset_round_count_id = ?", Autoclik_Round_CountID).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik, nil
}

func GetAutoclikCheckByautoclik_round_count_id(Round_CountID string) (assetChecks []*model.Autoclik_Fixed_Asset_Check, err error) {
	db := database.DB
	var assetCheckModels []*model.Autoclik_Fixed_Asset_Check

	tx := db.Where("autoclik_fixed_asset_round_count_id = ?", Round_CountID).Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var filteredAssetChecks []*model.Autoclik_Fixed_Asset_Check

	return filteredAssetChecks, nil
}

func GetByAutoclikRound_CountAllID(Round_CountID string) ([]*model.Autoclik_Fixed_Asset_Check, error) {
	db := database.DB
	var round_count []*model.Autoclik_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("autoclik_fixed_asset_round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetAutoclikCheckByFixedAssetID(Round_CountID string) (assetChecks []*model.Autoclik_Fixed_Asset_Check, err error) {
	db := database.DB
	var assetCheckModels []*model.Autoclik_Fixed_Asset_Check

	tx := db.Where("autoclik_fixed_asset_round_count_id = ?", Round_CountID).Order("no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Autoclik_Fixed_Asset_Check

	for _, assetCheck := range assetCheckModels {
		if !propertyCodes[*assetCheck.No] {
			propertyCodes[*assetCheck.No] = true
			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
		}
	}

	return filteredAssetChecks, nil
}

// func CreateNewAutoclik_Fixed_Asset_Check(autoclik_fixed_asset_check model.Autoclik_Fixed_Asset_Check) (*model.Autoclik_Fixed_Asset_Check, error) {
// 	db := database.DB

// 	assetCheckID := uuid.New()
// 	autoclik_fixed_asset_check.Autoclik_Fixed_Asset_CheckID = assetCheckID.String()

// 	for i := range autoclik_fixed_asset_check.Photos_check {
// 		photosCheckID := uuid.New()
// 		autoclik_fixed_asset_check.Photos_check[i].Photos_checkID = photosCheckID.String()
// 	}

// 	currentTime := time.Now()
// 	autoclik_fixed_asset_check.CreatedAt = &currentTime
// 	autoclik_fixed_asset_check.UpdatedAt = &currentTime

// 	tx := db.Create(&autoclik_fixed_asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &autoclik_fixed_asset_check, nil
// }
