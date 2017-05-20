package singlewolf

import (
	"net/http"
	"time"
)

// request/response
type Wrapper struct {
	Request
	ResponseWriter
}

// 回写到前端
type Result map[string]interface{}

type HandlerFunc func(*Wrapper, Result)

func wrapHandleFunc(handle HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
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

		// 执行函数
		var res Result = make(map[string]interface{})
		handle(wp, res)

		// 回写结果
		if err := wp.ResponseWriter.WriteJson(res); err != nil {
			logf("wp.ResponseWriter.WriteJson(%v) error(%v)", err)
		}

		// 记录日志
		writeLog(&wp.Request, time.Now(), res)
	}
}
