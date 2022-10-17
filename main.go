package main

import (
	"fmt"
	"golang-hacktiv8-assignment-2/controller"
	"golang-hacktiv8-assignment-2/database"
	"golang-hacktiv8-assignment-2/router"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("Error starting database", err)
		return
	}

	ctrl := controller.New(db)

	err = router.StartServer(ctrl)
	if err != nil {
		fmt.Println("Error start server", err)
		return
	}
}
