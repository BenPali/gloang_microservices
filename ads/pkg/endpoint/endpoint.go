package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/ads/pkg/service"
)

// CreateAdRequest collects the request parameters for the CreateAd method.
type CreateAdRequest struct {
	Jwt     string  `json:"jwt"`
	Title   string  `json:"title"`
	Desc    string  `json:"description"`
	Picture string  `json:"picture"`
	Price   float64 `json:"price"`
}

// CreateAdResponse collects the response parameters for the CreateAd method.
type CreateAdResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateAdEndpoint returns an endpoint that invokes CreateAd on the service.
func MakeCreateAdEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAdRequest)
		rs, err := s.CreateAd(ctx, req.Jwt, 0, req.Title, req.Desc, req.Picture, req.Price)
		return CreateAdResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAdResponse) Failed() error {
	return r.Err
}

// UpdateAdRequest collects the request parameters for the UpdateAd method.
type UpdateAdRequest struct {
	Jwt     string  `json:"jwt"`
	AddID   uint    `json:"add_id"`
	Title   string  `json:"title"`
	Desc    string  `json:"description"`
	Picture string  `json:"picture"`
	Price   float64 `json:"price"`
}

// UpdateAdResponse collects the response parameters for the UpdateAd method.
type UpdateAdResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateAdEndpoint returns an endpoint that invokes UpdateAd on the service.
func MakeUpdateAdEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAdRequest)
		rs, err := s.UpdateAd(ctx, req.Jwt, 0, req.AddID, req.Title, req.Desc, req.Picture, req.Price)
		return UpdateAdResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateAdResponse) Failed() error {
	return r.Err
}

// DeleteAdRequest collects the request parameters for the DeleteAd method.
type DeleteAdRequest struct {
	Jwt   string `json:"jwt"`
	AddID int    `json:"add_id"`
}

// DeleteAdResponse collects the response parameters for the DeleteAd method.
type DeleteAdResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteAdEndpoint returns an endpoint that invokes DeleteAd on the service.
func MakeDeleteAdEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAdRequest)
		rs, err := s.DeleteAd(ctx, req.Jwt, 0, uint(req.AddID))
		return DeleteAdResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteAdResponse) Failed() error {
	return r.Err
}

// GetAdRequest collects the request parameters for the GetAd method.
type GetAdRequest struct {
	AddID uint `json:"add_id"`
}

// GetAdResponse collects the response parameters for the GetAd method.
type GetAdResponse struct {
	S0 string      `json:"s0"`
	I1 interface{} `json:"i1"`
	E2 error       `json:"e2"`
}

// MakeGetAdEndpoint returns an endpoint that invokes GetAd on the service.
func MakeGetAdEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdRequest)
		s0, i1, e2 := s.GetAd(ctx, req.AddID)
		return GetAdResponse{
			E2: e2,
			I1: i1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdResponse) Failed() error {
	return r.E2
}

// SearchAdRequest collects the request parameters for the SearchAd method.
type SearchAdRequest struct {
	Keyword string `json:"keyword"`
}

// SearchAdResponse collects the response parameters for the SearchAd method.
type SearchAdResponse struct {
	S0 string      `json:"s0"`
	I1 interface{} `json:"i1"`
	E2 error       `json:"e2"`
}

// MakeSearchAdEndpoint returns an endpoint that invokes SearchAd on the service.
func MakeSearchAdEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchAdRequest)
		s0, i1, e2 := s.SearchAd(ctx, req.Keyword)
		return SearchAdResponse{
			E2: e2,
			I1: i1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SearchAdResponse) Failed() error {
	return r.E2
}

// GetUserAdsListRequest collects the request parameters for the GetUserAdsList method.
type GetUserAdsListRequest struct {
	PosterID uint `json:"poster_id"`
}

// GetUserAdsListResponse collects the response parameters for the GetUserAdsList method.
type GetUserAdsListResponse struct {
	S0 string      `json:"s0"`
	I1 interface{} `json:"i1"`
	E2 error       `json:"e2"`
}

// MakeGetUserAdsListEndpoint returns an endpoint that invokes GetUserAdsList on the service.
func MakeGetUserAdsListEndpoint(s service.AdsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserAdsListRequest)
		s0, i1, e2 := s.GetUserAdsList(ctx, req.PosterID)
		return GetUserAdsListResponse{
			E2: e2,
			I1: i1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserAdsListResponse) Failed() error {
	return r.E2
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateAd implements Service. Primarily useful in a client.
func (e Endpoints) CreateAd(ctx context.Context, jwt string, ID uint, title string, desc string, picture string, price float64) (rs string, err error) {
	request := CreateAdRequest{
		Desc:    desc,
		Jwt:     jwt,
		Picture: picture,
		Price:   price,
		Title:   title,
	}
	response, err := e.CreateAdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAdResponse).Rs, response.(CreateAdResponse).Err
}

// UpdateAd implements Service. Primarily useful in a client.
func (e Endpoints) UpdateAd(ctx context.Context, jwt string, ID uint, addID uint, title string, desc string, picture string, price float64) (rs string, err error) {
	request := UpdateAdRequest{
		AddID:   addID,
		Desc:    desc,
		Jwt:     jwt,
		Picture: picture,
		Price:   price,
		Title:   title,
	}
	response, err := e.UpdateAdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateAdResponse).Rs, response.(UpdateAdResponse).Err
}

// DeleteAd implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAd(ctx context.Context, jwt string, ID uint, addID uint) (rs string, err error) {
	request := DeleteAdRequest{
		AddID: int(addID),
		Jwt:   jwt,
	}
	response, err := e.DeleteAdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAdResponse).Rs, response.(DeleteAdResponse).Err
}

// GetAd implements Service. Primarily useful in a client.
func (e Endpoints) GetAd(ctx context.Context, addID uint) (s0 string, i1 interface{}, e2 error) {
	request := GetAdRequest{AddID: addID}
	response, err := e.GetAdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdResponse).S0, response.(GetAdResponse).I1, response.(GetAdResponse).E2
}

// SearchAd implements Service. Primarily useful in a client.
func (e Endpoints) SearchAd(ctx context.Context, keyword string) (s0 string, i1 interface{}, e2 error) {
	request := SearchAdRequest{Keyword: keyword}
	response, err := e.SearchAdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SearchAdResponse).S0, response.(SearchAdResponse).I1, response.(SearchAdResponse).E2
}

// GetUserAdsList implements Service. Primarily useful in a client.
func (e Endpoints) GetUserAdsList(ctx context.Context, posterID uint) (s0 string, i1 interface{}, e2 error) {
	request := GetUserAdsListRequest{PosterID: posterID}
	response, err := e.GetUserAdsListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserAdsListResponse).S0, response.(GetUserAdsListResponse).I1, response.(GetUserAdsListResponse).E2
}
