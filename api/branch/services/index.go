package BranchServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Branch, err error) {
	db := database.DB
	branch := model.Branch{BranchID: id}
	tx := db.Preload(clause.Associations).Find(&branch)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &branch, nil
}

func GetAll() (result []*model.Branch, err error) {
	db := database.DB
	var records []*model.Branch
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CountBranch() (count int, err error) {
	latestRecords, err := GetAll()
	if err != nil {
		return 0, err
	}
	return len(latestRecords), nil
}

func CreateNewBranch(branch *model.Branch, group *model.Group, mainBranch *model.Main_branch) (value *model.Branch, err error) {
	db := database.DB
	branch.BranchID = uuid.New().String()
	branch.Group = *group
	branch.Main_branch = *mainBranch
	tx := db.Create(branch)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return branch, nil
}

func UpdateBranch(id string, updatedType model.Branch, group *model.Group, mainBranch *model.Main_branch) (value *model.Branch, Error error) {
	db := database.DB
	existingType := model.Branch{BranchID: id}
	tx := db.Model(&existingType).Where("branch_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeletebranchbyID(branch *model.Branch) error {
	db := database.DB
	tx := db.Delete(branch)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
