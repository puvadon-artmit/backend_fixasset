package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Maliwan_Fixed_Asset_Check, Error error) {
	db := database.DB
	maliwan_fixed_asset_check := model.Maliwan_Fixed_Asset_Check{Maliwan_Fixed_Asset_CheckID: id}
	tx := db.Preload(clause.Associations).Find(&maliwan_fixed_asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_fixed_asset_check, nil
}

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetCheck *model.Maliwan_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModel model.Maliwan_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).First(&assetCheckModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &assetCheckModel, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Maliwan_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Maliwan_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return assetCheckModels, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Maliwan_Fixed_Asset_Check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Maliwan_Fixed_Asset_Check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Order("item_no desc").Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	propertyCodes := make(map[string]bool)
// 	var filteredAssetChecks []*model.Maliwan_Fixed_Asset_Check

// 	for _, assetCheck := range assetCheckModels {
// 		if !propertyCodes[*assetCheck.Item_no] {
// 			propertyCodes[*assetCheck.Item_no] = true
// 			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
// 		}
// 	}

// 	return filteredAssetChecks, nil
// }

func UpdateMaliwan_Fixed_Asset_CheckByID(id string, updatedItem model.Maliwan_Fixed_Asset_Check) error {
	db := database.DB

	existingItem := model.Maliwan_Fixed_Asset_Check{Maliwan_Fixed_Asset_CheckID: id}
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

func GetAllRecords() (result []*model.Maliwan_Fixed_Asset_Check, err error) {
	db := database.DB
	var records []*model.Maliwan_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CreateNewMaliwan_Fixed_Asset_Check(maliwan_fixed_asset_check model.Maliwan_Fixed_Asset_Check) (value *model.Maliwan_Fixed_Asset_Check, Error error) {
// 	db := database.DB
// 	maliwan_fixed_asset_check.Maliwan_Fixed_Asset_CheckID = uuid.New().String()
// 	tx := db.Create(&maliwan_fixed_asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &maliwan_fixed_asset_check, nil
// }

func CreateNewMaliwan_Fixed_Asset_Check(maliwan_check model.Maliwan_Fixed_Asset_Check, maliwan_photos_check []model.Maliwan_Fixed_Asset_Photos_check) (value *model.Maliwan_Fixed_Asset_Check, Error error) {
	db := database.DB
	maliwan_check.Maliwan_Fixed_Asset_CheckID = uuid.New().String()

	tx := db.Create(&maliwan_check)
	if tx.Error != nil {
		return nil, tx.Error
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	stmt, err := sqlDB.Prepare("INSERT INTO maliwan_fa_all_photos (maliwan_fixed_asset_check_id, maliwan_fixed_asset_photos_check_id) VALUES ($1, $2)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := range maliwan_photos_check {
		_, err := stmt.Exec(maliwan_check.Maliwan_Fixed_Asset_CheckID, maliwan_photos_check[i].Maliwan_Fixed_Asset_Photos_checkID)
		if err != nil {
			return nil, err
		}
	}

	return &maliwan_check, nil
}

func GetByMaliwanRound_CountID(Round_CountID string) ([]*model.Maliwan_Fixed_Asset_Check, error) {
	db := database.DB
	var round_count []*model.Maliwan_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("maliwan_fixed_asset_round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetByItem_noAndMaliwan_Round_CountID(Item_no string, Maliwan_Round_CountID string) ([]*model.Maliwan_Fixed_Asset_Check, error) {
	db := database.DB
	var maliwan []*model.Maliwan_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("no = ?", Item_no).Where("maliwan_fixed_asset_round_count_id = ?", Maliwan_Round_CountID).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}

func GetMaliwanCheckBymaliwan_round_count_id(Round_CountID string) (assetChecks []*model.Maliwan_Fixed_Asset_Check, err error) {
	db := database.DB
	var assetCheckModels []*model.Maliwan_Fixed_Asset_Check

	tx := db.Where("maliwan_fixed_asset_round_count_id = ?", Round_CountID).Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var filteredAssetChecks []*model.Maliwan_Fixed_Asset_Check

	return filteredAssetChecks, nil
}

func GetByMaliwanRound_CountAllID(Round_CountID string) ([]*model.Maliwan_Fixed_Asset_Check, error) {
	db := database.DB
	var round_count []*model.Maliwan_Fixed_Asset_Check
	tx := db.Preload(clause.Associations).Where("maliwan_fixed_asset_round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

func GetMaliwanCheckByFixedAssetID(Round_CountID string) (assetChecks []*model.Maliwan_Fixed_Asset_Check, err error) {
	db := database.DB
	var assetCheckModels []*model.Maliwan_Fixed_Asset_Check

	tx := db.Where("maliwan_fixed_asset_round_count_id = ?", Round_CountID).Order("no desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Maliwan_Fixed_Asset_Check

	for _, assetCheck := range assetCheckModels {
		if !propertyCodes[*assetCheck.No] {
			propertyCodes[*assetCheck.No] = true
			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
		}
	}

	return filteredAssetChecks, nil
}

// func CreateNewMaliwan_Fixed_Asset_Check(maliwan_fixed_asset_check model.Maliwan_Fixed_Asset_Check) (*model.Maliwan_Fixed_Asset_Check, error) {
// 	db := database.DB

// 	assetCheckID := uuid.New()
// 	maliwan_fixed_asset_check.Maliwan_Fixed_Asset_CheckID = assetCheckID.String()

// 	for i := range maliwan_fixed_asset_check.Photos_check {
// 		photosCheckID := uuid.New()
// 		maliwan_fixed_asset_check.Photos_check[i].Photos_checkID = photosCheckID.String()
// 	}

// 	currentTime := time.Now()
// 	maliwan_fixed_asset_check.CreatedAt = &currentTime
// 	maliwan_fixed_asset_check.UpdatedAt = &currentTime

// 	tx := db.Create(&maliwan_fixed_asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &maliwan_fixed_asset_check, nil
// }
