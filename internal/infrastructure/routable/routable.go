package routable

import "net/http"

type Routable interface {
	InitRoutes() http.Handler
}
