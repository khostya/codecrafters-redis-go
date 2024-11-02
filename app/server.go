package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/decode"
	"github.com/codecrafters-io/redis-starter-go/app/encode"
	redis2 "github.com/codecrafters-io/redis-starter-go/app/redis"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
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

	redis := redis2.NewRedis()

	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			return
		}

		arr, _, err := decode.Decode(buf)
		if err != nil {
			conn.Write([]byte(fmt.Errorf("error decode command: %w", err).Error()))
			continue
		}

		switch arr[0] {
		case "PING":
			_, _ = conn.Write(encode.String("PONG"))
		case "ECHO":
			_, _ = conn.Write(encode.String(arr[1]))
		case "GET":
			key := arr[1]

			value, ok := redis.Get(key)
			if !ok {
				conn.Write(encode.Null())
			} else {
				conn.Write(encode.String(value.(string)))
			}
		case "SET":
			key := arr[1]
			value := arr[2]

			var dur = time.Hour * 24 * 365
			if len(arr) > 3 && strings.ToLower(arr[3]) == "px" {
				ml, err := strconv.Atoi(arr[4])
				if err != nil {
					conn.Write([]byte(err.Error()))
				}
				dur = time.Millisecond * time.Duration(ml)
			}

			redis.Set(key, value, dur)
			conn.Write(encode.String("OK"))
		default:
			_, _ = conn.Write(encode.String("PONG"))
		}
	}
}
