package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Request struct {
	body          []byte
	token         string
	auth          bool
	wait_time_sec int
	startTime     time.Time
	res           *http.ResponseWriter
	req           *http.Request
	status        int
}

func Middleware(res http.ResponseWriter, req *http.Request) {
	var method string = req.Method
	go LogRequest(*req) // log the request
	switch method {
	case "GET":
		GetRequests(&res, req)
	case "POST":
		PostRequest(&res, req)
	case "PUT":
	}

}
func PutRequest(res *http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	var err error
	var put = Put{Request{res: res, req: req, wait_time_sec: 2}}
	err = put.Init()
	if err != nil {
		logrus.WithError(err).Error("failed to init put request")
		http.Error(*res, "bad request", 400)
		return
	}
	switch path {
	case "/updatetodo":
		err = put.UpdateToDo(res, req)

	default:
		http.Error(*res, "bad request", 400)
		// TODO add default handler for post
	}

	if err != nil {
		logrus.WithError(err).Error("Error occured during post request")
		http.Error(*res, "internal server error", http.StatusInternalServerError)
	}
}
func PostRequest(resp *http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	var err error
	var post = Post{Request{res: resp, req: req, wait_time_sec: 2}}
	if path != "/login" && path != "/register" {
		err = post.Init()
		if err != nil {
			logrus.WithError(err).Error("failed to init post request")
			http.Error(*resp, "bad request", 400)
			return
		}
	}

	switch path {
	case "/login":
		err = post.Login()

	case "/register":
		// TODO add register handler
		err = post.Register()
	case "/addtodo":
		err = post.AddToDo()
	default:
		// TODO add default handler for post
	}

	if err != nil {
		logrus.WithError(err).Error("Error occured during post request")
		http.Error(*resp, "internal server error", http.StatusInternalServerError)
	}
}

func GetRequests(res *http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	var err error
	switch path {
	case "/":
		// TODO send html page
		err = getRoot(res, req)
	default:
		// TODO add default response function
		err = getMainAssets(res, req)
	}
	if err != nil {
		logrus.WithError(err).Error("Error occurred in get request")
		http.Error(*res, "internal server error", http.StatusInternalServerError)
	}
}
func LogRequest(request http.Request) {
	var method string = request.Method
	var path string = request.URL.Path
	var remoteAddr string = request.RemoteAddr
	logrus.Infof("request from %v \tmethod: %v\tpath: %v", remoteAddr, method, path)

}

func sendJSON[T any](res *http.ResponseWriter, data T) error {
	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = (*res).Write(b)
	return err
}
func wait(start *time.Time, wait_time_sec int) {
	var end time.Time = time.Now()

	var duration int64 = end.Sub(*start).Nanoseconds()

	var waitime int64 = (int64(wait_time_sec) * 1_000_000_000) - duration
	time.Sleep(time.Duration(waitime))

}

type Default struct {
	Token string
}

type defaultResponse struct {
	Success bool
	Note    string
}
