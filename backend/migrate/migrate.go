package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kosuke9809/aws_icons_quiz_backend/database"
	"github.com/kosuke9809/aws_icons_quiz_backend/model"
)

func main() {
	fmt.Println("Start DB connect")
	dbConn := database.NewDB()
	defer fmt.Println("Successfully connected to database")
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Icon{})

	content, err := os.ReadFile("./migrate/sql/init_icons.sql")
	if err != nil {
		log.Fatal(err)
	}

	res := dbConn.Exec(string(content))
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	fmt.Println("complete insert init data")
}
