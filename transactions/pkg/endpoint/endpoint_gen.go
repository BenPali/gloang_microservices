// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/k4lii/golang_microservices/transactions/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	MakeOfferEndpoint     endpoint.Endpoint
	ManageOfferEndpoint   endpoint.Endpoint
	ListOwnOffersEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.TransactionsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		ListOwnOffersEndpoint: MakeListOwnOffersEndpoint(s),
		MakeOfferEndpoint:     MakeMakeOfferEndpoint(s),
		ManageOfferEndpoint:   MakeManageOfferEndpoint(s),
	}
	for _, m := range mdw["MakeOffer"] {
		eps.MakeOfferEndpoint = m(eps.MakeOfferEndpoint)
	}
	for _, m := range mdw["ManageOffer"] {
		eps.ManageOfferEndpoint = m(eps.ManageOfferEndpoint)
	}
	for _, m := range mdw["ListOwnOffers"] {
		eps.ListOwnOffersEndpoint = m(eps.ListOwnOffersEndpoint)
	}
	return eps
}
