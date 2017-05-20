package singlewolf

import (
	"testing"
	"time"
)

var (
	addrs       = []string{"127.0.0.1:9088"}
	httpTimeout = time.Second * 15
)

func TestA(t *testing.T) {

	// example

	handler, err := MakeHandler(
		NewRoute("/hello", hello),
	)
	if err != nil {
		logf("MakeHandler error(%v)", err)
		return
	}

	logf("%s\n", "start ...")

	if err := StartServe(handler, addrs, httpTimeout); err != nil {
		logf("Serve(handler, %v, %v) error(%v)", addrs, httpTimeout, err)
		return
	}
	defer Close()

	for {
	}

}

func hello(w *Wrapper, res Result) {

	// {"a":"hello","b":"abc","c":true,"d":12.34567888888,"e":1234567890}

	res["ret"] = 1

	res["params"] = w.Request.Params.GetAll()

	res["string"] = w.Request.Params.GetString("a")
	res["bytes"] = w.Request.Params.GetBytes("b")
	res["bool"] = w.Request.Params.GetBool("c")
	res["float64"] = w.Request.Params.GetFloat64("d")
	res["int64"] = w.Request.Params.GetInt64("e")

	res["xxxxx"] = w.Request.Params.GetString("xxxxx")

}
