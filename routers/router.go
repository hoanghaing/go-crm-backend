package routers

import (
	"go-crm/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(commonMiddleware)
	r.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	r.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
