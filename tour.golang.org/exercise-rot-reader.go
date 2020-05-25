package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)

	for i := range b {
		if b[i] < 65+13 {
			b[i] += 13
		} else if b[i] < 65+26 {
			b[i] -= 13
		} else if b[i] < 96+13 {
			b[i] += 13
		} else if b[i] < 96+26 {
			b[i] -= 13
		}
	}
	fmt.Println(">>>", "length is ", n)

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
