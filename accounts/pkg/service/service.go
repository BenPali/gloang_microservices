package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/k4lii/golang_microservices/auth/auth"
	"github.com/k4lii/golang_microservices/db/connection"
	"github.com/k4lii/golang_microservices/db/models"
)

// AccountsService describes the service.
type AccountsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	CreateAccount(ctx context.Context, username, password, mail string) (rs string, err error)
	DeleteAccount(ctx context.Context, jwt string, id uint) (rs string, err error)
	GetSelfAccount(ctx context.Context, jwt string, id uint) (rs string, profile interface{}, err error)
	UpdateAccount(ctx context.Context, username string, password string, mail string, jwt string, id uint) (rs string, err error)
	ReadOtherAccount(ctx context.Context, username string) (rs string, profile interface{}, err error)
	Login(ctx context.Context, username, password, mail string) (rs string, err error)
	AddToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error)
	SusbstractToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error)
}

type basicAccountsService struct{}

func (b *basicAccountsService) CreateAccount(ctx context.Context, username string, password string, mail string) (rs string, err error) {
	fmt.Println("Create account")

	db := connection.DB

	newAccount := models.Account{
		Username: username,
		Password: string(base64.StdEncoding.EncodeToString(sha256.New().Sum([]byte(password)))),
		Mail:     mail,
		Balance:  0,
	}

	res := db.Table("accounts").Create(&newAccount)
	if res.Error != nil {
		err = res.Error

		return rs, err
	}

	rs = "User created."

	// TODO implement the business logic of CreateAccount
	return rs, err
}
func (b *basicAccountsService) DeleteAccount(ctx context.Context, jwt string, id uint) (rs string, err error) {
	// TODO implement the business logic of DeleteAccount

	db := connection.DB

	var acc models.Account

	res := db.Table("accounts").Where("id = ?", id).Delete(&acc)
	if res.Error != nil {
		return "KO", res.Error
	}
	return "Account deleted", res.Error
}

// NewBasicAccountsService returns a naive, stateless implementation of AccountsService.
func NewBasicAccountsService() AccountsService {
	return &basicAccountsService{}
}

// New returns a AccountsService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountsService {
	var svc AccountsService = NewBasicAccountsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAccountsService) Login(ctx context.Context, username string, password string, mail string) (rs string, err error) {
	// TODO implement the business logic of Login

	db := connection.DB

	var acc models.Account

	hashed_passwd := string(base64.StdEncoding.EncodeToString(sha256.New().Sum([]byte(password))))

	res := db.Table("accounts").Where(
		map[string]interface{}{
			"username": username,
			"password": hashed_passwd,
		}).Or(map[string]interface{}{
		"mail":     mail,
		"password": hashed_passwd,
	}).Find(&acc)
	if res.Error != nil {
		return "KO", res.Error
	}

	token, err := auth.CreateToken(acc.ID)
	return token, err
}

func (b *basicAccountsService) UpdateAccount(ctx context.Context, username string, password string, mail string, jwt string, id uint) (rs string, err error) {
	// TODO implement the business logic of UpdateAccount
	db := connection.DB

	if username != "" {
		res := db.Table("accounts").Where("id = ?", id).Update("username", username)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	if password != "" {
		res := db.Table("accounts").Where("id = ?", id).Update("password", password)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	if mail != "" {
		res := db.Table("accounts").Where("id = ?", id).Update("mail", mail)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	return "Ok", err
}

func (b *basicAccountsService) AddToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error) {
	fmt.Println("Addtobalance")

	db := connection.DB
	var acc models.Account

	db.First(&acc, id)
	res := db.Table("accounts").Where("id = ?", id).Update("balance", amount+acc.Balance)
	if res.Error != nil {
		return "KO", res.Error
	}
	return "money added", err
}

func (b *basicAccountsService) GetSelfAccount(ctx context.Context, jwt string, id uint) (string, interface{}, error) {
	// TODO implement the business logic of GetSelfAccount

	db := connection.DB
	var acc models.Account
	res := db.Table("accounts").Select("username", "mail", "balance").Where("id = ?", id).Scan(&acc)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}
	ret := map[string]interface{}{
		"username": acc.Username,
		"mail":     acc.Mail,
		"balance":  acc.Balance,
	}
	return "OK", ret, nil
}

func (b *basicAccountsService) ReadOtherAccount(ctx context.Context, username string) (string, interface{}, error) {
	fmt.Println("ReadOtherAccount")

	var acc models.Account

	db := connection.DB
	res := db.Table("accounts").Select("username", "mail").Where("username = ?", username).Scan(&acc)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}
	ret := map[string]string{
		"username": acc.Username,
		"mail":     acc.Mail,
	}

	return "OK", ret, nil
}

func (b *basicAccountsService) SusbstractToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error) {
	fmt.Println("Substracttobalance")

	db := connection.DB
	var acc models.Account
	fmt.Println(jwt)
	fmt.Println(id)
	fmt.Println(amount)
	if jwt == os.Getenv("ADMIN_PASSWORD") {
		db.First(&acc, id)
		res := db.Table("accounts").Where("id = ?", id).Update("balance", acc.Balance-amount)
		if res.Error != nil {
			return "KO", res.Error
		}
		return "money substracted", nil
	}
	return "Ko", fmt.Errorf("You dont have access to this function")
}
