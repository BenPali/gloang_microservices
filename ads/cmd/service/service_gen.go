// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/k4lii/golang_microservices/ads/pkg/endpoint"
	http1 "github.com/k4lii/golang_microservices/ads/pkg/http"
	service "github.com/k4lii/golang_microservices/ads/pkg/service"
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
		"CreateAd":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateAd", logger))},
		"DeleteAd":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteAd", logger))},
		"GetAd":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetAd", logger))},
		"GetUserAdsList": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserAdsList", logger))},
		"SearchAd":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SearchAd", logger))},
		"UpdateAd":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateAd", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["CreateAd"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateAd")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateAd"))}
	mw["UpdateAd"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UpdateAd")), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateAd"))}
	mw["DeleteAd"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteAd")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteAd"))}
	mw["GetAd"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetAd")), endpoint.InstrumentingMiddleware(duration.With("method", "GetAd"))}
	mw["SearchAd"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SearchAd")), endpoint.InstrumentingMiddleware(duration.With("method", "SearchAd"))}
	mw["GetUserAdsList"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetUserAdsList")), endpoint.InstrumentingMiddleware(duration.With("method", "GetUserAdsList"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateAd", "UpdateAd", "DeleteAd", "GetAd", "SearchAd", "GetUserAdsList"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
