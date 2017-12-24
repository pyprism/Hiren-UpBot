package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/spf13/viper"
)

type Hiren struct {
	Server string
}

var db *gorm.DB
var err error

func (h *Hiren) Connect() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	h.Server = "host=localhost" + " user=" + viper.GetString("db_user") + " dbname=" + viper.GetString("db_name") + " sslmode=disable" + " password=" + viper.GetString("db_password")
	db, err = gorm.Open("postgres", h.Server)
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	//db.DropTables(&models.User{})
	//defer db.Close()

	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}

func (h *Hiren) UserCount() int64 {
	var count int64
	db.Model(models.User{}).Count(&count)
	return count
}
