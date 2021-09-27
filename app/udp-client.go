package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	connect := "" // [host:]port
	if tmp := os.Getenv("SERVER"); tmp != "" {
		connect = tmp
	}
	if len(os.Args) == 2 {
		connect = os.Args[1]
	}
	if connect == "" {
		log.Printf("Missing server to connect to (arg or SERVER env var)")
		os.Exit(1)
	}
	if strings.Index(connect, ":") < 0 {
		connect = "localhost:" + connect
	}

	s, err := net.ResolveUDPAddr("udp4", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Connecting to: %s", s)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("The UDP server is %q\n", c.RemoteAddr().String())
	defer c.Close()

	buffer := make([]byte, 1024)
	for i := 0; i < 10; i++ {
		_, err := c.Write([]byte(fmt.Sprintf("From client message #%d\n", i)))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
	time.Sleep(5 * time.Second)
}
