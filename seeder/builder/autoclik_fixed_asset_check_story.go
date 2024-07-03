package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAutoclik_Fixed_Asset_check_Story(db *gorm.DB, Autoclik_Fixed_Asset_check_Name string, Autoclik_Fixed_Asset_check_details string,
	No string, UserID string, Autoclik_Fixed_Asset_Round_CountID string) error {
	status := &model.Autoclik_Fixed_Asset_check_Story{
		Autoclik_Fixed_Asset_check_StoryID: uuid.New().String(),
		Autoclik_Fixed_Asset_check_Name:    Autoclik_Fixed_Asset_check_Name,
		Autoclik_Fixed_Asset_check_details: Autoclik_Fixed_Asset_check_details,
		No:                                 &No,
		UserID:                             UserID,
		Autoclik_Fixed_Asset_Round_CountID: Autoclik_Fixed_Asset_Round_CountID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
