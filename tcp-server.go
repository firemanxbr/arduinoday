package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "log"

func handleConnection(conn net.Conn){
   for {
      // will listen for message to process ending in newline (\n)
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil{ break }
      // output message received
      fmt.Print("Message Received:", string(message))

      // sample process for string received
      newmessage := strings.ToUpper(message)

      // send new string back to client
      conn.Write([]byte(newmessage + "\n"))
   }
}

func main() {

   fmt.Println("Launching server...")

   // listen on all interfaces
   ln, _ := net.Listen("tcp", ":8081")

   // run loop forever (or until ctrl-c)
   for {
      // accept connection on port
      conn, err := ln.Accept()
      if err != nil{ log.Fatal(err); break }
      go handleConnection(conn)
   }
}
