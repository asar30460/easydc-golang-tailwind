package main

import (
	"fmt"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Initilize db in main. ", err)
	}
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	r := router.InitRouter(userHandler)
	router.Start(r, ":8080")
}
