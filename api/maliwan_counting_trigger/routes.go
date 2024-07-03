package Maliwan_counting_triggerRoutes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	Controllermaliwan_counting_trigger "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_trigger/controllers"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/middleware"
	Model "github.com/puvadon-artmit/gofiber-template/model"
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func SetupaMaliwan_counting_triggerRoutes(router fiber.Router) {
	db := database.DB
	app := fiber.New()

	counting_trigger := router.Group("maliwan-counting-trigger")
	counting_trigger.Get("/get-by-id/:maliwan_counting_trigger_id", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.GetByTrigger_ID)
	counting_trigger.Get("/maliwan-count-data-store/:maliwan_count_id", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.GetByMaliwan_CountID)
	counting_trigger.Get("/maliwan-count-data-store-test/:maliwan_count_id", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.GetByMaliwan_CountIDOnFilterODataV4)
	// counting_trigger.Get("/get-maliwan-data", Controllermaliwan_counting_trigger.GetDatamaliwan)
	counting_trigger.Get("/get-pull-item-bin", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.PullItembinHandler)
	counting_trigger.Get("/get-maliwan-trigger", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.GetAllHandler)

	counting_trigger.Delete("/delete-all-data", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.DeleteItem_bin_maliwan)

	// counting_trigger.Get("/pull-data-by-id/:maliwan_counting_trigger_id", PullDataByBranch)

	counting_trigger.Post("/create-maliwan-counting-trigger", middleware.AuthorizationRequired(), Controllermaliwan_counting_trigger.Create)
	// counting_trigger.Delete("/delete-counting-trigger/:maliwan_counting_trigger_id", Controllermaliwan_counting_trigger.DeleteCountingTrigger)

	// go func() {
	// 	ticker := time.NewTicker(45 * time.Second)
	// 	defer ticker.Stop()
	// 	for range ticker.C {
	// 		checkMaliwanCountingTriggers(time.Now(), app, db)
	// 	}
	// }()

	// go func() {
	// 	ticker := time.NewTicker(45 * time.Second)
	// 	defer ticker.Stop()
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			checkMaliwanCountingTriggers(time.Now(), app, db)
	// 		}
	// 	}
	// }()

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	crontime := cron.New(cron.WithLocation(location))

	// ทุก 15 นาทีในช่วงเวลา 09:00 - 09:59
	// crontime.AddFunc("0,15,30,45 9 * * *", func() {
	// 	fmt.Println("ตามเวลาทุก 15 นาที แล้วจ้าาาา  //รอบ 9 โมง")
	// })

	// crontime.AddFunc("0,15,30,45 11 * * *", func() {
	// 	checkMaliwanCountingTriggers(time.Now(), app, db)
	// })

	// // ทุก 15 นาทีในช่วงเวลา 20:00 - 23:59
	// crontime.AddFunc("0,15,30,45 20-23 * * *", func() {
	// 	checkMaliwanCountingTriggers(time.Now(), app, db)
	// })

	// ทุก 15 นาทีในช่วงเวลา 00:00 - 05:59
	crontime.AddFunc("0,15,30,45 0-5 * * *", func() {
		checkMaliwanCountingTriggers(time.Now(), app, db)
	})

	crontime.AddFunc("32 19 * * *", func() {
		fctx := &fasthttp.RequestCtx{}

		c := app.AcquireCtx(fctx)

		defer app.ReleaseCtx(c)

		if err := Controllermaliwan_counting_trigger.DeleteItem_bin_maliwan(c); err != nil {
			fmt.Println("Error running job:", err)
		}

		time.Sleep(1 * time.Minute)

		fmt.Printf("เริ่ม pull data")

		if err := Controllermaliwan_counting_trigger.PullItembinHandler(c); err != nil {
			fmt.Println("Error running job:", err)
		}
	})

	// เริ่มต้น cron scheduler
	crontime.Start()

	// // ให้โปรแกรมรันต่อเนื่องโดยไม่หยุด
	// select {}

}

func checkMaliwanCountingTriggers(t time.Time, app *fiber.App, db *gorm.DB) {
	if db == nil {
		fmt.Println("Database is not initialized")
		return
	}

	var maliwanCountingTriggers []Model.Maliwan_Counting_Trigger

	db.Where("day_time = ? AND working_time = ?", t.Format("2006-01-02"), t.Format("15:04")).Find(&maliwanCountingTriggers)

	for _, trigger := range maliwanCountingTriggers {
		fmt.Printf("Found matching Maliwan_Counting_Trigger with ID: %s\n", trigger.Maliwan_Counting_TriggerID)

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{}) // Initialize new context with fasthttp.RequestCtx
		defer app.ReleaseCtx(ctx)

		// time.Sleep(2 * time.Minute)

		// err := PullDataByBranch(ctx, trigger.Maliwan_Counting_TriggerID)
		// if err != nil {
		// 	fmt.Printf("Error PullDataByBranch in GetByTrigger_ID: %v\n", err)
		// 	continue
		// }

		// time.Sleep(2 * time.Minute)

		err := GetByTrigger_ID(ctx, trigger.Maliwan_Counting_TriggerID)
		if err != nil {
			fmt.Printf("Error GetByTrigger_ID in GetByTrigger_ID: %v\n", err)
			continue
		}

		time.Sleep(1 * time.Minute)

		err = Controllermaliwan_counting_trigger.DeleteCountingTrigger(ctx, trigger.Maliwan_Counting_TriggerID)

		if err != nil {
			fmt.Printf("Error in GetByTrigger_ID: %v\n", err)
			continue
		}
		fmt.Println("Read Data And Create Successfully")

	}
}

// func GetByTrigger_ID(c *fiber.Ctx, Maliwan_Counting_TriggerID string) error {

// 	maliwanCountingTrigger, err := ServicesMaliwan_Counting_Trigger.GetById(Maliwan_Counting_TriggerID)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	// Ensure we have a valid Maliwan_Counting_Trigger
// 	if maliwanCountingTrigger == nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  "Maliwan_Counting_Trigger not found",
// 		})
// 	}

// 	// Extract Maliwan_countID from the fetched Maliwan_Counting_Trigger
// 	maliwanCountID := maliwanCountingTrigger.Maliwan_countID

// 	// Extract Branch_Maliwan from the associated Maliwan_count
// 	branchMaliwanCount := maliwanCountingTrigger.Maliwan_count.Branch_Maliwan
// 	if branchMaliwanCount == nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  "Branch_Maliwan not found",
// 		})
// 	}

// 	// Fetch Item_Category_Code by Maliwan_countID from Maliwan_Counts_Item_Category_Code table
// 	itemCategoryCodes, err := ServicesMaliwan_Counting_Trigger.GetBymaliwan_count_IDDB(maliwanCountID)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	Gen_prod_posting_Group_Maliwan := maliwanCountingTrigger.Maliwan_count.Gen_Prod_Posting_Group
// 	if Gen_prod_posting_Group_Maliwan == "" {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  "Branch_Maliwan not found",
// 		})
// 	}

// 	// No need to loop through itemCategoryCodes to extract Item_Category_Code again
// 	itemCategoryCodeList := itemCategoryCodes

// 	// Fetch Maliwan_data by item category codes and branch
// 	maliwanData, err := ServicesMaliwan_Counting_Trigger.GetByMaliwan_dataNoDB(itemCategoryCodeList, *branchMaliwanCount, Gen_prod_posting_Group_Maliwan)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	// Prepare the response
// 	response := fiber.Map{
// 		"status": "success",
// 		"result": fiber.Map{
// 			"maliwan_count_id":         maliwanCountID,
// 			"maliwan_counting_trigger": maliwanCountingTrigger,
// 			"item_category_codes":      itemCategoryCodeList,
// 			"maliwan_data":             maliwanData,
// 		},
// 	}

// 	// err = Controllermaliwan_counting_trigger.ReaddataAndCreate(maliwanData, maliwanCountID)
// 	// if err != nil {
// 	// 	return c.Status(500).JSON(fiber.Map{
// 	// 		"status": "error",
// 	// 		"error":  err.Error(),
// 	// 	})
// 	// }

// 	return c.JSON(response)
// }
