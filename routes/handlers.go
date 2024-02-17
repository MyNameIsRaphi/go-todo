package routes

import (
	"html/template"
	"os"
	"net/http"
	"todo/jwt"

	"github.com/sirupsen/logrus"
)

func getRoot(res http.ResponseWriter, req *http.Request) {
	//TODO render html file

	// Check if request is already logged in, if yes send them the corresponding html document
	// otherwise send login form
	token, err := req.Cookie("jwt")
	if  err == http.ErrNoCookie || !jwt.Validate(token.String()) {
	
		tmpl,err := template.ParseFiles("views/login.ejs")	
		if err != nil {
			http.Error(res,"Not found" , http.StatusInternalServerError)
			return
		}

		data := map[string]bool{"unauth":false}
		err = tmpl.Execute(res, data)
		if err != nil {
			logrus.WithError(err).Errorf("Failed to send data to %s", req.Host)
		}
	} else {
			
		http.Redirect(res, req,"/", http.StatusMovedPermanently)
		
	}
}

func getCSS(res http.ResponseWriter, req *http.Request){
	content, err := os.ReadFile("views/style.css")
	if err != nil {
		http.Error(res, "File not found", 404)
	}
	res.Write(content)
}
