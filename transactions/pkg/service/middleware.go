package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
	"github.com/k4lii/golang_microservices/auth/auth"
)

// Middleware describes a service middleware.
type Middleware func(TransactionsService) TransactionsService

type loggingMiddleware struct {
	logger log.Logger
	next   TransactionsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TransactionsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TransactionsService) TransactionsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) MakeOffer(ctx context.Context, jwt string, id uint, adID uint, price float64) (rs string, err error) {
	id, err = auth.DecodeToken(jwt)
	fmt.Println(id)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "KO", err
	}
	if price == 0 {
		return "KO, no price given", nil
	}
	return l.next.MakeOffer(ctx, jwt, id, adID, price)
}
func (l loggingMiddleware) ManageOffer(ctx context.Context, jwt string, id uint, transactionID uint, accepted bool) (rs string, err error) {
	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "KO", err
	}
	return l.next.ManageOffer(ctx, jwt, id, transactionID, accepted)
}
func (l loggingMiddleware) ListOwnOffers(ctx context.Context, jwt string, id uint) (rs string, offers interface{}, err error) {
	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "KO", nil, err
	}
	return l.next.ListOwnOffers(ctx, jwt, id)
}
