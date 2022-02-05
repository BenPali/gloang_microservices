package endpoint

import (
	"context"
	"fmt"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/accounts/pkg/service"
)

// CreateAccountRequest collects the request parameters for the CreateAccount method.
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

// CreateAccountResponse collects the response parameters for the CreateAccount method.
type CreateAccountResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateAccountEndpoint returns an endpoint that invokes CreateAccount on the service.
func MakeCreateAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		rs, err := s.CreateAccount(ctx, req.Username, req.Password, req.Mail)
		return CreateAccountResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAccountResponse) Failed() error {
	return r.Err
}

// DeleteAccountRequest collects the request parameters for the DeleteAccount method.
type DeleteAccountRequest struct {
	Jwt string `json:"jwt"`
}

// DeleteAccountResponse collects the response parameters for the DeleteAccount method.
type DeleteAccountResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteAccountEndpoint returns an endpoint that invokes DeleteAccount on the service.
func MakeDeleteAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAccountRequest)
		rs, err := s.DeleteAccount(ctx, req.Jwt, 0)
		return DeleteAccountResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteAccountResponse) Failed() error {
	return r.Err
}

// LoginRequest collects the request parameters for the CreateAccount method.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

// LoginResponse collects the response parameters for the CreateAccount method.
type LoginResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes CreateAccount on the service.
func MakeLoginEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		rs, err := s.Login(ctx, req.Username, req.Password, req.Mail)
		return LoginResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateAccount implements Service. Primarily useful in a client.
func (e Endpoints) CreateAccount(ctx context.Context, username string, password string, mail string) (rs string, err error) {
	request := CreateAccountRequest{
		Mail:     mail,
		Password: password,
		Username: username,
	}
	response, err := e.CreateAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAccountResponse).Rs, response.(CreateAccountResponse).Err
}

// DeleteAccount implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAccount(ctx context.Context, id uint) (rs string, err error) {
	request := DeleteAccountRequest{Jwt: fmt.Sprint(id)}
	response, err := e.DeleteAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAccountResponse).Rs, response.(DeleteAccountResponse).Err
}

func (e Endpoints) Login(ctx context.Context, username string, password string, mail string) (rs string, err error) {
	request := LoginRequest{
		Mail:     mail,
		Password: password,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Rs, response.(LoginResponse).Err
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// UpdateAccountRequest collects the request parameters for the UpdateAccount method.
type UpdateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	JWT      string `json:"jwt"`
	Id       uint   `json:"id"`
}

// UpdateAccountResponse collects the response parameters for the UpdateAccount method.
type UpdateAccountResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateAccountEndpoint returns an endpoint that invokes UpdateAccount on the service.
func MakeUpdateAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAccountRequest)
		rs, err := s.UpdateAccount(ctx, req.Username, req.Password, req.Mail, req.JWT, req.Id)
		return UpdateAccountResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateAccountResponse) Failed() error {
	return r.Err
}

// UpdateAccount implements Service. Primarily useful in a client.
func (e Endpoints) UpdateAccount(ctx context.Context, username string, password string, mail string, id uint) (rs string, err error) {
	request := UpdateAccountRequest{
		Id:       id,
		Mail:     mail,
		Password: password,
		Username: username,
	}
	response, err := e.UpdateAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateAccountResponse).Rs, response.(UpdateAccountResponse).Err
}

// AddToBalanceRequest collects the request parameters for the AddToBalance method.
type AddToBalanceRequest struct {
	Id     uint    `json:"id"`
	JWT    string  `json:"jwt"`
	Amount float64 `json:"amount"`
}

// AddToBalanceResponse collects the response parameters for the AddToBalance method.
type AddToBalanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeAddToBalanceEndpoint returns an endpoint that invokes AddToBalance on the service.
func MakeAddToBalanceEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddToBalanceRequest)
		rs, err := s.AddToBalance(ctx, req.JWT, req.Id, req.Amount)
		return AddToBalanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddToBalanceResponse) Failed() error {
	return r.Err
}

// AddToBalance implements Service. Primarily useful in a client.
func (e Endpoints) AddToBalance(ctx context.Context, id uint, amount float64) (rs string, err error) {
	request := AddToBalanceRequest{
		Amount: amount,
		Id:     id,
	}
	response, err := e.AddToBalanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddToBalanceResponse).Rs, response.(AddToBalanceResponse).Err
}

// GetSelfAccountRequest collects the request parameters for the GetSelfAccount method.
type GetSelfAccountRequest struct {
	Id  uint   `json:"id"`
	JWT string `json:"jwt"`
}

// GetSelfAccountResponse collects the response parameters for the GetSelfAccount method.
type GetSelfAccountResponse struct {
	Rs      string      `json:"rs"`
	Err     error       `json:"err"`
	Profile interface{} `json:"profile"`
}

// MakeGetSelfAccountEndpoint returns an endpoint that invokes GetSelfAccount on the service.
func MakeGetSelfAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSelfAccountRequest)
		rs, profile, err := s.GetSelfAccount(ctx, req.JWT, req.Id)
		return GetSelfAccountResponse{
			Err:     err,
			Rs:      rs,
			Profile: profile,
		}, nil
	}
}

// Failed implements Failer.
func (r GetSelfAccountResponse) Failed() error {
	return r.Err
}

// ReadOtherAccountRequest collects the request parameters for the ReadOtherAccount method.
type ReadOtherAccountRequest struct {
	Username string `json:"username"`
}

// ReadOtherAccountResponse collects the response parameters for the ReadOtherAccount method.
type ReadOtherAccountResponse struct {
	Rs      string      `json:"rs"`
	Err     error       `json:"err"`
	Profile interface{} `json:"profile"`
}

// MakeReadOtherAccountEndpoint returns an endpoint that invokes ReadOtherAccount on the service.
func MakeReadOtherAccountEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadOtherAccountRequest)
		rs, profile, err := s.ReadOtherAccount(ctx, req.Username)
		return ReadOtherAccountResponse{
			Err:     err,
			Rs:      rs,
			Profile: profile,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadOtherAccountResponse) Failed() error {
	return r.Err
}

// GetSelfAccount implements Service. Primarily useful in a client.
func (e Endpoints) GetSelfAccount(ctx context.Context, id uint) (rs string, err error) {
	request := GetSelfAccountRequest{Id: id}
	response, err := e.GetSelfAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetSelfAccountResponse).Rs, response.(GetSelfAccountResponse).Err
}

// ReadOtherAccount implements Service. Primarily useful in a client.
func (e Endpoints) ReadOtherAccount(ctx context.Context, username string) (rs string, err error) {
	request := ReadOtherAccountRequest{Username: username}
	response, err := e.ReadOtherAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadOtherAccountResponse).Rs, response.(ReadOtherAccountResponse).Err
}

// SusbstractToBalanceRequest collects the request parameters for the SusbstractToBalance method.
type SusbstractToBalanceRequest struct {
	Jwt    string  `json:"jwt"`
	Id     uint    `json:"id"`
	Amount float64 `json:"amount"`
}

// SusbstractToBalanceResponse collects the response parameters for the SusbstractToBalance method.
type SusbstractToBalanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSusbstractToBalanceEndpoint returns an endpoint that invokes SusbstractToBalance on the service.
func MakeSusbstractToBalanceEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SusbstractToBalanceRequest)
		rs, err := s.SusbstractToBalance(ctx, req.Jwt, req.Id, req.Amount)
		return SusbstractToBalanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SusbstractToBalanceResponse) Failed() error {
	return r.Err
}

// SusbstractToBalance implements Service. Primarily useful in a client.
func (e Endpoints) SusbstractToBalance(ctx context.Context, jwt string, id uint, amount float64) (rs string, err error) {
	request := SusbstractToBalanceRequest{
		Amount: amount,
		Id:     id,
		Jwt:    jwt,
	}
	response, err := e.SusbstractToBalanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SusbstractToBalanceResponse).Rs, response.(SusbstractToBalanceResponse).Err
}
