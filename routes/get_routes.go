package routes

import (
	"html/template"
	"os"
	"net/http"
	"todo/jwt"

)

func getRoot(res *http.ResponseWriter, req *http.Request) error{
	//TODO render html file

	// Check if request is already logged in, if yes send them the corresponding html document
	// otherwise send login form
	token, err := req.Cookie("jwt")
	if  err == http.ErrNoCookie || !jwt.Validate(token.String()) {
	
		tmpl,err := template.ParseFiles("views/login.ejs")	
		if err != nil {
			http.Error(*res,"Not found" , http.StatusInternalServerError)
			return err
		}

		data := loginData{false} 
		err = tmpl.Execute(*res, data)
		if err != nil {
			return err
		}
	} else {
		http.Redirect(*res, req,"/", http.StatusMovedPermanently)
		
	}
	return nil
}

func getRegister(res *http.ResponseWriter, req *http.Request) error {
	templ, err := template.ParseFiles("views/register.ejs")
	if err != nil {
		http.Error(*res, "Internal Server error", http.StatusInternalServerError) 
		
		return err
	}
	data := registerData{false, false, false, false } 

	return templ.Execute(*res, data)

}

func getCSS(res *http.ResponseWriter, req *http.Request) error{
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
		http.Error(*res,"404 not found", 404)
		return err
	}
	(*res).Header().Set("Content-Type", "text/css")
	_, err = (*res).Write(content)
	return err
}

