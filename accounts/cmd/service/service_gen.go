// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/k4lii/golang_microservices/accounts/pkg/endpoint"
	http1 "github.com/k4lii/golang_microservices/accounts/pkg/http"
	service "github.com/k4lii/golang_microservices/accounts/pkg/service"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"AddToBalance":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddToBalance", logger))},
		"CreateAccount":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateAccount", logger))},
		"DeleteAccount":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteAccount", logger))},
		"GetSelfAccount":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetSelfAccount", logger))},
		"Login":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Login", logger))},
		"ReadOtherAccount":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ReadOtherAccount", logger))},
		"SusbstractToBalance": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SusbstractToBalance", logger))},
		"UpdateAccount":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateAccount", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["CreateAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateAccount")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateAccount"))}
	mw["DeleteAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteAccount")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteAccount"))}
	mw["GetSelfAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetSelfAccount")), endpoint.InstrumentingMiddleware(duration.With("method", "GetSelfAccount"))}
	mw["UpdateAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UpdateAccount")), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateAccount"))}
	mw["ReadOtherAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ReadOtherAccount")), endpoint.InstrumentingMiddleware(duration.With("method", "ReadOtherAccount"))}
	mw["Login"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Login")), endpoint.InstrumentingMiddleware(duration.With("method", "Login"))}
	mw["AddToBalance"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddToBalance")), endpoint.InstrumentingMiddleware(duration.With("method", "AddToBalance"))}
	mw["SusbstractToBalance"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SusbstractToBalance")), endpoint.InstrumentingMiddleware(duration.With("method", "SusbstractToBalance"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateAccount", "DeleteAccount", "GetSelfAccount", "UpdateAccount", "ReadOtherAccount", "Login", "AddToBalance", "SusbstractToBalance"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
