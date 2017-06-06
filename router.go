package singlewolf

import (
	"net/http"
	"time"
)

// Route storage uri-handler pair
type Route struct {
	pattern string //uri to access
	handler HandlerFunc
}

// NewRoute is
func NewRoute(pattern string, handler HandlerFunc) *Route {
	return &Route{pattern, handler}
}

// Router implement http.ServeHTTP, it get corresponding handler and process client request
type Router map[string]*Route

func (rtr Router) matchPattern(pattern string) HandlerFunc {
	if v, ok := rtr[pattern]; ok {
		return v.handler
	}
	return nil
}

// MakeHandler set http.Handler
func MakeHandler(routes ...*Route) (Router, error) {
	rtr := make(Router)
	for _, r := range routes {
		// uri must begin with /
		if r.pattern == "" || r.pattern[0:1] != "/" {
			return nil, ErrURIInvalid
		}
		// uri cannot be repeat
		if _, ok := rtr[r.pattern]; ok {
			return nil, ErrURIRepeat
		}
		rtr[r.pattern] = r
	}
	return rtr, nil
}

func (rtr Router) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	// 只允许post提交
	if r.Method != "POST" {
		invalidMethodHandle(wr)
		return
	}
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
		if err := wp.ResponseWriter.WriteJSON(res); err != nil {
			logf("wp.ResponseWriter.WriteJSON(%v) error(%v)", res, err)
		}

		// 记录日志
		writeLog(&wp.Request, start, res)
		return
	}

	notFoundHandle(wr)
	return
}
