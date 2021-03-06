// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/ads/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateAdEndpoint       endpoint.Endpoint
	UpdateAdEndpoint       endpoint.Endpoint
	DeleteAdEndpoint       endpoint.Endpoint
	GetAdEndpoint          endpoint.Endpoint
	SearchAdEndpoint       endpoint.Endpoint
	GetUserAdsListEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.AdsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateAdEndpoint:       MakeCreateAdEndpoint(s),
		DeleteAdEndpoint:       MakeDeleteAdEndpoint(s),
		GetAdEndpoint:          MakeGetAdEndpoint(s),
		GetUserAdsListEndpoint: MakeGetUserAdsListEndpoint(s),
		SearchAdEndpoint:       MakeSearchAdEndpoint(s),
		UpdateAdEndpoint:       MakeUpdateAdEndpoint(s),
	}
	for _, m := range mdw["CreateAd"] {
		eps.CreateAdEndpoint = m(eps.CreateAdEndpoint)
	}
	for _, m := range mdw["UpdateAd"] {
		eps.UpdateAdEndpoint = m(eps.UpdateAdEndpoint)
	}
	for _, m := range mdw["DeleteAd"] {
		eps.DeleteAdEndpoint = m(eps.DeleteAdEndpoint)
	}
	for _, m := range mdw["GetAd"] {
		eps.GetAdEndpoint = m(eps.GetAdEndpoint)
	}
	for _, m := range mdw["SearchAd"] {
		eps.SearchAdEndpoint = m(eps.SearchAdEndpoint)
	}
	for _, m := range mdw["GetUserAdsList"] {
		eps.GetUserAdsListEndpoint = m(eps.GetUserAdsListEndpoint)
	}
	return eps
}
