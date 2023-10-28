package routes

import (
	"net/http"
)

func Middleware(res http.ResponseWriter, req *http.Request) {
	var method string = req.Method

	switch method {
	case "GET":
		GetRequests(res, req)
	case "POST":

	}
}

func PostRequest(res http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path

	switch path {
	case "/login":
		// TODO Add login handler function
	case "/register":
		// TODO add register handler
	default:
		// TODO add default handler for post
	}
}

func GetRequests(res http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path

	switch path {
	case "/":
		// TODO Add handler function
	default:
		// TODO add default response function
	}
}
