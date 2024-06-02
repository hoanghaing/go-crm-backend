package routers

import (
	"go-crm/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	r.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return r
}
