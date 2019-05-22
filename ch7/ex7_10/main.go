package main 

import (
    "sort"
    "fmt"
)

type Palindromechecker []byte

func (p Palindromechecker) Len() int { return len(p) }
func (p Palindromechecker) Less(i, j int) bool { return p[i] < p[j] }
func (p Palindromechecker) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func IsPalindrome(s sort.Interface) bool {
    for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
        if !(!s.Less(i, j) && !s.Less(j, i)) {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(IsPalindrome(Palindromechecker([]byte("abcdefdsfeqfeq"))))
    fmt.Println(IsPalindrome(Palindromechecker([]byte("12345677654321"))))
    fmt.Println(IsPalindrome(Palindromechecker([]byte("abcdefgfedcba"))))
}

