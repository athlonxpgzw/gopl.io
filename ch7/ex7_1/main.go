// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
    "bufio"
    "strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(strings.NewReader(string(p)))
    scanner.Split(bufio.ScanWords)
    c := 0
    for scanner.Scan() {
        c++
    }
    *wc += WordCounter(c)
    return c, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(strings.NewReader(string(p)))
    scanner.Split(bufio.ScanLines)
    c := 0
    for scanner.Scan() {
        c++
    }
    *lc += LineCounter(c)
    return c, nil
}


//!-bytecounter

func main() {
	//!+main
	var wc WordCounter
	wc.Write([]byte("hello"))
	fmt.Println(wc) // "5", = len("hello")

	wc = 0 // reset the counter
	var name = "Dolly, marry had an lamb"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc) // "12", = len("hello, Dolly")
	//!-main
    var lc LineCounter
    var lines = []byte(`hello \nmy\nname\nis\nGGG
    fqfeq
    feqfeq
    feqfeq
    431h
    43134
    fejqkwfeqh
    j43k14hj31
    furhbn43
    feqfeq
    feqfeq`)
    lc.Write(lines)
    fmt.Println(lc)
    wc.Write(lines)
    fmt.Println(wc)
}
