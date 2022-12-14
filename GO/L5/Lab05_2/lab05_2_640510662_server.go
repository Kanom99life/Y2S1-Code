// ธีรภัทร์ นิลศิริ
// 640510662
// Lab05_2
// 204203 Sec 001
package main
import ("net"
		"fmt"
		"bufio"
		"os"
)
func main() {
	fmt.Println("Server is ready.")
	ln, _ := net.Listen("tcp", ":8000")
	for{
		conn, _ := ln.Accept()
		fmt.Println("Accepted connection.")
		go func(){ // Do in another thread (Create new Goroutine)
			for {
				client_buffer := bufio.NewReader(conn)
				message, _ := client_buffer.ReadString('\n')
				fmt.Print("Message Received: ", string(message))
				// send to client
				keyboard_buffer := bufio.NewReader(os.Stdin)
				fmt.Print("Text to send: ")
				text, _ := keyboard_buffer.ReadString('\n')
				fmt.Fprintf(conn, text + "\n")
			}
		}()
	}
}