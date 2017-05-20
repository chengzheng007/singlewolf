package singlewolf

import (
	"net/http"
	"time"
)

// http.Handler
func MakeHandler(routes ...*route) (*router, error) {
	rtr := &router{RT: map[string]*route{}}
	for _, r := range routes {
		// uri must begin with /
		if r.pattern == "" || r.pattern[0:1] != "/" {
			return nil, ErrUriInvalid
		}
		// uri cannot be repeat
		if _, ok := rtr.RT[r.pattern]; ok {
			return nil, ErrUriRepeat
		}

		rtr.RT[r.pattern] = r
	}

	return rtr, nil
}

func (rtr *router) ServeHTTP(wr http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	if handler := rtr.matchPattern(path); handler != nil {
		start := time.Now()

		// 执行函数
		wp := &Wrapper{
			Request{
				r,
				getRequestParams(r),
			},
			&responseWriter{
				wr,
				false,
			},
		}
		var res Result = make(map[string]interface{})
		handler(wp, res)

		// 回写结果
		if err := wp.ResponseWriter.WriteJson(res); err != nil {
			logf("wp.ResponseWriter.WriteJson(%v) error(%v)", err)
		}

		// 记录日志
		writeLog(&wp.Request, start, res)
		return
	}

	notFoundHandle(wr)
	return
}

func (rtr *router) matchPattern(pattern string) HandlerFunc {
	for k, v := range rtr.RT {
		if k == pattern {
			return v.handler
		}
	}

	return nil
}
