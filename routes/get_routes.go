package routes

import (
	"net/http"
	"os"
)

func getMainAssets(res *http.ResponseWriter, req *http.Request) error {
	path := req.URL.Path
	b, err := os.ReadFile("./views/dist" + path)
	if path[len(path)-3:] == ".js" {
		(*res).Header().Set("Content-Type", "application/javascript")
	} else if path[len(path)-4:] == ".css" {
		(*res).Header().Set("Content-Type", "text/css")
	}
	if err != nil {
		http.Error(*res, "Not found", 404)
		return err
	}
	_, err = (*res).Write(b)
	return err
}

func getRoot(res *http.ResponseWriter, req *http.Request) error {
	b, err := os.ReadFile("./views/dist/index.html")
	if err != nil {
		http.Error(*res, "Not found", 404)
		return err
	}
	_, err = (*res).Write(b)
	return err
}
