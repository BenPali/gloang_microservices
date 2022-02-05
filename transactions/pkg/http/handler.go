package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/k4lii/golang_microservices/transactions/pkg/endpoint"
	"net/http"
)

// makeMakeOfferHandler creates the handler logic
func makeMakeOfferHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/make-offer", http1.NewServer(endpoints.MakeOfferEndpoint, decodeMakeOfferRequest, encodeMakeOfferResponse, options...))
}

// decodeMakeOfferRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMakeOfferRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.MakeOfferRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMakeOfferResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMakeOfferResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeManageOfferHandler creates the handler logic
func makeManageOfferHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/manage-offer", http1.NewServer(endpoints.ManageOfferEndpoint, decodeManageOfferRequest, encodeManageOfferResponse, options...))
}

// decodeManageOfferRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeManageOfferRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ManageOfferRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeManageOfferResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeManageOfferResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListOwnOffersHandler creates the handler logic
func makeListOwnOffersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/list-own-offers", http1.NewServer(endpoints.ListOwnOffersEndpoint, decodeListOwnOffersRequest, encodeListOwnOffersResponse, options...))
}

// decodeListOwnOffersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListOwnOffersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ListOwnOffersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListOwnOffersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListOwnOffersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
