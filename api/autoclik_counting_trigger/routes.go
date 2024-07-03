package autoclik_counting_triggerRoutes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	Controllerautoclik_counting_trigger "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_trigger/controllers"
	ServicesAutoclik_Counting_Trigger "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_trigger/services"
	ControllerAutoclik_data "github.com/puvadon-artmit/gofiber-template/api/autoclik_data/controllers"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/middleware"
	Model "github.com/puvadon-artmit/gofiber-template/model"
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func SetupaAtoclik_counting_triggerRoutes(router fiber.Router) {
	db := database.DB
	app := fiber.New()

	counting_trigger := router.Group("autoclik-counting-trigger")
	counting_trigger.Get("/get-autoclik-trigger", middleware.AuthorizationRequired(), Controllerautoclik_counting_trigger.GetAllHandler)
	counting_trigger.Get("/get-by-id/:autoclik_counting_trigger_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_trigger.GetById)
	counting_trigger.Get("/get-by-trigger-id/:autoclik_counting_trigger_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_trigger.GetByIdAutoclik_Count)

	counting_trigger.Get("/autoclik-count-data-store/:autoclik_count_id", middleware.AuthorizationRequired(), Controllerautoclik_counting_trigger.GetByAutoclik_Count_ID)

	counting_trigger.Get("/by-id/:autoclik_count_id", Controllerautoclik_counting_trigger.GetByAutoclik_Count_ID2)

	// ทดสอบ
	counting_trigger.Get("/get-itembin", Controllerautoclik_counting_trigger.GetItem_binJSonData)
	counting_trigger.Get("/get-data2/:autoclik_count_id", Controllerautoclik_counting_trigger.GetByAutoclik_Count_ID2)
	counting_trigger.Get("/get-data2/:autoclik_count_id", Controllerautoclik_counting_trigger.GetByAutoclik_Count_ID2)

	counting_trigger.Post("/create-autoclik-counting-trigger", middleware.AuthorizationRequired(), Controllerautoclik_counting_trigger.Create)
	// counting_trigger.Delete("/delete-counting-trigger/:autoclik_counting_trigger_id", Controllerautoclik_counting_trigger.DeleteCountingTrigger)

	counting_trigger.Get("/get-product-groups", Controllerautoclik_counting_trigger.GetByProductGroupHandler)

	// go func() {
	// 	ticker := time.NewTicker(45 * time.Second)
	// 	defer ticker.Stop()
	// 	for range ticker.C {
	// 		checkAutoclikCountingTriggers(time.Now(), app, db)
	// 	}
	// }()

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	c := cron.New(cron.WithLocation(location))

	// ทุก 15 นาทีในช่วงเวลา 09:00 - 09:59
	// c.AddFunc("0,15,30,45 9 * * *", func() {
	// 	fmt.Println("ตามเวลาทุก 15 นาที แล้วจ้าาาา  //รอบ 9 โมง")
	// })

	// c.AddFunc("0,15,30,45 10 * * *", func() {
	// 	fmt.Println("ตามเวลาทุก 15 นาที แล้วจ้าาาา // รอบ 10 โมง")
	// })

	// // ทุก 15 นาทีในช่วงเวลา 20:00 - 23:59
	// c.AddFunc("0,15,30,45 20-23 * * *", func() {
	// 	checkAutoclikCountingTriggers(time.Now(), app, db)
	// })

	// ทุก 15 นาทีในช่วงเวลา 00:00 - 05:59
	c.AddFunc("0,15,30,45 0-5 * * *", func() {
		checkAutoclikCountingTriggers(time.Now(), app, db)
	})

	c.Start()

}

func checkAutoclikCountingTriggers(t time.Time, app *fiber.App, db *gorm.DB) {
	if db == nil {
		fmt.Println("Database is not initialized")
		return
	}

	var autoclikCountingTriggers []Model.Autoclik_Counting_Trigger

	db.Where("day_time = ? AND working_time = ?", t.Format("2006-01-02"), t.Format("15:04")).Find(&autoclikCountingTriggers)

	for _, trigger := range autoclikCountingTriggers {
		fmt.Printf("Found matching Autoclik_Counting_Trigger with ID: %s\n", trigger.Autoclik_Counting_TriggerID)

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(ctx)

		err := ControllerAutoclik_data.GetAll(ctx)
		if err != nil {
			fmt.Printf("Error in ControllerAutoclik_data.GetAll: %v\n", err)
			continue
		}

		err = ControllerAutoclik_data.Clearandsuckbinitem(ctx)
		if err != nil {
			fmt.Printf("Error in ControllerAutoclik_data.Clearandsuckbinitem: %v\n", err)
			continue
		}

		time.Sleep(2 * time.Minute)

		err = GetByIdAutoclik_Count(ctx, trigger.Autoclik_Counting_TriggerID)
		if err != nil {
			fmt.Printf("Error in GetByIdAutoclik_Count :  %v\n", err)
			continue
		}

		time.Sleep(1 * time.Minute)

		err = Controllerautoclik_counting_trigger.DeleteCountingTrigger(ctx, trigger.Autoclik_Counting_TriggerID)

		if err != nil {
			fmt.Printf("Error in DeleteCountingTrigger :  %v\n", err)
			continue
		}
		fmt.Println("Read Data AndCreate Successfully")

	}
}

func GetByIdAutoclik_Count(c *fiber.Ctx, Autoclik_Counting_TriggerID string) error {

	autoclikCountingTriggerID := Autoclik_Counting_TriggerID

	autoclikCountingTrigger, err := ServicesAutoclik_Counting_Trigger.GetByIdAutoclik_CountingArray(autoclikCountingTriggerID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if autoclikCountingTrigger != nil {

		Autoclik_Count_ID := autoclikCountingTrigger.Autoclik_count.Autoclik_countID
		branchAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByBranchIDDB(autoclikCountingTrigger.Autoclik_count.BranchAutoclik_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		product_groupAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByAutoclik_countIDDB(Autoclik_Count_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Extract required data
		genProdPostingGroup := autoclikCountingTrigger.Autoclik_count.Gen_Prod_Posting_Group
		branchCode := branchAutoclik[0].Branch_Autoclik.Branch_Code
		if branchCode == "" {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch code is empty",
			})
		}

		var nameProductGroups []string
		for _, pg := range product_groupAutoclik {
			if pg.Name_Product_Group != nil {
				nameProductGroups = append(nameProductGroups, *pg.Name_Product_Group)
			}
		}

		// Call GetByProductGroupHandler with the extracted genProdPostingGroup
		genProdPostingGroupResults, err := ServicesAutoclik_Counting_Trigger.GetByGen_Prod_Posting_GroupDB(genProdPostingGroup)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		productGroups, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(nameProductGroups, branchCode) // ใช้ branchCode แทน genProdPostingGroup
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Filter Product_Groups based on the condition Item_No == No
		var filteredProductGroups []Model.Item_Autoclik_Bin_Code
		for _, pg := range productGroups {
			if pg.Item_No != nil {
				for _, gpg := range genProdPostingGroupResults {
					if gpg.No != nil && *pg.Item_No == *gpg.No {
						filteredProductGroups = append(filteredProductGroups, *pg)
					}
				}
			}
		}

		response := fiber.Map{
			"status": "success",
			"result": fiber.Map{
				"autoclik_count_id": Autoclik_Count_ID,
				"product_group":     filteredProductGroups, // ใช้ filteredProductGroups ที่กรองแล้ว
			},
		}

		err = Controllerautoclik_counting_trigger.ReaddataAndCreate(filteredProductGroups, Autoclik_Count_ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(response)
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"error":  "Data not found",
	})
}
