package singlewolf

import (
	"errors"
	"net/http"
)

var ErrUriInvalid = errors.New("uri pattern error")
var ErrUriRepeat = errors.New("uri is repeat")

// function type to process your business
type HandlerFunc func(*Wrapper, Result)

// route storage uri-handler pair
type route struct {
	pattern string //uri to access
	handler HandlerFunc
}

func NewRoute(pattern string, handler HandlerFunc) *route {
	return &route{pattern, handler}
}

// router implement http.ServeHTTP, it get corresponding handler and process client request
type router struct {
	RT map[string]*route
}

// paramsData storage paramsters data client sent, and it ha been Json Unmarshaled to map
type paramsData struct {
	data map[string]interface{}
}

type Request struct {
	*http.Request
	Params paramsData
}

// write back to front end
type Result map[string]interface{}

// Wrapper wrap request and response,
type Wrapper struct {
	Request
	ResponseWriter
}
