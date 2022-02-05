package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/messages/pkg/service"
)

// AddMessageRequest collects the request parameters for the AddMessage method.
type AddMessageRequest struct {
	Jwt           string `json:"jwt"`
	Id            uint   `json:"id"`
	TransactionId uint   `json:"transaction_id"`
	Messages      string `json:"messages"`
}

// AddMessageResponse collects the response parameters for the AddMessage method.
type AddMessageResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeAddMessageEndpoint returns an endpoint that invokes AddMessage on the service.
func MakeAddMessageEndpoint(s service.MessagesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddMessageRequest)
		rs, err := s.AddMessage(ctx, req.Jwt, req.Id, req.TransactionId, req.Messages)
		return AddMessageResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddMessageResponse) Failed() error {
	return r.Err
}

// GetMessageRequest collects the request parameters for the GetMessage method.
type GetMessageRequest struct {
	Jwt           string `json:"jwt"`
	Id            uint   `json:"id"`
	TransactionId uint   `json:"transaction_id"`
}

// GetMessageResponse collects the response parameters for the GetMessage method.
type GetMessageResponse struct {
	Rs       string      `json:"rs"`
	Messages interface{} `json:"messages"`
	Err      error       `json:"err"`
}

// MakeGetMessageEndpoint returns an endpoint that invokes GetMessage on the service.
func MakeGetMessageEndpoint(s service.MessagesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMessageRequest)
		rs, messages, err := s.GetMessage(ctx, req.Jwt, req.Id, req.TransactionId)
		return GetMessageResponse{
			Err:      err,
			Messages: messages,
			Rs:       rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetMessageResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AddMessage implements Service. Primarily useful in a client.
func (e Endpoints) AddMessage(ctx context.Context, jwt string, id uint, transactionId uint, messages string) (rs string, err error) {
	request := AddMessageRequest{
		Id:            id,
		Jwt:           jwt,
		Messages:      messages,
		TransactionId: transactionId,
	}
	response, err := e.AddMessageEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddMessageResponse).Rs, response.(AddMessageResponse).Err
}

// GetMessage implements Service. Primarily useful in a client.
func (e Endpoints) GetMessage(ctx context.Context, jwt string, id uint, transactionId uint) (rs string, err error) {
	request := GetMessageRequest{
		Id:            id,
		Jwt:           jwt,
		TransactionId: transactionId,
	}
	response, err := e.GetMessageEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetMessageResponse).Rs, response.(GetMessageResponse).Err
}
