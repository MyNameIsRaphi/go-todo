package routes

import (
	"net/http"
	"os"
)





func getMainAssets(res *http.ResponseWriter, req *http.Request) error {
	path := req.URL.Path
	b, err := os.ReadFile("./views/home/public/dist" + path)
	if err != nil {
		http.Error(*res, "Not found", 404)
		return err
	}
	_, err = (*res).Write(b)
	return err
}
