package routers

import (
	"go-crm/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// r.Use(commonMiddleware)
	// Apply common middleware to all routes except "/"
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				commonMiddleware(next).ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	})

	r.HandleFunc("/", controllers.GetHome).Methods("GET")
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
