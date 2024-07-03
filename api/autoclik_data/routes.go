package Autoclik_dataRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAutoclik_data "github.com/puvadon-artmit/gofiber-template/api/autoclik_data/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupAutoclik_dataRoutes(router fiber.Router) {
	app := router.Group("autoclik-data")
	app.Get("/clear-and-pull-autoclik", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAll)
	app.Get("/get-alldata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByLimitandOffsetHandler)
	app.Get("/get-all", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllHandler)
	app.Get("/get-category", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCategory)
	app.Get("/get-count", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByCountHandler)
	app.Get("/get-no/:No", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByNoHandler)
	app.Get("/get-gen-posting-group/:gen_prod_posting_group", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetBygen_prod_posting_groupHandler)

	// app.Get("/clear-and-suck", ControllerAutoclik_data.ClearAndSuck)

	// -------------------------------------------------------------------------------------------------------
	app.Get("/get-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllBindata)
	// app.Get("/suck-information", middleware.AuthorizationRequired(), ControllerAutoclik_data.Suck_information)
	app.Get("/get-all-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllAutoclik_dataHandler)
	app.Get("/get-count-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCountAllItemBin)
	app.Get("/clear-and-suck-item-bin", middleware.AuthorizationRequired(), ControllerAutoclik_data.Clearandsuckbinitem)
}

func SetupAutoclik_bin_dataRoutes(router fiber.Router) {
	app := router.Group("autoclik-bin-data")

	app.Get("/get-item-no", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByItem_No_Autoclik_BinHandler)
	app.Get("/get-product-group", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetByProductGroupHandler)
	app.Get("/get-location-code", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllGetLocation_Code)

	app.Get("/get-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllBindata)
	app.Get("/get-count-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCountAutoclik_data)
	app.Get("/suck-information", middleware.AuthorizationRequired(), ControllerAutoclik_data.Suck_information)
	app.Get("/get-all-autoclik-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetAllAutoclik_dataHandler)
	app.Get("/get-count-bindata", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetCountAllItemBin)

	//ดึงข้อมูล autoclik-bin-data
	app.Get("/clear-and-pull-item-bin", middleware.AuthorizationRequired(), ControllerAutoclik_data.Clearandsuckbinitem)

	app.Get("/filter-item-bin", middleware.AuthorizationRequired(), ControllerAutoclik_data.GetChoosebytypeAutoclikbin)
}
