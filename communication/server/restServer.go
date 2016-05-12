package server

import (
	"net/http"
)

type Route struct {
	Url string
	Method string
	Handler HandlerSessionFunc
}

type RestServer struct{
	Root string
	Routes[]*Route
}

func (this *RestServer) HandleRequest(httpResponse http.ResponseWriter, httpRequest *http.Request, session *Session) {
	routePath := httpRequest.URL.Path[len(this.Root):];
	for _,route := range this.Routes{
		if route.Url == routePath && route.Method == httpRequest.Method{
			route.Handler(httpResponse,httpRequest, session)
			break
		}
	}
}

