package server

import (
	"net/http"

	"google.golang.org/grpc"
)

// Service is the most basic interface of a service that can be received and
// hosted by a Server.
type Service interface {
	Prefix() string

	// Middleware is a hook to enable services to add
	// any additional middleware.
	Middleware(http.Handler) http.Handler
}

// SimpleService is an interface defining a service that
// is made up of http.HandlerFuncs.
type SimpleService interface {
	Service

	// route - method - func
	Endpoints() map[string]map[string]http.HandlerFunc
}

// JSONService is an interface defining a service that
// is made up of JSONEndpoints.
type JSONService interface {
	Service

	// Ensure that the route syntax is compatible with the router
	// implementation chosen in cfg.RouterType.
	// route - method - func
	JSONEndpoints() map[string]map[string]JSONEndpoint
	JSONMiddleware(JSONEndpoint) JSONEndpoint
}

// MixedService is an interface defining service that
// offer JSONEndpoints and simple http.HandlerFunc endpoints.
type MixedService interface {
	Service

	// route - method - func
	Endpoints() map[string]map[string]http.HandlerFunc

	// Ensure that the route syntax is compatible with the router
	// implementation chosen in cfg.RouterType.
	// route - method - func
	JSONEndpoints() map[string]map[string]JSONEndpoint
	JSONMiddleware(JSONEndpoint) JSONEndpoint
}

// RPCService is an interface defining an grpc-compatible service that
// also offers JSONContextEndpoints and ContextHandlerFuncs.
type RPCService interface {
	ContextService

	Service() (*grpc.ServiceDesc, interface{})

	// Ensure that the route syntax is compatible with the router
	// implementation chosen in cfg.RouterType.
	// route - method - func
	JSONEndpoints() map[string]map[string]JSONContextEndpoint
	JSONMiddleware(JSONContextEndpoint) JSONContextEndpoint
}

// JSONEndpoint is the JSONService equivalent to SimpleService's http.HandlerFunc.
type JSONEndpoint func(*http.Request) (int, interface{}, error)

// ContextService is an interface defining a service that
// is made up of ContextHandlerFuncs.
type ContextService interface {
	Service

	// route - method - func
	ContextEndpoints() map[string]map[string]ContextHandlerFunc
	ContextMiddleware(ContextHandler) ContextHandler
}

// MixedContextService is an interface defining a service that
// is made up of JSONContextEndpoints and ContextHandlerFuncs.
type MixedContextService interface {
	ContextService

	// route - method - func
	JSONEndpoints() map[string]map[string]JSONContextEndpoint
	JSONContextMiddleware(JSONContextEndpoint) JSONContextEndpoint
}
