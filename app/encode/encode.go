package encode

import "fmt"

func String(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}

func Null() []byte {
	return []byte(fmt.Sprintf("$-1\r\n"))
}
