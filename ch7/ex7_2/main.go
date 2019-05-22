package main

import (
    "fmt"
    "io"
    "os"
)

type ByteCounter struct {
    io io.Writer
    counter int64
}

func (b *ByteCounter) Write(p []byte) (int, error) {
    n, err := b.io.Write(p)
    b.counter += int64(n)
    return n, err
}


func CountingWriter(w io.Writer) (io.Writer, *int64) {
    b := ByteCounter{w, 0}
    return &b, &b.counter
}


func main() {
    writer, counter := CountingWriter(os.Stdout)
    fmt.Fprintf(writer, "hello, the world\n")
    fmt.Println(*counter)
}

