// https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a host:port string")
                return
        }
        CONNECT := arguments[1]

        s, err := net.ResolveUDPAddr("udp4", CONNECT)
        c, err := net.DialUDP("udp4", nil, s)
	handleError.HandleError(err)

        fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
        defer c.Close()

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                data := []byte(text + "\n")
                _, err = c.Write(data)
                if strings.TrimSpace(string(data)) == "STOP" {
                        fmt.Println("Exiting UDP client!")
                        return
                }
		handleError.HandleError(err)

                buffer := make([]byte, 1024)
                n, _, err := c.ReadFromUDP(buffer)
		handleError.HandleError(err)

	        fmt.Printf("Reply: %s\n", string(buffer[0:n]))
        }
}
