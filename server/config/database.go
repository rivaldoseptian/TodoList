package config

import (
	"fmt"
	"server/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV.MYSQL_USER, ENV.MYSQL_PASSWORD, ENV.MYSQL_HOST, ENV.MYSQL_PORT, ENV.MYSQL_DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DataBase Error")
	}

	db.AutoMigrate(&models.Activities{}, &models.Todo{})

	DB = db
	log.Println("Sucess Connect To DataBase")

}
