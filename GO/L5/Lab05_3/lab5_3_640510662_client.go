// ธีรภัทร์ นิลศิริ
// 640510662
// Lab05_3
// 204203 Sec 001
package main
import ("net"
		"fmt"
		"bufio"
		"os"
)
func main(){
	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	// input user name
	fmt.Print("Input your name : ")
	keyboard_buffer_name := bufio.NewReader(os.Stdin)
	strUser, _ := keyboard_buffer_name.ReadString('\n')
	fmt.Print("Welcome ",strUser)
	// connect username to server
	fmt.Fprintf(conn, strUser+"\n")
	
	for {
		

		// create buffer to get keyboard input
		keyboard_buffer := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := keyboard_buffer.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text+"\n")
		// create buffer and recieve from conn
		server_buffer := bufio.NewReader(conn)
		message, _ := server_buffer.ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}