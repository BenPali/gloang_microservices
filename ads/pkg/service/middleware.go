package service

import (
	"context"
	"errors"
	"fmt"

	log "github.com/go-kit/log"
	"github.com/k4lii/golang_microservices/auth/auth"
)

// Middleware describes a service middleware.
type Middleware func(AdsService) AdsService

type loggingMiddleware struct {
	logger log.Logger
	next   AdsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AdsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AdsService) AdsService {
		return &loggingMiddleware{logger, next}
	}

}

var errDefValue = errors.New("incorrect default value")

func (l loggingMiddleware) CreateAd(ctx context.Context, jwt string, ID uint, title string, desc string, picture string, price float64) (string, error) {
	if ID != 0 {
		return "ko", fmt.Errorf("%w: expected 0, got %d", errDefValue, ID)
	}

	var err error

	if title == "" || desc == "" || price <= 0 {
		return "ko", errors.New("title, desc and prices are mandatory")
	}

	ID, err = auth.DecodeToken(jwt)
	if err != nil {
		return "ko", err
	}

	return l.next.CreateAd(ctx, jwt, ID, title, desc, picture, price)
}

func (l loggingMiddleware) UpdateAd(ctx context.Context, jwt string, ID uint, addID uint, title string, desc string, picture string, price float64) (string, error) {
	if ID != 0 {
		return "ko", fmt.Errorf("%w: expected 0, got %d", errDefValue, ID)
	}

	var err error

	ID, err = auth.DecodeToken(jwt)
	if err != nil {
		return "ko", err
	}

	if title == "" && desc == "" && price <= 0 {
		return "ko", errors.New("title, desc and prices are mandatory")
	}

	return l.next.UpdateAd(ctx, jwt, ID, addID, title, desc, picture, price)
}
func (l loggingMiddleware) DeleteAd(ctx context.Context, jwt string, ID uint, addID uint) (string, error) {
	if ID != 0 {
		return "ko", fmt.Errorf("%w: expected 0, got %d", errDefValue, ID)
	}

	var err error

	ID, err = auth.DecodeToken(jwt)
	if err != nil {
		return "ko", err
	}

	return l.next.DeleteAd(ctx, jwt, ID, addID)
}
func (l loggingMiddleware) GetAd(ctx context.Context, addID uint) (s0 string, i1 interface{}, e2 error) {
	return l.next.GetAd(ctx, addID)
}
func (l loggingMiddleware) SearchAd(ctx context.Context, keyword string) (s0 string, i1 interface{}, e2 error) {
	if keyword == "" {
		return "ko", nil, errors.New("cannot search for an empty keyword")
	}
	return l.next.SearchAd(ctx, keyword)
}
func (l loggingMiddleware) GetUserAdsList(ctx context.Context, posterID uint) (s0 string, i1 interface{}, e2 error) {

	return l.next.GetUserAdsList(ctx, posterID)
}
