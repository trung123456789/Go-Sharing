package main

import (
	"urls"
	"utils"
)

func main() {
	// Auto generate db
	utils.AutoMigrateDb()
	urls.GeneralURL()
}
