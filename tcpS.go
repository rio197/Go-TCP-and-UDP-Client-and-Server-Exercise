// https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "time"
	"./handleError"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        PORT := ":" + arguments[1]
        
	l, err := net.Listen("tcp", PORT)
	handleError.HandleError(err)
        defer l.Close()

        c, err := l.Accept()
	handleError.HandleError(err) 

        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
         	handleError.HandleError(err) 

               if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                c.Write([]byte(myTime))
        }
}
    
