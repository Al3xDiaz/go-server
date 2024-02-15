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
	db.DB.AutoMigrate(
		models.User{},
		models.Profile{},
		models.Telephone{},
		models.Course{},
		models.Permision{},
		models.Site{},
		models.Commentary{},
	)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/version", routes.Version).Methods(http.MethodGet)
	r.HandleFunc("/vcard/{username}", routes.VCard).Methods(http.MethodGet)

	r.HandleFunc("/auth/login", routes.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/signup", routes.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/auth/userdata", utils.RequireAuth(routes.UserData)).Methods(http.MethodGet)
	r.HandleFunc("/profile", utils.RequireAuth(routes.UpdateProfile)).Methods(http.MethodPatch)
	r.HandleFunc("/profile", utils.RequireAuth(routes.GetProfile)).Methods(http.MethodGet)
	r.HandleFunc("/telephone", utils.RequireAuth(routes.PostTelephone)).Methods(http.MethodPost)

	r.HandleFunc("/users", routes.GetUsers).Methods(http.MethodGet)

	r.HandleFunc("/commentaries", routes.GetCommentaries).Methods(http.MethodGet)
	r.HandleFunc("/commentaries/{id}", utils.RequireAuth(routes.GetCommentary)).Methods(http.MethodGet)
	r.HandleFunc("/commentaries/{id}", utils.RequireAuth(routes.DeleteCommentary)).Methods(http.MethodDelete)
	r.HandleFunc("/commentaries", utils.RequireAuth(routes.CreateCommentary)).Methods(http.MethodPost)

	http.Handle("/", r)
	return r
}
func RunServer() {
	router := newREST()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{
		// "GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	})
	// origins := handlers.AllowedOrigins([]string{"https://alex.chaoticteam.com", "http://localhost:3000"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Authorization", "content-type", "X-Requested-With"})
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
		log.Output(0, "cleaning data...")
		db.Connect()
		db.DB.Exec("Delete from users")
		db.DB.Exec("Delete from commentaries")
		db.DB.Exec("Delete from sites")
		log.Output(0, "data cleaned")
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
