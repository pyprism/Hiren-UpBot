package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/spf13/viper"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("postgres", "host=localhost"+" user="+viper.GetString("db_user")+
		" dbname="+viper.GetString("db_name")+" sslmode=disable"+" password="+viper.GetString("db_password"))
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

