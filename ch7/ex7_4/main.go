// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
    "io"

	"golang.org/x/net/html"
)

type MyReader string

func NewReader(s string) *MyReader {
    var my MyReader
    my = MyReader(s)
    return &my
}

func (my *MyReader) Read(b []byte) (n int, err error) {
    n = len(*my)
    copy(b, []byte(*my))
    err = io.EOF
    return
}


func main() {
	doc, _ := html.Parse(NewReader("<html><body><h1>hello</h1></body></html>"))
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
//		os.Exit(1)
//	}
//	for _, link := range visit(nil, doc) {
//		fmt.Println(link)
//	}
    fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
//func visit(links []string, n *html.Node) []string {
//	if n.Type == html.ElementNode && n.Data == "a" {
//		for _, a := range n.Attr {
//			if a.Key == "href" {
//				links = append(links, a.Val)
//			}
//		}
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		links = visit(links, c)
//	}
//	return links
//}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
