package main

import (
	"fmt"
	"runtime/debug"

	"github.com/jimpelton/tstat/pkg/venstar"
)

func run(host string, port int) {
	cli := venstar.NewVenstarClient(host, port)
	if rep, err := cli.Info(); err != nil {
		fmt.Printf("error getting info: %s", err.Error())
	} else {
		fmt.Println("Info: %+v", rep)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			// run()
		}
	}()
	run("192.168.0.22", 80)
}
