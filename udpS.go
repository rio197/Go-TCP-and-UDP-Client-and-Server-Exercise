// https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
package main

import (
        "fmt"
        "math/rand"
        "net"
        "os"
        "strconv"
        "strings"
        "time"
)

func random(min, max int) int {
        return rand.Intn(max-min) + min
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a port number!")
                return
        }
        PORT := ":" + arguments[1]

        s, err := net.ResolveUDPAddr("udp4", PORT)
	handleError.HandleError(err)

        connection, err := net.ListenUDP("udp4", s)
	handleError.HandleError(err)

        defer connection.Close()
        buffer := make([]byte, 1024)
        rand.Seed(time.Now().Unix())

        for {
                n, addr, err := connection.ReadFromUDP(buffer)
		handleError.HandleError(err)

                fmt.Print("-> ", string(buffer[0:n-1]))

                if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
                        fmt.Println("Exiting UDP server!")
                        return
                }

                data := []byte(strconv.Itoa(random(1, 1001)))
                fmt.Printf("data: %s\n", string(data))
                _, err = connection.WriteToUDP(data, addr)
		handleError.HandleError(err)
       }
}
