package controller

import (
	"github.com/camotts/bggroup/controller/store"
	"github.com/camotts/bggroup/rest_server_bggroup/operations/account"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type registerHandler struct{}

func newRegisterHandler() *registerHandler {
	return &registerHandler{}
}

func (h *registerHandler) Handle(params account.RegisterParams) middleware.Responder {
	if params.Body == nil || params.Body.Email == "" || params.Body.Password == "" {
		logrus.Error("missing email or password")
		return account.NewRegisterNotFound()
	}

	tx, err := str.Begin()
	if err != nil {
		logrus.Errorf("unable to start transaction for email %v: %v", params.Body.Email, err)
		return account.NewRegisterInternalServerError()
	}
	defer tx.Rollback()

	hash, err := hashPassword(params.Body.Password)
	if err != nil {
		logrus.Errorf("unable to hash password for email %v: %v", params.Body.Email, err)
		return account.NewRegisterInternalServerError()
	}
	a := &store.Account{
		Email:    params.Body.Email,
		Password: string(hash),
	}

	if err := str.CreateAccount(tx, a); err != nil {
		logrus.Errorf("unable to create account for email %v: %v", params.Body.Email, err)
		return account.NewRegisterInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("unable to commit '%v': %v", params.Body.Email, err)
		return account.NewRegisterInternalServerError()
	}
	logrus.Infof("created account '%v'", a.Email)

	return account.NewRegisterOK()
}

func hashPassword(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return
}
