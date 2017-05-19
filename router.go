package singlewolf

import (
	"errors"
	"net/http"
)

type router struct {
	RT map[string]*route
}

var ErrUriInvalid = errors.New("uri pattern error")

func (r *router) MuxFunc() HandlerFunc {
	return func(wp *Wrapper, res Result) {
		// TODO 根据请求URI匹配具体handler
		// 遍历Router.RT

		path := wp.Request.URL.Path
		if handler := r.matchPatern(path); handler != nil {
			handler(wp, res)
			return
		}

		notFoundHandle(wp.ResponseWriter, res, http.StatusNotFound)
		return
	}
}

func MakeRouter(routes ...*route) (MuxI, error) {

	rtr := &router{RT: map[string]*route{}}
	// 如果有相同的url后边的覆盖前面的
	for _, r := range routes {
		// uri不能为空，且必须以/开头
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
