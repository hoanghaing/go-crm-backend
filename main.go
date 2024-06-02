package main

import (
	"fmt"
	"log"
	"net/http"

	"go-crm/controllers"
	"go-crm/models"
	"go-crm/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	dsn := "host=localhost user=postgres password=12345678aA@ dbname=go_crm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	controllers.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	controllers.DB.AutoMigrate(&models.Customer{})

	r := routers.SetupRouter()
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
