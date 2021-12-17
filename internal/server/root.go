package server

import (
	"bufio"
	"fmt"
	"go-cloud-calculator/internal/server/ops"
	"net"
	"strings"
)

func RunServer(port string) {

	fmt.Println("-> Running server on port: " + port)

	// lunching a tcp listener on `port`
	PORT := ":" + port
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Printf("-> err: %v\n", err)
		return
	}
	defer l.Close()

	// Accepting new connection
	c, err := l.Accept()
	if err != nil {
		fmt.Printf("-> err: %v\n", err)
		return
	}

	// Contracting with client in a infinite loop
	for {
		// Reading client order
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("-> err: %v\n", err)
			return
		}

		// Checking if the client whants to stop the process
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("-> Exiting TCP server!")
			return
		}

		// Printing what we recieved and analizing it
		input := string(netData)
		// fmt.Print("Client input: ", input)
		output, errH := ops.Handle(input)
		// fmt.Println("-> Result is " + output)
		// Writing answer
		if errH != nil {
			fmt.Printf("-> err: %v\n", errH)
		}
		output = output + "\n"
		// fmt.Println("-> Writing back to client...")
		c.Write([]byte(output))
		// fmt.Println("-> Write back completed.")
	}

}
