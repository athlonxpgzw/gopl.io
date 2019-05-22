// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
    "math/rand"
    "fmt"
    "strconv"
    "strings"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string{
//    var s string
    sli := make([]string, 50)
//    if t != nil {
//        t.left.String()
//        t.right.String()
//    }
//    return strconv.Itoa(t.value)
    var visit func(t *tree)
    visit = func(t *tree) {
        if t.left != nil {
            visit(t.left)
        }
//        fmt.Sprintf(s, strconv.Itoa(t.value))
//        s = fmt.Sprintf("%s %d", s, t.value)
        sli = append(sli, strconv.Itoa(t.value))
        if t.right != nil {
            visit(t.right)
        }
    }
    visit(t)
    return strings.Join(sli, " ")
}


//// Sort sorts values in place.
//func Sort(values []int) {
//	var root *tree
//	for _, v := range values {
//		root = add(root, v)
//	}
//	appendValues(values[:0], root)
//}
//
//// appendValues appends the elements of t to values in order
//// and returns the resulting slice.
//func appendValues(values []int, t *tree) []int {
//	if t != nil {
//		values = appendValues(values, t.left)
//		values = append(values, t.value)
//		values = appendValues(values, t.right)
//	}
//	return values
//}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
    var tree = new(tree)
    for i := range data {
        add(tree, data[i])
    }
    fmt.Println(tree)
//	treesort.Sort(data)
//	if !sort.IntsAreSorted(data) {
//		t.Errorf("not sorted: %v", data)
//	}
}
//!-
