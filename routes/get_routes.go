package routes

import (
	"html/template"
	"net/http"
	"os"
	"todo/jwt"
)

func getRoot(res *http.ResponseWriter, req *http.Request) error {
	//TODO render html file

	// Check if request is already logged in, if yes send them the corresponding html document
	// otherwise send login form
	token, err := req.Cookie("JWT")
	if err == http.ErrNoCookie || !jwt.Validate(token.String()) {

		tmpl, err := template.ParseFiles("views/login.ejs")
		if err != nil {
			http.Error(*res, "Not found", http.StatusInternalServerError)
			return err
		}

		data := loginData{false}
		err = tmpl.Execute(*res, data)
		if err != nil {
			return err
		}
	} else {
		// send home page
		http.Redirect(*res, req, "/home", http.StatusMovedPermanently)
	}
	return nil
}

func getRegister(res *http.ResponseWriter, req *http.Request) error {
	templ, err := template.ParseFiles("views/register.ejs")
	if err != nil {
		http.Error(*res, "Internal Server error", http.StatusInternalServerError)

		return err
	}
	data := registerData{false, false, false, false}

	return templ.Execute(*res, data)

}

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

func getCSS(res *http.ResponseWriter, req *http.Request) error {
	content, err := os.ReadFile("views/style.css")
	if err != nil {
		http.Error(*res, "File not found", 404)
		return err
	}
	(*res).Header()["Content-Type"] = []string{"text/css"}
	(*res).WriteHeader(http.StatusOK)
	_, err = (*res).Write(content)
	return err
}

func sendRegisterCSS(res *http.ResponseWriter, req *http.Request) error {
	content, err := os.ReadFile("views/registerCSS.css")
	if err != nil {
		http.Error(*res, "404 not found", 404)
		return err
	}
	(*res).Header().Set("Content-Type", "text/css")
	_, err = (*res).Write(content)
	return err
}

func getHome(res *http.ResponseWriter, req *http.Request) error {
	token, err := req.Cookie("JWT")
	if err != nil || !jwt.Validate(token.String()) {
		http.Redirect(*res, req, "/", http.StatusMovedPermanently)
		return err
	} else {
		b, err := os.ReadFile("./views/home/public/dist/index.html")
		if err != nil {
			return err
		}
		_, err = (*res).Write(b)
		return err
	}

}
