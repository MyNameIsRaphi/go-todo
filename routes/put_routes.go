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

	"github.com/sirupsen/logrus"
)

type Put struct {
	Request
}

func (P Put) Init() error {
	P.status = 200
	P.startTime = time.Now()
	b, err := io.ReadAll(P.req.Body)
	if err != nil {
		return err
	}
	if P.wait_time_sec == 0 {
		P.wait_time_sec = 2
	}
	P.body = b
	var d Default
	err = json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	P.auth = jwt.Validate(d.Token)
	if !P.auth {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func (P Put) UpdateToDo(res *http.ResponseWriter, req *http.Request) error {
	var body updateToDo
	var err error
	var foundUser types.User
	var resp defaultResponse
	// read response
	b, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.WithError(err).Error("failed to read request body")
		resp = defaultResponse{false, "failed to read data"}
		P.status = 400
		(*res).WriteHeader(P.status)
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON(res, resp)
	}

	err = json.Unmarshal(b, &body)

	if err != nil {
		P.status = 400
		logrus.WithError(err).Error("request has a invalid format")
		resp = defaultResponse{false, "request has a invalid format"}
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON(res, resp)
	}

	// verify token
	if !jwt.Validate(body.Token) {
		P.status = 401
		logrus.Infof("invalid token from %s to change todo", req.Host)
		resp = defaultResponse{false, "invalid token"}
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON(res, resp)
	}

	// read email
	email, err := jwt.GetEmail(body.Token)

	if err != nil {
		logrus.WithError(err).Error("failed to read email from token")
		P.status = 401
		resp = defaultResponse{false, "invalid token"}
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON(P.res, resp)
	}

	// get user

	foundUser, err = database.GetUser(email)

	if err != nil {
		P.status = http.StatusBadRequest
		logrus.WithError(err).Error("failed to get user")
		resp = defaultResponse{false, "no user found for email"}
		wait(&P.startTime, P.wait_time_sec)
		return sendJSON(res, resp)

	}

	// update todo
	var newUser = foundUser
	for i := 0; i < len(foundUser.ToDos); i++ {
		currentToDo := foundUser.ToDos[i]
		if currentToDo == body.OldToDo {
			newUser.ToDos[i] = body.NewToDo
		}
	}
	// update user
	err = database.UpdateUser(foundUser, newUser)

	if err != nil {
		resp = defaultResponse{false, "failed to update user"}
		P.status = 400

	} else {
		resp = defaultResponse{true, "successfully updated todo"}
		P.status = 200
	}
	(*res).WriteHeader(P.status)
	return sendJSON(res, resp)

}
