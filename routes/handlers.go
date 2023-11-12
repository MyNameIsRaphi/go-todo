package routes

import (
	"net/http"
)

func getRoot(res http.ResponseWriter, req *http.Request) {
	//TODO render html file

	// Check if request is already logged in, if yes send them the corresponding html document
	// otherwise send login form
}
