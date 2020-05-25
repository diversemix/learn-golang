package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (n int, err error) {
	n = 1
	err = nil
	b[0] = 65
	return
}

func main() {
	reader.Validate(MyReader{})
}
