package routes

import (
	"net/http"
	"os"

	request "github.com/al3xdiaz/go-server/utils"
)

func Version(w http.ResponseWriter, r *http.Request) {
	// ...
	var version = os.Getenv("API_VERSION")
	request.Ok(w, version)
}
