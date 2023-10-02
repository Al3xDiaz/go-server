package main

import (
	"log"
	"net/http"
	"os"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/routes"
	"github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RunServer() {

	db.Connect()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Permision{})

	db.DB.AutoMigrate(models.Commentary{})

	r := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("/static"))

	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	r.HandleFunc("/auth/login", routes.Login).Methods("POST")
	r.HandleFunc("/auth/signup", routes.SignUp).Methods("POST")
	r.HandleFunc("/auth/userdata", utils.RequireAuth(routes.UserData)).Methods("GET")

	r.HandleFunc("/commentaries", routes.GetCommentaries).Methods("GET")
	r.HandleFunc("/commentaries/{id}", utils.RequirePermision(routes.GetCommentary)).Methods("GET")
	r.HandleFunc("/commentaries", utils.RequireAuth(routes.CreateCommentary)).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func main() {
	command := os.Args[1]
	switch command {
	case "runserver":
		RunServer()
	case "permisions":
		SelectUser()
	default:
		log.Output(0, "command not found")
	}

}
func SelectUser() {
	db.Connect()
	username := os.Args[2]

	user := models.User{}
	db.DB.First(&user, "user_name = ?", username)
	AddPermisions(db.DB, &user)
}
func AddPermisions(db *gorm.DB, user *models.User) {
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("insert path in regex:")
	// path, _ := reader.ReadString('\n')
	// fmt.Println("insert methods in regex:")
	// methods, _ := reader.ReadString('\n')

	db.Model(&user).Association("Permisions").Append(&models.Permision{
		// Path:    path,
		// Methods: methods,
		Path:    "/.*",
		Methods: ".*",
	})

	// fmt.Println("cluld you add more permisions?:(y/N)")
	// resp, _ := reader.ReadString('\n')
	// if resp == "y" {
	// 	AddPermisions(db, user)
	// }
}
