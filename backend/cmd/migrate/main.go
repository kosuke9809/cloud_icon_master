package migrate

import (
	"fmt"

	database "github.com/kosuke9809/cloud_icon_master/internal/infra/database/config"
	migration "github.com/kosuke9809/cloud_icon_master/internal/infra/database/migration"
)

func Start() {
	fmt.Println("Start db connection.")
	dbConn := database.NewDB()
	defer database.CloseDB(dbConn)
	fmt.Println("Connected...")

	fmt.Println("Start db migration.")
	migration.Migrate(dbConn)
	fmt.Println("Complete insert seed data.")
}
