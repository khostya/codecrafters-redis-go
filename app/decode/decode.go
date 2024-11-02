package decode

import (
	"errors"
	"strconv"
	"unicode"
)

func Decode(buf []byte) ([]string, []byte, error) {
	s := string(buf[0])
	buf = buf[1:]

	if s == "*" {
		length, buf, err := scanInt(buf)
		if err != nil {
			return nil, nil, err
		}

		return array(length, buf)
	}
	return nil, nil, errors.New("Unknown command: " + s)
}

func array(length int, buf []byte) ([]string, []byte, error) {
	arr := make([]string, 0)

	for i := 0; i < length; i++ {
		log(buf)

		ch := string(buf[0])
		if ch == "$" {
			s, b, err := str(buf[1:])
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
	n, buf, err := scanInt(buf)
	if err != nil {
		return "", nil, err
	}

	log(buf)
	return string(buf[:n]), buf[n+2:], nil
}

func scanInt(buf []byte) (int, []byte, error) {
	b := make([]byte, 0)

	for i := 0; i < len(buf); i++ {
		if unicode.IsDigit(rune(buf[i])) {
			b = append(b, buf[i])
			continue
		}
		break
	}

	log(buf)

	n, err := strconv.Atoi(string(b))
	return n, buf[2+len(b):], err
}

func log(buf []byte) {

}
