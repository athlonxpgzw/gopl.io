// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
    "strings"
)

type host struct {
    city string
    addr string
    port string
}

func main() {
//	conn, err := net.Dial("tcp", "localhost:8000")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//	mustCopy(os.Stdout, conn)
    hosts := make([]host, 0)
    if len(os.Args) > 1 {
        for i, arg := os.Args {
            if i == 0 {
                continue
            }
            str := strings.Split(arg, "=")
            h := host{city:str[0]}
            str = strings.Split(str, ":")
            h.addr = str[0]
            h.port = str[1]
            hosts = append(hosts, h)
        }
    } else {
        hosts = append(hosts, host{"local", "localhost", "8000"})
    }
    for _, h := range hosts {
        fmt.Fprintf(os.Stdout, "%s\t", h.city)
    }


}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
