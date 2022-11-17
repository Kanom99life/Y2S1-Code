// ธีรภัทร์ นิลศิริ
// 640510662
// Lab05_3
// 204203 Sec 001
package main
import ("net"
		"fmt"
		"bufio"
		"os"
		"strings"
)
func main() {
	fmt.Println("Server is ready.")
	ln, _ := net.Listen("tcp", ":8000") // message
	
	for{
		conn, _ := ln.Accept() //Accept message
		fmt.Println("Accepted connection.")
		
		client_name := bufio.NewReader(conn)
		user_name, _:= client_name.ReadString('\n')
		user_name = string(user_name)
		user_name = strings.Replace(user_name, "\n", "", 1)
		go func(){ // Do in another thread (Create new Goroutine)
			for {
				client_buffer := bufio.NewReader(conn)
				message, _ := client_buffer.ReadString('\n')
				fmt.Print(string(user_name)," said: ",string(message))
				// send to client
				keyboard_buffer := bufio.NewReader(os.Stdin)
				fmt.Print("reply to ",string(user_name),": ")
				text, _ := keyboard_buffer.ReadString('\n')
				fmt.Fprintf(conn, text + "\n")
			}
		}()
	}
}