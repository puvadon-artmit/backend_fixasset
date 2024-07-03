package Posting_groupsController

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	Posting_groupsvices "github.com/puvadon-artmit/gofiber-template/api/posting_group/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func GetAllHandler(c *fiber.Ctx) error {
	records, err := Posting_groupsvices.GetAllPostingGroups()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func ClearAndPull(c *fiber.Ctx) error {
	err := Posting_groupsvices.ClearAndFetchAllItems()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Autoclik retrieved and created successfully",
	})
}

func GetPosting_group_Maliwan(c *fiber.Ctx) error {
	// Retrieve branchMaliwanCounts from query parameters
	branchMaliwanCounts := c.Query("branchMaliwanCounts")
	if branchMaliwanCounts == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "branchMaliwanCounts parameter is required",
		})
	}

	// Split the comma-separated values into a slice of strings
	counts := strings.Split(branchMaliwanCounts, ",")

	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	var records model.ItemPostingGroupMaliwan
	var result []model.GenProductPostingGroupMaliwan

	// วนลูปเพื่อดึงข้อมูลสำหรับทุกสาขา
	for _, branchMaliwanCount := range counts {
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/GenProductPostingGroupsAPI"

		skip := 0
		top := 100 // Number of records to fetch per request
		for {
			param := fmt.Sprintf("?$top=%d&$skip=%d", top, skip)
			pathURL := NAV_SERVER_HMW + Company + ItemApi + param
			req, err := http.NewRequest("GET", pathURL, nil)
			if err != nil {
				return err
			}

			resp, err := Client.Do(req)
			if err != nil {
				return err
			}
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				resp.Body.Close()
				break
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return err
			}
			err = json.Unmarshal(body, &records)
			if err != nil {
				return err
			}

			if len(records.Value) == 0 {
				break
			}

			result = append(result, records.Value...)
			skip += top
		}
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}
