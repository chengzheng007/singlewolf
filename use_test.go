package singlewolf

import (
	"fmt"
	"testing"
	"time"
)

var (
	addrs       = []string{"127.0.0.1:9088"}
	httpTimeout = time.Second * 15
)

func TestA(t *testing.T) {

	// example

	// you can set logger to print it to your disk, default it print to screen
	// SetErrorLog(err)

	// make handler
	handler, err := MakeHandler(
		NewRoute("/hello", hello),
	)
	if err != nil {
		logf("MakeHandler error(%v)", err)
		return
	}

	logf("%s\n", "start ...")

	// start serve request
	if err := StartServe(handler, addrs, httpTimeout); err != nil {
		logf("Serve(handler, %v, %v) error(%v)", addrs, httpTimeout, err)
		return
	}
	// close serve
	defer Close()

	for {
	}

}

// specific function to process client request
func hello(w *Wrapper, res Result) {
	// test request data，you can use it to send a post request and verify result
	// {"a":"hello","b":"abc","c":true,"d":12.34567888888,"e":1234567890, "array":[{"name":"Jack", "age":23}], "map":{"corp":"google", "address":"us"}}

	res["ret"] = 1

	res["params_of_all"] = w.Request.Params.GetAll()

	res["string"] = w.Request.Params.GetString("a")
	res["bytes"] = w.Request.Params.GetBytes("b")
	res["bool"] = w.Request.Params.GetBool("c")
	res["float64"] = w.Request.Params.GetFloat64("d")
	res["int64"] = w.Request.Params.GetInt64("e")
	// single object node
	res["map"] = w.Request.Params.GetArray("map")
	// array of object node
	res["array"] = w.Request.Params.GetArray("array")

	// if you want get element in array, you can iterate it,
	// call type-associated function to get its value, for example v.GetString("key")
	// arr := w.Request.Params.GetArray("array")
	// for _, v := range arr {
	// fmt.Println(v.GetString("key"))
	// }

	res["not_exist"] = w.Request.Params.GetString("not_exist")

}
