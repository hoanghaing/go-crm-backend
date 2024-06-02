package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-crm/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("GetCustomers")
	var customers []models.Customer
	if err := DB.Find(&customers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println("GetCustomers: ", customers)
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCustomer")
	params := mux.Vars(r)
	var customer models.Customer
	DB.First(&customer, params["id"])
	json.NewEncoder(w).Encode(&customer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateCustomer")
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	DB.Create(&customer)
	json.NewEncoder(w).Encode(&customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCustomer")
	params := mux.Vars(r)
	var customer models.Customer
	DB.First(&customer, params["id"])
	json.NewDecoder(r.Body).Decode(&customer)
	DB.Save(&customer)
	json.NewEncoder(w).Encode(&customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteCustomer")
	params := mux.Vars(r)
	var customer models.Customer
	DB.Delete(&customer, params["id"])
	json.NewEncoder(w).Encode("Customer Deleted")
}
