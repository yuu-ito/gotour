package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		if b[i] >= 'A' && b[i] <= 'L' {
			b[i] = b[i] + 13
		} else if b[i] >= 'M' && b[i] <= 'Z' {
			b[i] = b[i] - 13
		} else if b[i] >= 'a' && b[i] <= 'l' {
			b[i] = b[i] + 13
		} else if b[i] >= 'm' && b[i] <= 'z' {
			b[i] = b[i] - 13
		}
	}
	return n, err
}

func main() {
	str := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(str)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
