package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	//fmt.Printf("%q\n", fields)

	m := make(map[string]int)
	for _, sub := range fields {
		m[sub] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
