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

// Result will be written to front end
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

		// execute your request logic
		var res Result = make(map[string]interface{})
		handle(wp, res)

		// write back result to front end
		if err := wp.ResponseWriter.WriteJson(res); err != nil {
			logf("wp.ResponseWriter.WriteJson(%v) error(%v)", err)
		}

		//  write log
		writeLog(&wp.Request, time.Now(), res)
	}
}
