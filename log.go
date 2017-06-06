package singlewolf

import (
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	errorLog *log.Logger
)

// SetErrorLog 用于配置日志输出
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
	ip := getRealIP(r.Header)
	hs := time.Now().Sub(start)
	ret, _ := res["ret"]
	if r.Method == "GET" {
		logf("[%s]get_url:%s(params:%v,time:%fms,ret:%v)", ip, r.URL.Path, r.Params.GetAll(), hs.Seconds()*1000, ret)
	} else {
		logf("[%s]post_url:%s(params:%v,time:%fms,ret:%v)", ip, r.URL.Path, r.Params.GetAll(), hs.Seconds()*1000, ret)
	}

}

func getRealIP(header http.Header) string {
	ip := header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	remote := header.Get("X-Forwarded-For")
	if remote == "" {
		return header.Get("X-Real-IP")
	}
	idx := strings.LastIndex(remote, ",")
	if idx > -1 {
		remote = strings.TrimSpace(remote[idx+1:])
	}
	return remote
}
