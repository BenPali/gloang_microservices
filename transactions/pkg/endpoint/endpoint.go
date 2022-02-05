package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/transactions/pkg/service"
)

// MakeOfferRequest collects the request parameters for the MakeOffer method.
type MakeOfferRequest struct {
	Jwt   string  `json:"jwt"`
	Id    uint    `json:"id"`
	AdID  uint    `json:"ad_id"`
	Price float64 `json:"price"`
}

// MakeOfferResponse collects the response parameters for the MakeOffer method.
type MakeOfferResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeMakeOfferEndpoint returns an endpoint that invokes MakeOffer on the service.
func MakeMakeOfferEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MakeOfferRequest)
		rs, err := s.MakeOffer(ctx, req.Jwt, req.Id, req.AdID, req.Price)
		return MakeOfferResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r MakeOfferResponse) Failed() error {
	return r.Err
}

// ManageOfferRequest collects the request parameters for the ManageOffer method.
type ManageOfferRequest struct {
	Jwt           string `json:"jwt"`
	Id            uint   `json:"id"`
	TransactionID uint   `json:"transaction_id"`
	Accepted      bool   `json:"accepted"`
}

// ManageOfferResponse collects the response parameters for the ManageOffer method.
type ManageOfferResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeManageOfferEndpoint returns an endpoint that invokes ManageOffer on the service.
func MakeManageOfferEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ManageOfferRequest)
		rs, err := s.ManageOffer(ctx, req.Jwt, req.Id, req.TransactionID, req.Accepted)
		return ManageOfferResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ManageOfferResponse) Failed() error {
	return r.Err
}

// ListOwnOffersRequest collects the request parameters for the ListOwnOffers method.
type ListOwnOffersRequest struct {
	Jwt string `json:"jwt"`
	Id  uint   `json:"id"`
}

// ListOwnOffersResponse collects the response parameters for the ListOwnOffers method.
type ListOwnOffersResponse struct {
	Rs     string      `json:"rs"`
	Offers interface{} `json:"offers"`
	Err    error       `json:"err"`
}

// MakeListOwnOffersEndpoint returns an endpoint that invokes ListOwnOffers on the service.
func MakeListOwnOffersEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListOwnOffersRequest)
		rs, offers, err := s.ListOwnOffers(ctx, req.Jwt, req.Id)
		return ListOwnOffersResponse{
			Err:    err,
			Offers: offers,
			Rs:     rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ListOwnOffersResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// MakeOffer implements Service. Primarily useful in a client.
func (e Endpoints) MakeOffer(ctx context.Context, jwt string, id uint, adID uint, price float64) (rs string, err error) {
	request := MakeOfferRequest{
		AdID:  adID,
		Id:    id,
		Jwt:   jwt,
		Price: price,
	}
	response, err := e.MakeOfferEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MakeOfferResponse).Rs, response.(MakeOfferResponse).Err
}

// ManageOffer implements Service. Primarily useful in a client.
func (e Endpoints) ManageOffer(ctx context.Context, jwt string, id uint, transactionID uint, accepted bool) (rs string, err error) {
	request := ManageOfferRequest{
		Accepted:      accepted,
		Id:            id,
		Jwt:           jwt,
		TransactionID: transactionID,
	}
	response, err := e.ManageOfferEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ManageOfferResponse).Rs, response.(ManageOfferResponse).Err
}

// ListOwnOffers implements Service. Primarily useful in a client.
func (e Endpoints) ListOwnOffers(ctx context.Context, jwt string, id uint) (rs string, offers interface{}, err error) {
	request := ListOwnOffersRequest{
		Id:  id,
		Jwt: jwt,
	}
	response, err := e.ListOwnOffersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListOwnOffersResponse).Rs, response.(ListOwnOffersResponse).Offers, response.(ListOwnOffersResponse).Err
}
