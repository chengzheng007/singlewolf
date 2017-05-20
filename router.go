package singlewolf

import (
	"errors"
	"net/http"
)

type router struct {
	RT map[string]*route
}

var ErrUriInvalid = errors.New("uri pattern error")

func (r *router) getHandlerFunc() HandlerFunc {
	return func(wp *Wrapper, res Result) {
		path := wp.Request.URL.Path
		if handler := r.matchPatern(path); handler != nil {
			handler(wp, res)
			return
		}
		notFoundHandle(wp.ResponseWriter, res, http.StatusNotFound)
		return
	}
}

func MakeRouter(routes ...*route) (*router, error) {

	rtr := &router{RT: map[string]*route{}}
	// if set same uri pattern, the latter will cover ahead
	for _, r := range routes {
		// uri must begin with /
		if r.pattern == "" || r.pattern[0:1] != "/" {
			return rtr, ErrUriInvalid
		}
		rtr.RT[r.pattern] = r
	}

	return rtr, nil
}

func (r *router) matchPatern(pattern string) HandlerFunc {
	if r.RT == nil {
		return nil
	}

	for k, v := range r.RT {
		if k == v.pattern {
			return v.handler
		}
	}

	return nil
}
