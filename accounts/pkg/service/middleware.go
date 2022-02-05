package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	log "github.com/go-kit/log"
	"github.com/k4lii/golang_microservices/auth/auth"
)

// Middleware describes a service middleware.
type Middleware func(AccountsService) AccountsService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountsService) AccountsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateAccount(ctx context.Context, username string, password string, email string) (rs string, err error) {
	if username == "" || password == "" || email == "" {
		return "KO", errors.New("could not create account: mandatory field is empty")
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		return "KO", err
	}

	return l.next.CreateAccount(ctx, username, password, email)
}
func (l loggingMiddleware) DeleteAccount(ctx context.Context, jwt string, id uint) (string, error) {
	if id != 0 {
		l.logger.Log("id does not have the expected zero value. value", id)
	}

	id, err := auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())

		return "KO", err
	}

	return l.next.DeleteAccount(ctx, jwt, id)
}

func (l loggingMiddleware) Login(ctx context.Context, username string, password string, email string) (rs string, err error) {

	if (username == "" && email == "") || password == "" {
		return "KO", errors.New("could not login: mandatory field is empty")
	}

	_, err = mail.ParseAddress(email)
	if err != nil && email != "" {
		return "KO", err
	}
	return l.next.Login(ctx, username, password, email)
}

func (l loggingMiddleware) UpdateAccount(ctx context.Context, username string, password string, email string, jwt string, id uint) (rs string, err error) {
	if id != 0 {
		l.logger.Log("id does not have the expected zero value. value", id)
	}

	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())

		return "KO", err
	}

	if username == "" && password == "" && email == "" {
		return "KO", errors.New("could not update account: all fields are empty, nothing to update")
	}

	_, err = mail.ParseAddress(email)
	if err != nil && email != "" {
		return "KO: Not a mail", err
	}
	return l.next.UpdateAccount(ctx, username, password, email, jwt, id)
}

func (l loggingMiddleware) AddToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error) {
	if id != 0 {
		l.logger.Log("id does not have the expected zero value. value", id)
	}

	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())

		return "KO", err
	}
	if amount == 0 {
		return "KO : no specific value", err
	}
	return l.next.AddToBalance(ctx, jwt, id, amount)
}

func (l loggingMiddleware) GetSelfAccount(ctx context.Context, jwt string, id uint) (rs string, profile interface{}, err error) {
	if id != 0 {
		l.logger.Log("id does not have the expected zero value. value", id)
	}

	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())

		return "KO", nil, err
	}
	return l.next.GetSelfAccount(ctx, jwt, id)
}
func (l loggingMiddleware) ReadOtherAccount(ctx context.Context, username string) (rs string, profile interface{}, err error) {
	if username == "" {
		return "Ko: no username provided", nil, errors.New("no username provided")
	}
	return l.next.ReadOtherAccount(ctx, username)
}

func (l loggingMiddleware) SusbstractToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "SusbstractToBalance", "jwt", jwt, "id", id, "amount", amount, "rs", rs, "err", err)
	}()
	return l.next.SusbstractToBalance(ctx, jwt, id, amount)
}
