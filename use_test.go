package singlewolf


import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

var (
	addr        = "127.0.0.1:9088"
	httpTimeout = time.Second * 15
)

func TestA(t *testing.T) {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("net.Listen(\"tcp\", addr) error(%v)\n", err)
		return
	}

	httpServeMux, err := InitServerMux()
	if err != nil {
		fmt.Printf("InitServerMux() error(%v)\n", err)
		return
	}

	server := &http.Server{Handler: httpServeMux, ReadTimeout: httpTimeout, WriteTimeout: httpTimeout}

	fmt.Println("start ...")

	if err := server.Serve(l); err != nil {
		fmt.Printf("server.Serve(l) error(%v)\n", err)
		return
	}
}

func InitServerMux() (http.Handler, error) {
	mux := NewMux()

	router, err := MakeRouter(
		NewRoute("/hello", hello),
	)

	if err != nil {
		return nil, err
	}

	return mux.MakeHandler(router), nil
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

}
