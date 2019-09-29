// https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/#net-package-functions
package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strconv"
        "strings"
)

var count = 0

func handleConnection(c net.Conn) {
        fmt.Print(".")
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
		handleError.HandleError(err)

                temp := strings.TrimSpace(string(netData))
                if temp == "STOP" {
                        break
                }
                fmt.Println(temp)
                counter := strconv.Itoa(count) + "\n"
                c.Write([]byte(string(counter)))
        }
        c.Close()
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a port number!")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp4", PORT)
	handleError.HandleError(err)
	defer l.Close()

        for {
                c, err := l.Accept()
		handleError.HandleError(err) 

               	go handleConnection(c)
                count++
        }
}
