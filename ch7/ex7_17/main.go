// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
//			stack = append(stack, tok.Name.Local) // push
            attributeList := make([]string, 0)
            attributeList = append(attributeList, tok.Name.Local)
            for _, attr := range tok.Attr {
                if (attr.Name.Local == "id" || attr.Name.Local == "class") {
                    attributeList = append(attributeList, attr.Value)
                }
            }
			stack = append(stack, attributeList) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
                for _, nameAndAttributes := range stack {
                    fmt.Printf("%s ", strings.Join(nameAndAttributes, "|"))
                }
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
//	for len(y) <= len(x) {
//		if len(y) == 0 {
//			return true
//		}
//		if x[0] == y[0] {
//			y = y[1:]
//		}
//		x = x[1:]
//	}
    for len(y) <= len(x) {
        if len(y) == 0 {
            return true
        }
        for _, attr := range x[0] {
            if attr == y[0] {
                y = y[1:]
                break
            }
        }
        x = x[1:]
    }
	return false
}

//!-
