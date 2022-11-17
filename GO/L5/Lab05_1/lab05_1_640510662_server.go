// ธีรภัทร์ นิลศิริ
// 640510662
// Lab05_1
// 204203 Sec 001
package main
import ("net"
		"fmt"
		"bufio"
		"os"
)
func main() {
	fmt.Println("Start server...")
	// listen on port 800
	ln, _ := net.Listen("tcp", ":8000")
	//accept connection
	conn, _ := ln.Accept()
	// run loop forever (or until ctrl-c)
	for {
	// get message, outpu
		client_buffer := bufio.NewReader(conn)
		message, _ := client_buffer.ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		// send to client
		keyboard_buffer := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := keyboard_buffer.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
	}
}