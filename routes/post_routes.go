package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"todo/database"
	"todo/jwt"
	"todo/types"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Post struct {
	Request
}

func (P Post) Init() error {
	P.startTime = time.Now()
	b, err := io.ReadAll(P.req.Body)
	P.status = 200
	if err != nil {
		P.status = 400
		(*P.res).WriteHeader(P.status)
		return err
	}
	P.body = b

	var d Default

	err = json.Unmarshal(b, &d)

	if err != nil {
		P.status = 400
		(*P.res).WriteHeader(P.status)
		return err
	}

	P.token = d.Token
	P.auth = false
	if d.Token != "" {

		P.auth = jwt.Validate(P.token)
	}
	if !P.auth {
		P.status = 401
		(*P.res).WriteHeader(P.status)
		return fmt.Errorf("invalid token")
	}
	return nil
}

func (P Post) Login() error {

	var body loginRequest
	res := P.res
	err := json.Unmarshal(P.body, &body)
	if err != nil {
		return sendJSON[loginData](res, loginData{false, ""})
	}

	var data loginData
	// check if username and password is valid
	var isValid bool = database.CheckCreditentials(body.Email, body.Password)

	if isValid {
		token, err := jwt.CreateJWT(body.Email)
		if err == nil {
			data = loginData{true, token}
		}
	} else {
		data = loginData{false, ""}
	}

	// send response
	(*res).Header().Set("Content-Type", "application/json")
	wait(&P.startTime, P.wait_time_sec)
	return sendJSON[loginData](res, data)

}

func (P Post) Register() error {
	var body registerRequest
	var data registerData
	var inValid = false
	var err error
	b, err := io.ReadAll(P.req.Body)
	data = registerData{false, false, false, true}
	if err != nil {
		(*P.res).WriteHeader(400)
		return sendJSON[registerData](P.res, data)
	}
	P.body = b
	// read request body

	err = json.Unmarshal(P.body, &body)
	if err != nil {
		logrus.Error(err)
		data = registerData{false, false, false, true}
		P.status = 400
		(*P.res).WriteHeader(P.status)
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON[registerData](P.res, data)

	}
	pw1 := body.ConfirmedPassword
	pw2 := body.Password
	email := body.Email

	// validate creditentials
	if pw1 == "" || pw2 == "" || email == "" || pw1 != pw2 || len(pw1) < 8 {
		inValid = true
		logrus.Info("Email or password don't match")
		data = registerData{false, false, true, false}

	} else if database.Exists(email) {
		inValid = true
		logrus.Info("User already exists")
		data = registerData{false, true, false, false}
	}

	// send response
	if inValid {
		P.status = 400
		(*P.res).WriteHeader(P.status)
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON[registerData](P.res, data)

	}
	// Add email validation

	var new_uuid uuid.UUID = uuid.New()
	newUser := types.User{Email: email, Password: pw1, NotificationsGranted: false, ID: new_uuid.String()}
	err = database.AddUser(newUser)
	if err != nil {
		P.status = http.StatusInternalServerError
		logrus.WithError(err).Error("failed to add user to database")
		data = registerData{false, false, false, false}

	} else {
		P.status = 200
		logrus.Info("successfully added user to database")
		data = registerData{true, false, false, false}
	}
	(*P.res).WriteHeader(P.status)
	wait(&P.startTime, P.wait_time_sec)
	return sendJSON[registerData](P.res, data)

}
func (P Post) AddToDo() error {
	var body addToDoRequest
	var err error
	var email string
	var user types.User

	err = json.Unmarshal(P.body, &body)

	if err != nil {
		return err
	}
	email, err = jwt.GetEmail(P.token)
	if err != nil {
		return err
	}

	user, err = database.GetUser(email)

	if err != nil {
		return err
	}

	user.ToDos = append(user.ToDos, body.New)

	err = sendJSON(P.res, defaultResponse{true, "Successfully added todo"})
	return err
}

type addToDoRequest struct {
	Token string
	New   types.ToDo
}

type registerData struct {
	Success        bool
	EmailUsedError bool
	InCorrectData  bool
	NoBodyError    bool
}

type loginData struct {
	Auth  bool
	Token string
}

type registerRequest struct {
	Email                string
	Password             string
	ConfirmedPassword    string
	FirstName            string
	LastName             string
	NotificationsGranted bool
}

type loginRequest struct {
	Password string
	Email    string
}
