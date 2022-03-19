package router

import (
	"net/http"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

type Route struct {
	Get    RouteHandler
	Post   RouteHandler
	Put    RouteHandler
	Patch  RouteHandler
	Delete RouteHandler
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func handleRoute(route RouteHandler, writer http.ResponseWriter, request *http.Request) {
	if route != nil {
		route(writer, request)
		return
	}

	handleNotFound(writer, request)
}

func (route *Route) MakeHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			handleRoute(route.Get, writer, request)
		case "POST":
			handleRoute(route.Post, writer, request)
		case "PUT":
			handleRoute(route.Put, writer, request)
		case "PATCH":
			handleRoute(route.Patch, writer, request)
		case "DELETE":
			handleRoute(route.Delete, writer, request)
		default:
			handleNotFound(writer, request)
		}
	}
}
