package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
)

type echoServer struct {
	gnet.BuiltinEventEngine
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	data, _ := c.Next(-1)
	c.Write(data)
	return gnet.None
}

func main() {
	var port int
	var multicore, reuseport bool

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.BoolVar(&reuseport, "reuseport", false, "--reuseport true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Run(echo, fmt.Sprintf("udp://:%d", port), gnet.WithMulticore(multicore), gnet.WithReusePort(reuseport)))
}
