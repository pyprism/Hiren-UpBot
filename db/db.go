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

func (h *Hiren) UrlList(username string) []models.URL {
	var user models.User
	var urls []models.URL
	db.Where(&models.User{UserName: username}).First(&user)
	db.Where("user_id = ?", user.ID).Find(&urls)
	return urls
}

func (h *Hiren) FindHostById(username, id string) (models.URL, error) {
	var user models.User
	var url models.URL
	db.Where(&models.User{UserName: username}).First(&user)
	db.Where("user_id = ? and id =?", user.ID, id).First(&url)
	if user.ID == 0 || url.ID == 0 {
		err := errors.New("not found")
		return url, err
	} else {
		var err error = nil
		return url, err
	}
}

func (h *Hiren) DeleteUrlByID(username, id string) error {
	var user models.User
	db.Where(&models.User{UserName: username}).First(&user) // just for security !

	if user.ID == 0 {
		err := errors.New("not found")
		return err
	} else {
		var err error = nil
		db.Delete(&models.URL{}, "id = ?", id)
		return err
	}
}

//func (h *Hiren) UpdateUrlById(username, id, name, url string)
