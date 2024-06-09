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
	fmt.Println("GetCustomers")
	var customers []models.Customer
	if err := DB.Find(&customers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customersJSON, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(customersJSON)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer models.Customer
	DB.First(&customer, params["id"])
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(customerJSON)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Create(&customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(customerJSON)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCustomer")
	params := mux.Vars(r)
	var customer models.Customer
	DB.First(&customer, params["id"])
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(&customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(customerJSON)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteCustomer")
	params := mux.Vars(r)
	var customer models.Customer
	if err := DB.Delete(&customer, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Customer Deleted"}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
