package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"todo/database"
	"todo/encrypt"
	"todo/jwt"
	"todo/types"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func postRoot(res *http.ResponseWriter, req *http.Request) error {
	
	var start = time.Now()
	const WAIT_TIME_SEC = 2
	

	var err error
	parse_err := (*req).ParseForm()	

	if parse_err != nil {
		wait(&start, WAIT_TIME_SEC)
		err = sendInvalidtempl(res)	
		if err != nil {
			http.Error(*res, "Internal server error", http.StatusInternalServerError)
		}
		return parse_err

	} else {
		email := req.FormValue("email")	
		password := req.FormValue("password")
		if email == "" || password == "" {
			wait(&start, WAIT_TIME_SEC)
			err = sendInvalidtempl(res)	
			if err != nil  {
				http.Error(*res,"Internal server error", http.StatusInternalServerError)
				return err
			}
			return nil 
		}
		found_usr, usr_err := database.GetUser(email)	
		if usr_err != nil {
			wait(&start, WAIT_TIME_SEC)
			
			err = sendInvalidtempl(res)
			if err != nil {
				http.Error(*res, "Internal server error", http.StatusInternalServerError)
				return err
			}
			return usr_err
		}
		
		hashedPassword, hash_err := encrypt.Hash(password)
		if hash_err != nil {
			wait(&start, WAIT_TIME_SEC)
			err = sendInvalidtempl(res)
			if err != nil {
				http.Error(*res, "Internal server error", http.StatusInternalServerError)
				return err
			}
			return hash_err 
		}

		if found_usr.Password == hashedPassword {
			wait(&start, WAIT_TIME_SEC)
			token, err := jwt.CreateJWT(found_usr.ID)
			if err == nil {
				(*res).Header().Add("Cookie", "JWT="+token)

			} else {
				logrus.WithError(err).Info("Failed to create JWT")
			}
			
			// send todo home page with content
			
			
		} else {
			return sendInvalidtempl(res)
			// send invalid login
		}
		


	}
	return nil
}

func postRegister(res *http.ResponseWriter, req *http.Request) error {
	var start time.Time = time.Now()
	const WAIT_TIME_SEC = 2

	parse_err := req.ParseForm()

	if parse_err != nil {
		wait(&start, WAIT_TIME_SEC)
		return sendInvalidRegister(res, false, true, false)
	}
	pw1 := req.FormValue("password")
	pw2 := req.FormValue("confirmedPassword")
	email := req.FormValue("email")
	if pw1 == "" || pw2 == "" || email == ""|| pw1 != pw2 || len(pw1) < 8{
		wait(&start, WAIT_TIME_SEC)
		return sendInvalidRegister(res, false, true, false)
	} else if database.Exists(email) {
		wait(&start, WAIT_TIME_SEC)
		return sendInvalidRegister(res, false, false, true)
	}
	// Add email validation	
	hash_pw, hash_err := encrypt.Hash(pw1)
	if hash_err != nil {
		return hash_err
	}
	var new_uuid uuid.UUID = uuid.New()
	
	newUser := types.User{Email:email, Password:hash_pw, NotificationsGranted:false, ID: new_uuid.String()}	
	database.AddUser(newUser)	
	return nil
}



func sendInvalidRegister(res *http.ResponseWriter,user_exists, pw_correct, pw_match bool) error {
	
	var data registerData  = registerData{false, false, false, false}
	switch (true) {
		case user_exists :
			data.EmailUsedError = true
		case pw_correct :
			data.NoEmailError = true
		case pw_match:
			data.NoMatchError = true
		default:
			return fmt.Errorf("None of entered errors are true")

	}
	templ, err := template.ParseFiles("html/register.ejs")
	if err != nil {
		return err
	}

	return templ.Execute(*res, data) 
}

func wait(start *time.Time, wait_time_sec  int)  {
	var end time.Time = time.Now();

	var duration int64 = end.Sub(*start).Nanoseconds()
	
	var waitime int64 = (int64(wait_time_sec) * 1_000_000_000) - duration
	time.Sleep(time.Duration(waitime))

}


func sendInvalidtempl(res *http.ResponseWriter) (error) {
	templ, err := template.ParseFiles("html/login.ejs")		
	if err != nil {
		return err
	}

	data := loginData{true}
	return templ.Execute(*res, data) 

}
