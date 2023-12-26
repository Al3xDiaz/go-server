package main

import (
	"log"
	"net/http"
	"os"

	"github.com/al3xdiaz/go-server/db"
	"github.com/al3xdiaz/go-server/models"
	"github.com/al3xdiaz/go-server/routes"
	"github.com/al3xdiaz/go-server/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func newREST() *mux.Router {

	db.Connect()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Permision{})
	db.DB.AutoMigrate(models.Site{})
	db.DB.AutoMigrate(models.Commentary{})

	r := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("./static"))

	r.HandleFunc("/version", routes.Version).Methods(http.MethodGet)
	r.HandleFunc("/auth/login", routes.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/signup", routes.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/auth/userdata", utils.RequireAuth(routes.UserData)).Methods(http.MethodGet)

	r.HandleFunc("/commentaries", routes.GetCommentaries).Methods(http.MethodGet)
	r.HandleFunc("/commentaries/{id}", utils.RequireAuth(routes.GetCommentary)).Methods(http.MethodGet)
	r.HandleFunc("/commentaries/{id}", utils.RequireAuth(routes.DeleteCommentary)).Methods(http.MethodDelete)
	r.HandleFunc("/commentaries", utils.RequireAuth(routes.CreateCommentary)).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.Handle("/", r)
	return r
}
func RunServer() {
	router := newREST()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(credentials, methods, origins, headers)(router)))
}

func main() {
	command := os.Args[1]
	switch command {
	case "runserver":
		RunServer()
	case "permisions":
		SelectUser()
	case "cleandata":
		db.Connect()
		db.DB.Exec("Delete from users")
		db.DB.Exec("Delete from commentaries")
		db.DB.Exec("Delete from sites")
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
