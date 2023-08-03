package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/jimpelton/tstat/pkg/venstar"
)

func run(host string, port int) {
	cli := venstar.NewVenstarClient(host, port)
	if rep, err := cli.Alerts(); err != nil {
		fmt.Printf("error getting info: %s", err.Error())
	} else {
		fmt.Println("Info: %+v", rep)
	}
}

type Options struct {
	DeviceHost string
	DevicePort uint
}

const venstarDeviceHostEnv = "VENSTAR_DEVICE_HOST"
const venstarDevicePortEnv = "VENSTAR_DEVICE_PORT"

func optionsFromEnv() (opts Options) {
	if devHost, exists := os.LookupEnv(venstarDeviceHostEnv); exists {
		opts.DeviceHost = devHost
	}

	if devPort, exists := os.LookupEnv(venstarDevicePortEnv); exists {
		p, _ := strconv.Atoi(devPort)
		opts.DevicePort = uint(p)
	}

	return
}

func optionsFromCmdLine(opts *Options) error {
	flag.StringVar(&opts.DeviceHost, "host", "", "Host/IP of Venstar device")
	flag.UintVar(&opts.DevicePort, "port", 0, "Port for Venstar device")
	flag.Parse()

	var err error
	if opts.DeviceHost == "" {
		err = errors.Join(
			err,
			errors.New("host option not provided but is required"))
	}

	if opts.DevicePort == 0 {
		err = errors.Join(
			err,
			errors.New("port option not provided but is required"))
	}

	return err
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			// run()
		}
	}()
	opts := Options{}
	if err := optionsFromCmdLine(&opts); err != nil {
		log.Printf("Errors when parsing cmd line:\n%v\n", err.Error())
	}

	run("192.168.0.23", 80)
}
