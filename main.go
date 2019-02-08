package main

import (
	"flag"
	"fmt"
	"os"
	"net"
	"log"
	"./lib/lwm2m"
)


type (
	Opts struct {
		LWM2MHost string
		LWM2MPort string
		Help bool
	}
)


const (
	help string = `
	You need help ... why?
	`
)

func initialize() (*net.UDPConn,error) {
	flag.StringVar(&opts.LWM2MHost, "a", "localhost", "LwM2M host.")
	flag.StringVar(&opts.LWM2MPort, "p", "5683", "LwM2M port.")
	flag.BoolVar(&opts.Help, "h", false, "Show help.")
	flag.BoolVar(&opts.Help, "help", false, "Show help.")
	flag.Parse()
	if opts.Help {
		fmt.Printf("%s\n",help)
		os.Exit(0)
	}
	addr := fmt.Sprintf("%s:%s", opts.LWM2MHost, opts.LWM2MPort)
	uaddr, err := net.ResolveUDPAddr("udp",addr)
	l, err := net.ListenUDP("udp", uaddr)
	return l,err
}


var opts Opts


func main() {
	l,err := initialize()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server started")
		lwm2m.Start(l)
	}
}
