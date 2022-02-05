package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
	"github.com/k4lii/golang_microservices/auth/auth"
)

// Middleware describes a service middleware.
type Middleware func(MessagesService) MessagesService

type loggingMiddleware struct {
	logger log.Logger
	next   MessagesService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a MessagesService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MessagesService) MessagesService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) AddMessage(ctx context.Context, jwt string, id uint, transactionId uint, messages string) (rs string, err error) {
	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "KO", err
	}
	return l.next.AddMessage(ctx, jwt, id, transactionId, messages)
}
func (l loggingMiddleware) GetMessage(ctx context.Context, jwt string, id uint, transactionId uint) (rs string,	messages interface{}, err error) {
	id, err = auth.DecodeToken(jwt)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "KO",nil, err
	}
	return l.next.GetMessage(ctx, jwt, id, transactionId)
}
