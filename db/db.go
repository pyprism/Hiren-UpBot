package db

import (
	"errors"
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

	db.AutoMigrate(&models.User{}, &models.URL{}, &models.Mailgun{})
	db.Model(&models.URL{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}

func GetDB() *gorm.DB {
	return db
}

func (h *Hiren) UserCount() int64 {
	var count int64
	db.Model(models.User{}).Count(&count)
	return count
}

func (h *Hiren) UserCreate(username, hash string, admin bool) bool {
	user := models.User{UserName: username, Password: hash, Admin: admin}
	db.Create(&user)
	return db.NewRecord(user)
}

func (h *Hiren) FindUserByUsername(username string) (models.User, error) {
	var user models.User
	var error error
	db.Where(&models.User{UserName: username}).First(&user)
	if user.ID == 0 {
		error = errors.New("not found")
		return user, error
	} else {
		error = nil
		return user, error
	}

}

func (h *Hiren) UrlCreate(name, url, username string, poll, alert int64) bool {
	var user models.User
	db.Where(&models.User{UserName: username}).First(&user)
	urlObj := models.URL{
		Name:            name,
		Url:             url,
		UserID:          user.ID,
		PollingInterval: poll,
		AlertThreshold:  alert,
	}
	db.Create(&urlObj)
	return db.NewRecord(urlObj)

}
