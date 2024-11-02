package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379: ", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)

	redis := NewRedis()

	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			return
		}

		println(string(buf[:n]))

		arr, _, err := decode(buf)
		if err != nil {
			conn.Write([]byte(fmt.Errorf("error decode command: %w", err).Error()))
			continue
		}

		switch arr[0] {
		case "PING":
			_, _ = conn.Write(encodeString("PONG"))
		case "ECHO":
			_, _ = conn.Write(encodeString(arr[1]))
		case "GET":
			key := arr[1]

			value, ok := redis.Get(key)
			if !ok {
				conn.Write(encodeString("key not found"))
			} else {
				conn.Write(encodeString(value.(string)))
			}
		case "SET":
			key := arr[1]
			value := arr[2]

			redis.Set(key, value)
			conn.Write(encodeString("OK"))
		default:
			_, _ = conn.Write(encodeString("PONG"))
		}
	}
}
