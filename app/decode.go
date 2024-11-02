package main

import (
	"errors"
	"strconv"
)

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
