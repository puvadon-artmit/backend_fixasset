package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/puvadon-artmit/gofiber-template/seeder/seeds"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// p := config.GetEnvConfig("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)

	// if err != nil {
	// 	log.Println("Wrong port!")
	// }

	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Silent,          // Log level
			Colorful:      false,                  // Disable color
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	if err = DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Permission{},
		&model.PermissionComponent{},
		&model.PermissionGroup{},
		&model.Category{},
		&model.Branch{},
		&model.Group{},
		&model.Group_Story{},
		&model.Status{},
		&model.Type_things{},
		&model.Manufacturer{},
		&model.Item_model{},
		&model.Story{},
		&model.Ground{},
		&model.Ground_Story{},
		&model.Responsible{},
		&model.Responsible_Story{},
		&model.Assets{},
		&model.Assets_Story{},
		&model.Typeplan{},
		&model.Typeplan_story{},
		&model.Asset_count{},
		&model.Asset_check{},
		&model.Location{},
		&model.Location_story{},
		&model.Signature{},
		&model.Main_branch{},
		&model.Branch_Story{},
		&model.Main_Branch_Story{},
		&model.Item_Autoclik{},
		&model.Maliwan_data{},
		&model.Round_Count{},
		&model.Round_Count_Story{},
		&model.Main_Category{},
		&model.Branch_Autoclik{},
		&model.Count_Autoclik{},
		&model.Asset_count_Main_Category{},
		&model.Asset_count_Category{},
		&model.Category_Story{},
		&model.Counting_rights{},
		&model.Manufacturer_Story{},
		&model.Asset_count_story{},
		&model.Scan_story{},
		&model.Main_Category_story{},
		&model.Status_story{},
		&model.Type_things_story{},
		&model.Autoclik_count{},
		&model.Autoclik_count_Story{},
		&model.Autoclik_Round_Count{},
		&model.Autoclik_Round_Count_Story{},
		&model.Autoclik_check{},
		&model.Autoclik_Counting_Rights{},
		&model.Autoclik_AllPhoto{},
		&model.Autoclik_Count_Product_Group{},
		&model.Item_Autoclik_Bin_Code{},
		&model.Autoclik_check_Story{},
		&model.Signature_Autoclik{},
		&model.Autoclik_Update_Story{},
		&model.Autoclik_Fixed_Asset{},
		&model.Maliwan_Update_Story{},
		&model.User_story{},
		&model.Maliwan_count{},
		&model.Maliwan_counts_story{},
		&model.Maliwan_Round_Count{},
		&model.Maliwan_Counts_Item_Category_Code{},
		&model.Maliwan_Round_Count_Story{},
		&model.Maliwan_Counting_Rights{},
		&model.Maliwan_check{},
		&model.PhotoMaliwanCheck{},
		&model.Maliwan_Photos_check{},
		&model.Signature_Maliwan{},
		&model.Maliwan_check_Story{},
		&model.Autoclik_Counting_Trigger{},
		&model.GenProductPostingGroups{},
		&model.Autoclik_Count_Store{},
		&model.Maliwan_Counting_Trigger{},
		&model.Maliwan_Count_Store{},
		&model.Item_bin_maliwan{},
		&model.Autoclik_Fixed_Asset_Count{},
		&model.Autoclik_Fixed_Asset_Store{},
		&model.Autoclik_Fixed_Asset_Counting_Rights{},
		&model.Autoclik_Fixed_Asset_Round_Count{},
		&model.Autoclik_Fixed_Asset_Round_Count_Story{},
		&model.Autoclik_Fixed_Asset_Check{},
		&model.Autoclik_Fixed_Asset_Photos_check{},
		&model.Autoclik_Fixed_Asset_count_Story{},
		&model.Autoclik_Fixed_Asset_check_Story{},
		&model.Signature_Autoclik_Fixed_Asset{},
		&model.Assets_Count_Store{},
		&model.Maliwan_Fixed_Asset_Count{},
		&model.Maliwan_Fixed_Asset_count_Story{},
		&model.Maliwan_Fixed_Asset_Round_Count{},
		&model.Maliwan_Fixed_Asset_Round_Count_Story{},
		&model.Maliwan_Fixed_Asset_Counting_Rights{},
		&model.Maliwan_Fixed_Asset_Photos_check{},
		&model.Maliwan_Fixed_Asset_Check{},
		&model.Signature_Maliwan_Fixed_Asset{},
		&model.Maliwan_Item_Category_Code{},
		&model.Maliwan_Fixed_Asset_Store{},
		&model.Maliwan_Fixed_Asset{},
		&model.Maliwan_Fixed_Asset_check_Story{},
		&model.Maliwan_Update_Fixed_Asset_Story{},
		&model.Request_Update_Data{},
	); err == nil && DB.Migrator().HasTable(&model.Role{}) {
		if err := DB.First(&model.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			for _, seed := range seeds.All() {
				if err := seed.Run(DB); err != nil {
					log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
				}
			}
		}
	}
	fmt.Println("Database Migrated")
}
