package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
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

		if arr[0] == "PING" {
			_, _ = conn.Write(encodeString("PONG"))
		} else if arr[0] == "ECHO" {
			_, _ = conn.Write(encodeString(arr[1]))
		}
	}
}

func decode(buf []byte) ([]string, []byte, error) {
	s := string(buf[0])
	buf = buf[1:]

	if s == "*" {
		return array(buf)
	}
	return nil, nil, errors.New("Unknown command: " + s)
}

func array(buf []byte) ([]string, []byte, error) {
	s := string(buf)
	println(s)

	arr := make([]string, 0)

	n, err := strconv.Atoi(string(buf[0]))
	if err != nil {
		return nil, nil, err
	}
	buf = buf[3:]
	println(arr, n)

	for i := 0; i < n; i++ {
		s := string(buf)
		println(s)

		ch := string(buf[0])
		if ch == "$" {
			s, b, err := str(buf)
			buf = b

			if err != nil {
				return nil, nil, err
			}
			arr = append(arr, s)
		}
	}

	return arr, buf, nil
}

func str(buf []byte) (string, []byte, error) {
	buf = buf[1:]

	n, err := strconv.Atoi(string(buf[0]))
	if err != nil {
		return "", nil, err
	}

	return string(buf[3 : n+3]), buf[n+5:], nil
}

func encodeString(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}
