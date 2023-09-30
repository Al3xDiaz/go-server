package routes

import (
	"net/http"

	"github.com/al3xdiaz/go-server/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != utils.API_KEY {
			return
		} else {
			token, err := utils.CreateJWT()
			if err != nil {
				return
			}
			utils.Ok(w, token)
		}
	}
}
