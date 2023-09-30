package main

import (
	"net/http"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/routes"
	"github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
)

func main() {

	db.Connect()
	db.DB.AutoMigrate(models.Users{})
	db.DB.AutoMigrate(models.Permisions{})
	db.DB.AutoMigrate(models.Commentary{})

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/auth/login", routes.Login).Methods("POST")
	r.HandleFunc("/auth/signup", routes.SignUp).Methods("POST")

	r.HandleFunc("/commentaries", routes.GetCommentaries).Methods("GET")
	r.HandleFunc("/commentaries/{id}", utils.ValidateJWT(routes.GetCommentary)).Methods("GET")
	r.HandleFunc("/commentaries", routes.CreateCommentary).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
