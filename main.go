package main

import (
	"fmt"
	"log"
	"net/http"
	"soul_api/middleware"
	"soul_api/routes/customers"
	"soul_api/routes/team"
	"soul_api/routes/users"
	"soul_api/routes/partners"
    "soul_api/routes/customer_with_partner"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)

	r.HandleFunc("/api/users", users.Create).Methods("POST")
	r.HandleFunc("/api/users/show", middleware.ValidateTokenMiddleware(users.List)).Methods("GET")
	r.HandleFunc("/api/users/update", middleware.ValidateTokenMiddleware(users.Update)).Methods("PUT")
	r.HandleFunc("/api/users/delete", middleware.ValidateTokenMiddleware(users.Delete)).Methods("DELETE")
	r.HandleFunc("/api/users/login", users.Login).Methods("POST")

	r.HandleFunc("/team/login", team.Login).Methods("POST")
	// r.HandleFunc("/team/list", middleware.ValidateTokenMiddleware(team.List)).Methods("GET")
	r.HandleFunc("/team/list", team.List).Methods("GET")
	r.HandleFunc("/team/add-member", team.Create).Methods("POST")
	r.HandleFunc("/team/update-member", middleware.ValidateTokenMiddleware(team.Update)).Methods("PUT")
	r.HandleFunc("/team/update-member/password", middleware.ValidateTokenMiddleware(team.UpdatePassword)).Methods("PUT")
	r.HandleFunc("/team/update-team-member", middleware.ValidateTokenMiddleware(team.UpdateMember)).Methods("PUT")
	r.HandleFunc("/team/view-member", middleware.ValidateTokenMiddleware(team.View)).Methods("GET")
	// r.HandleFunc(`/team/edit-team-member/{:id}`, middleware.ValidateTokenMiddleware(team.View)).Methods("POST")
	r.HandleFunc("/team/update-status", middleware.ValidateTokenMiddleware(team.UpdateStatus)).Methods("POST")
	r.HandleFunc("/team/logout", middleware.ValidateTokenMiddleware(team.Logout)).Methods("GET")
	r.HandleFunc("/update/profilePic", team.UploadImage).Methods("POST")

	r.HandleFunc("/team/has-role", team.HasRole).Methods("GET")

	r.HandleFunc("/customers/registration", customers.Customers).Methods("POST")
	r.HandleFunc("/customers/booking", customers.Booking).Methods("POST")
	r.HandleFunc("/customers/transaction", customers.Transaction).Methods("POST")
	r.HandleFunc("/customers/view", middleware.ValidateTokenMiddleware(customers.View)).Methods("GET")
	r.HandleFunc("/customers/list", middleware.ValidateTokenMiddleware(customers.List)).Methods("GET")
	r.HandleFunc("/customers/booking/view", middleware.ValidateTokenMiddleware(customers.ViewBooking)).Methods("GET")
	r.HandleFunc("/customers/transaction/view", middleware.ValidateTokenMiddleware(customers.ViewTransaction)).Methods("GET")
	//r.HandleFunc("/customers/list", middleware.ValidateTokenMiddleware(customers.List)).Methods("GET")
	r.HandleFunc("/customers/update", middleware.ValidateTokenMiddleware(customers.Update)).Methods("PUT")
	r.HandleFunc("/customers/transaction/update", middleware.ValidateTokenMiddleware(customers.UpdateTransaction)).Methods("PUT")

	r.HandleFunc("/partner/register", partner.Create).Methods("POST")
	r.HandleFunc("/partner/update-partner", middleware.ValidateTokenMiddleware(partner.Update)).Methods("PUT")

	r.HandleFunc("/customer-with-partner/add", customer_with_partner.Create).Methods("POST")
    r.HandleFunc("/customer-with-partner/update", middleware.ValidateTokenMiddleware(customer_with_partner.Update)).Methods("PUT")

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}
