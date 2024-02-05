package main

import (
	"github.com/kosuke9809/cloud_icon_master/cmd/api"
	"github.com/kosuke9809/cloud_icon_master/cmd/migrate"
)

func main() {
	migrate.Start()
	api.Start()
}
