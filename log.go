package singlewolf

import (
	"log"
	"strings"
	"time"
)

var (
	errorLog *log.Logger
)

func SetErrorLog(err *log.Logger) {
	errorLog = err
}

func logf(format string, args ...interface{}) {
	if errorLog != nil {
		errorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func writeLog(r *Request, start time.Time, res Result) {
	ip := getRealIp(r)
	hs := time.Now().Sub(start)
	ret, _ := res["ret"]

	if r.Method == "GET" {
		logf("[%s]get_url:%s(params:%v,time:%fms,ret:%v)", ip, r.URL.Path, r.Params.GetAll(), hs.Seconds()*1000, ret)
	} else {
		logf("[%s]post_url:%s(params:%v,time:%fms,ret:%v)", ip, r.URL.Path, r.Params.GetAll(), hs.Seconds()*1000, ret)
	}

}

// getRealIp get real ip.
func getRealIp(r *Request) string {
	ip := r.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	remote := r.Request.Header.Get("X-Forwarded-For")
	if remote == "" {
		return r.Request.Header.Get("X-Real-IP")
	}
	idx := strings.LastIndex(remote, ",")
	if idx > -1 {
		remote = strings.TrimSpace(remote[idx+1:])
	}
	return remote
}
