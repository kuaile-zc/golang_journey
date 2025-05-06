package goroutines

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func Chat() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleChatConn(conn)
	}

}

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// broadcast message to all clients
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleChatConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	userName := conn.RemoteAddr().String()
	ch <- "You are " + userName
	messages <- userName + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- userName + ": " + input.Text()
	}

	leaving <- ch
	messages <- userName + " has left"
	err := conn.Close()
	if err != nil {
		return
	}

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			return
		}
	}
}
