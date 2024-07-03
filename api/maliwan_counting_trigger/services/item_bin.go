package ServicesMaliwan_Counting_Trigger

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetAllbranchsMaliwan(branchs string, item_no []string) ([]*model.Item_bin_maliwan, error) {
	db := database.DB
	var maliwan []*model.Item_bin_maliwan
	tx := db.Preload(clause.Associations).Where("branch = ? AND item_no IN (?)", branchs, item_no).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}

func CreateMaliwan_Count_StoreBatch(maliwanItems []model.Maliwan_Count_Store) error {
	db := database.DB

	// ตั้งค่า Maliwan_Count_StoreID สำหรับแต่ละ element ใน slice
	for i := range maliwanItems {
		maliwanItems[i].Maliwan_Count_StoreID = uuid.New().String()
	}

	tx := db.Create(&maliwanItems)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func CreateMaliwan_Count_Store(maliwan_data model.Maliwan_Count_Store) (value *model.Maliwan_Count_Store, Error error) {
	db := database.DB
	maliwan_data.Maliwan_Count_StoreID = uuid.New().String()

	tx := db.Create(&maliwan_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_data, nil
}

func DeleteAllItem_bin_maliwan() error {
	db := database.DB
	result := db.Exec("DELETE FROM item_bin_maliwans")
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("ลบข้อมูลทั้งหมดเสร็จสิ้น")
	return nil
}
