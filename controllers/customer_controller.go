package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"go-crm/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetHome")
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>API Overview</title>
		</head>
		<body>
			<h1>Welcome to the API</h1>
			<p>This API provides the following endpoints:</p>
			<ul>
				<li><strong>GET /customers</strong>: Get all customers</li>
				<li><strong>GET /customers/{id}</strong>: Get a specific customer by ID</li>
				<li><strong>POST /customers</strong>: Create a new customer</li>
				<li><strong>PUT /customers/{id}</strong>: Update an existing customer</li>
				<li><strong>DELETE /customers/{id}</strong>: Delete a customer</li>
			</ul>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and write the output to the response writer
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
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
