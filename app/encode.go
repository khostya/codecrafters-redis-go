package main

import "fmt"

func encodeString(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}
