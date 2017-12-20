package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/spf13/viper"
)

var db *xorm.Engine
var err error

func Init() {
	db, err = xorm.NewEngine("postgres", "host=localhost"+" user="+viper.GetString("db_user")+
		" dbname="+viper.GetString("db_name")+" sslmode=disable"+" password="+viper.GetString("db_password"))
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.ShowSQL(true)
	//db.DropTables(&models.User{})
	db.CreateTables(&models.User{})
	db.CreateUniques(new(models.User))
}

func GetDB() *xorm.Engine {
	return db
}

//func CloseDB() {
//	db.Close()
//}
