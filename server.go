package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

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
	db.DB.AutoMigrate(models.Site{})
	db.DB.AutoMigrate(models.Commentary{})

	// db.DB.Exec("Delete from users")
	// db.DB.Exec("Delete from commentaries")
	// db.DB.Exec("Delete from sites")

	r := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("/static"))

	r.HandleFunc("/auth/login", routes.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/auth/signup", routes.SignUp).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/auth/userdata", utils.RequireAuth(routes.UserData)).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/commentaries", routes.GetCommentaries).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/commentaries/{id}", utils.RequirePermision(routes.GetCommentary)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/commentaries", utils.RequireAuth(routes.CreateCommentary)).Methods(http.MethodPost, http.MethodOptions)
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.Handle("/", r)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://*.chaoticteam.com", "https://chaoticteam.com"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(r)
	http.ListenAndServe(":8000", handler)
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
