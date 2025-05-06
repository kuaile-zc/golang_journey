package goroutines

import (
	"io"
	"log"
	"net"
	"time"
)

// ReturnListenTime 顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端
// 测试使用 nc localhost 8080
func ReturnListenTime() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

	}
}
