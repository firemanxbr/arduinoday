package main

import "net"
import "fmt"
import "bufio"
import "os"
import "log"

func main() {

   // connect to this socket
   conn, err := net.Dial("tcp", "127.0.0.1:8081")
   if err != nil{ log.Fatal(err); return }

   for {
      // read in input from stdin
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Text to send: ")
      text, err := reader.ReadString('\n')
      if err != nil{ log.Fatal(err); break }

      // send to socket
      fmt.Fprintf(conn, text+"\n")

      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil{ log.Fatal(err); break }
      fmt.Print("Message from server: " + message)
   }
}
