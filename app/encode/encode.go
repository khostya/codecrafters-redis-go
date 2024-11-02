package encode

import (
	"fmt"
	"strings"
)

func String(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}

func ListItem(s string) []byte {
	return []byte(fmt.Sprintf("$%v\r\n%v\r\n", len(s), s))
}

func List(s ...string) []byte {
	var items []string

	for _, ss := range s {
		items = append(items, string(ListItem(ss)))
	}

	return []byte(fmt.Sprintf("*%v\r\n%v", len(items), strings.Join(items, "")))
}

func Null() []byte {
	return []byte(fmt.Sprintf("$-1\r\n"))
}
