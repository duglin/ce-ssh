package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	port := "" // [interface:]port
	if tmp := os.Getenv("PORT"); tmp != "" {
		port = tmp
	}
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	if port == "" {
		log.Printf("Missing port to listen on (arg or PORT env)")
		os.Exit(1)
	}
	if strings.Index(port, ":") < 0 {
		port = "0.0.0.0:" + port
	}

	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Listening on: %s\n", port)
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n]))

		data := []byte(fmt.Sprintf("yes I got:%s", buffer[0:n-1]))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
