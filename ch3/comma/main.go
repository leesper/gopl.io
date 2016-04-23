// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"sort"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", nonrecursiveComma(os.Args[i]))
	}
	fmt.Printf("  %t\n", anagrams("integral", "triangle"))
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
func nonrecursiveComma(s string) string {
	if len(s) <= 3 {
		return s
	}

	secs := []string{}
	for len(s) > 3 {
		secs = append(secs, s[len(s) - 3:])
		s = s[:len(s) - 3]
	}
	secs = append(secs, s)

	var buf bytes.Buffer
	for i := len(secs)-1; i >= 0; i-- {
		buf.WriteString(secs[i])
		if i > 0 {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}

func anagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)
	for i := 0; i < len(str1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
