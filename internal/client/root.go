package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// This function runs the client on 'port'
func Execute(IP_PORT string) {

	c, err := net.Dial("tcp", IP_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// Reading user input
		// fmt.Println("-> Reading user input...")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		// fmt.Println("-> Sending completed")

		// Reading server message
		// fmt.Println("-> Reading server ...")
		message, _ := bufio.NewReader(c).ReadString('\n')
		// fmt.Println("-> Reading server completed.")
		fmt.Print("-> " + message)
		// fmt.Println("-> Server message printed")
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("-> TCP client exiting...")
			return
		}
	}

}
