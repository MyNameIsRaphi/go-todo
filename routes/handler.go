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
		GetRequests(&res, req)
	case "POST":
		PostRequest(&res, req)
	}
	
}

func PostRequest(res *http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	var err error
	switch path {
	case "/":
		err = postRoot(res, req)		

	case "/register":
		// TODO add register handler
		
	default:
		// TODO add default handler for post
	}

	if err != nil {
		logrus.WithError(err).Error("Error occured during post request")
	}
}

func GetRequests(res *http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	var err error
	switch path {
	case "/":
		// TODO Add handler function
		err = getRoot(res, req)
	case "/home":
		// TODO add home handler
	case "/register":
		err = getRegister(res, req)
	case "/style.css":
		err = getCSS(res, req)		
	case "/registerCSS":
		err = sendRegisterCSS(res, req)
	default:
		// TODO add default response function
		http.Error(*res, "Not found", 404)
	}
	if err != nil {
		logrus.WithError(err).Error("Error occurred in get request")
	}
}
func LogRequest(request http.Request) {
	var method string = request.Method
	var path string = request.URL.Path
	var remoteAddr string = request.RemoteAddr
	logrus.Infof("request from %v \tmethod: %v\tpath: %v", remoteAddr, method, path)
	
}


type registerData struct {
	EmailUsedError bool
	NoEmailError bool 
	NoMatchError bool 
	NoPassWordError bool
}

type loginData struct {
	Unauth bool
}





