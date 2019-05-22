package main

import (
    "fmt"
    "os"
    "io"
)

type limitreader struct {
    io io.Reader
    n int64
}

func (limit *limitreader) Read(p []byte) (n int, err error) {
    if (limit.n <= 0) {
        return 0, io.EOF
    }
    if (int64(len(p)) > limit.n) {
        p = p[0:limit.n]
    }
    n, err = limit.io.Read(p)
    limit.n -= int64(n)
    return
}



func LimitReader(r io.Reader, n int64) io.Reader {
    return &limitreader{r, n}
}

func main() {
    var p = make([]byte, 1024)
    reader := LimitReader(os.Stdin, 16)
    n, _:= reader.Read(p)
    p = p[:n]
    fmt.Println(string(p), n)
    return
}


