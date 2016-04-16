// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s := ""
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
	fmt.Println(s)
	fmt.Printf("runs %f seconds\n", time.Since(start).Seconds())
}

//!-
