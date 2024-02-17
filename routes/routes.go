package routes

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Middleware(res http.ResponseWriter, req *http.Request) {
	var method string = req.Method
	go LogRequest(*req) // log the request
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
		getRoot(res, req)
	case "/home":
		// TODO add home handler
	case "/style.css":
		getCSS(res, req)		
	default:
		// TODO add default response function
	}
}
func LogRequest(request http.Request) {
	var method string = request.Method
	var path string = request.URL.Path
	var remoteAddr string = request.RemoteAddr
	logrus.Infof("request from %v \tmethod: %v\tpath: %v", remoteAddr, method, path)
}
