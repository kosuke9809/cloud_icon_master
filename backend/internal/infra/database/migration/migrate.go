package database

import (
	"fmt"
	"log"
	"os"

	"github.com/kosuke9809/cloud_icon_master/internal/domain/model"
	"gorm.io/gorm"
)

func Migrate(dbConn *gorm.DB) {
	fmt.Println("Start DB Migration")
	dbConn.AutoMigrate(&model.User{}, &model.Icon{})
	content, err := os.ReadFile("./internal/infra/database/seed/init_icons.sql")
	if err != nil {
		log.Println(err)
	}

	res := dbConn.Exec(string(content))
	if res.Error != nil {
		log.Println(res.Error)
	}

}
