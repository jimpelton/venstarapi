package main

import (
	"net/http"
	"runtime/debug"

	"github.com/jimpelton/tstat/lib/venstar"
)

func run() {
	body, code, err := venstar.Get("http://192.168.0.23/query/info")
	if err != nil {
		println("Error: ", err.Error())
	}

	println("run(): reply status:", code)

	if code != http.StatusOK {
		return
	}

	println(string(body[:]))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			// run()
		}
	}()
	run()
}
