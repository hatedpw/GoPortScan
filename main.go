package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ip_input() string {
	var ip string
	fmt.Print("Enter an IP address: ")
	fmt.Scanln(&ip)
	return ip
}

//ping a host port from 1 to 3306 if no response after 2 seconds close connection and continue to next port and print the response
func scan_tcp(ip string) {
	for i := 22; i < 3306; i++ {
		conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(i))
		if err != nil {
			continue
		}
		conn.SetDeadline(time.Now().Add(2 * time.Second))
		defer conn.Close()
		fmt.Println("Port", i, "is open")
		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			continue
		}
		fmt.Println(string(response))
		conn.Close()
	}
}

func main() {
	ip := ip_input()
	scan_tcp(ip)
}
