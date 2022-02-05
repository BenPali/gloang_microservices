package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/k4lii/golang_microservices/accounts/pkg/endpoint"
)

// makeCreateAccountHandler creates the handler logic
func makeCreateAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-account", http1.NewServer(endpoints.CreateAccountEndpoint, decodeCreateAccountRequest, encodeCreateAccountResponse, options...))
}

// decodeCreateAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAccountHandler creates the handler logic
func makeDeleteAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-account", http1.NewServer(endpoints.DeleteAccountEndpoint, decodeDeleteAccountRequest, encodeDeleteAccountResponse, options...))
}

// decodeDeleteAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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

// makeLoginHandler creates the handler logic
func makeLoginHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/login", http1.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...))
}

// decodeLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateAccountHandler creates the handler logic
func makeUpdateAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-account", http1.NewServer(endpoints.UpdateAccountEndpoint, decodeUpdateAccountRequest, encodeUpdateAccountResponse, options...))
}

// decodeUpdateAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddToBalanceHandler creates the handler logic
func makeAddToBalanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/add-to-balance", http1.NewServer(endpoints.AddToBalanceEndpoint, decodeAddToBalanceRequest, encodeAddToBalanceResponse, options...))
}

// decodeAddToBalanceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddToBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddToBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddToBalanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddToBalanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetSelfAccountHandler creates the handler logic
func makeGetSelfAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-self-account", http1.NewServer(endpoints.GetSelfAccountEndpoint, decodeGetSelfAccountRequest, encodeGetSelfAccountResponse, options...))
}

// decodeGetSelfAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetSelfAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetSelfAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetSelfAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetSelfAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadOtherAccountHandler creates the handler logic
func makeReadOtherAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/read-other-account", http1.NewServer(endpoints.ReadOtherAccountEndpoint, decodeReadOtherAccountRequest, encodeReadOtherAccountResponse, options...))
}

// decodeReadOtherAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadOtherAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadOtherAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadOtherAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadOtherAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSusbstractToBalanceHandler creates the handler logic
func makeSusbstractToBalanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/susbstract-to-balance", http1.NewServer(endpoints.SusbstractToBalanceEndpoint, decodeSusbstractToBalanceRequest, encodeSusbstractToBalanceResponse, options...))
}

// decodeSusbstractToBalanceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSusbstractToBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SusbstractToBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSusbstractToBalanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSusbstractToBalanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
