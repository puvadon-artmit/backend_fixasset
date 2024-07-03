package Autoclik_dataServices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
	"gorm.io/gorm/clause"
)

// var Client *http.Client

func GetPostingGroups() (result []model.ItemPostingGroups, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/GenProductPostingGroupsAPI"
	path_url := NAV_SERVER_AUTOCLIK + ItemAPI
	req, err := http.NewRequest("GET", path_url, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var autoclik model.ItemPostingGroups
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.ItemPostingGroups{autoclik}, nil
}

func CreateNewAutoclik_data(autoclik_data model.GenProductPostingGroups) (value *model.GenProductPostingGroups, Error error) {
	db := database.DB
	autoclik_data.GenProductPostingGroupsID = uuid.New().String()
	tx := db.Create(&autoclik_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_data, nil
}

func GetAllRecordsAndCreate() error {
	values, err := GetPostingGroups()
	if err != nil {
		return err
	}

	for _, item := range values {
		for _, group := range item.Value {
			autoclik_data := model.GenProductPostingGroups{
				GenProductPostingGroupsID:  group.GenProductPostingGroupsID,
				Code:                       group.Code,
				Description:                group.Description,
				Def_VAT_Prod_Posting_Group: group.Def_VAT_Prod_Posting_Group,
				Auto_Insert_Default:        group.Auto_Insert_Default,
			}

			_, err := CreateNewAutoclik_data(autoclik_data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func GetById(id string) (value *model.Item_Autoclik, Error error) {
	db := database.DB
	autoclik := model.Item_Autoclik{Autoclik_dataID: id}
	tx := db.Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik, nil
}

func GetAllPostingGroups() (result []*model.GenProductPostingGroups, err error) {

	db := database.DB
	var autoclik []*model.GenProductPostingGroups
	tx := db.Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return autoclik, nil
}

func ClearAllItems() error {
	db := database.DB
	if err := db.Where("true").Delete(&model.GenProductPostingGroups{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAndFetchAllItems() error {
	// เคลียร์ข้อมูลทั้งหมดจากตาราง
	if err := ClearAllItems(); err != nil {
		return err
	}

	// สร้างข้อมูลใหม่จากข้อมูลที่มีอยู่
	if err := GetAllRecordsAndCreate(); err != nil {
		return err
	}

	return nil
}
