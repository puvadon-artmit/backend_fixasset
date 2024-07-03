package Maliwan_dataRoutes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	ControllerMaliwan_data "github.com/puvadon-artmit/gofiber-template/api/maliwan_data/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
)

func SetupMaliwan_dataRoutes(router fiber.Router) {
	app := router.Group("maliwan-data")

	// ดูดข้อมูลสินค้ามะลิวัลทั้งหมด
	app.Get("/pull-data-maliwan", middleware.AuthorizationRequired(), ControllerMaliwan_data.Clearandsuckbmaliwan)

	// ดูดข้อมูล Head Office
	app.Get("/pull-data-head-office", middleware.AuthorizationRequired(), ControllerMaliwan_data.ClearandpullHeadOffice)

	// ดูดข้อมูล Naka
	app.Get("/pull-data-naka", middleware.AuthorizationRequired(), ControllerMaliwan_data.ClearandpullNaka)

	// ดูดข้อมูล Burirum
	app.Get("/pull-data-burirum", middleware.AuthorizationRequired(), ControllerMaliwan_data.ClearandpullBurirum)

	// ดูดข้อมูล Mueang Krabi
	app.Get("/pull-data-mueang-krabi", middleware.AuthorizationRequired(), ControllerMaliwan_data.ClearandpullMueang_Krabi)

	// ดูดข้อมูล Surin
	app.Get("/pull-data-surin", middleware.AuthorizationRequired(), ControllerMaliwan_data.ClearandpullSurin)

	// -----------------------------------------------------------------------------------------------------------------

	// app.Get("/retrieve-maliwan", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetAll)

	app.Get("/get-count-maliwan", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetCountAllMaliwan)

	app.Get("/get-item-category-code", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetByItem_Category_CodeHandler)
	app.Get("/get-all", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetAllHandler)

	app.Get("/get-no", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetByNoHandler)

	app.Get("/get-branch", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetByBranch)

	app.Get("/get-branchs", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetByBranchtestHandler) // ทดสอบ

	app.Get("/get-count-by-branch", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetCountMaliwandata)
	app.Get("/get-all-maliwan-data", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetAllDataMaliwanHandler)
	app.Get("/get-maliwan-by-branch", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetCountMaliwanDataByBranch)
	app.Get("/filter-maliwan-data", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetAllGetItem_Category_Code)
	app.Get("/pull-item-category-code", middleware.AuthorizationRequired(), ControllerMaliwan_data.GetFilterRecords)

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	services := fiber.New()
	crontime := cron.New(cron.WithLocation(location))

	crontime.AddFunc("49 19 * * *", func() {
		fctx := &fasthttp.RequestCtx{}

		c := services.AcquireCtx(fctx)

		defer services.ReleaseCtx(c)

		// fmt.Printf("เริ่ม pull data")

		if err := ControllerMaliwan_data.PullDataMaliwanAllBranchHandler(c); err != nil {
			fmt.Println("Error running job:", err)
		}

		if err := ControllerMaliwan_data.CreateMaliwan_Update_Story(c); err != nil {
			fmt.Println("Error running job:", err)
		}
	})

	// เริ่มต้น cron scheduler
	crontime.Start()

	// go func() {
	// 	ticker := time.NewTicker(60000 * time.Second)
	// 	defer ticker.Stop()
	// 	for range ticker.C {
	// 		err := ControllerMaliwan_data.Clearandsuckbmaliwan(nil)
	// 		if err != nil {
	// 			log.Println("Error while executing Clearandsuckbmaliwan:", err)
	// 		} else {
	// 			log.Println("Clearandsuckbmaliwan completed successfully.")
	// 		}
	// 	}
	// }()

}
