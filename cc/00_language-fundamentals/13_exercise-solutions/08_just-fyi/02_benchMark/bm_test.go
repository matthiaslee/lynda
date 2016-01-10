package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

// run this at command:
// go test -bench='.*'

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
