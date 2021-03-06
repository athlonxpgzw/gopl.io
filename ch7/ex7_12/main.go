// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
    "strconv"
    "html/template"
)

var listTemp = template.Must(template.New("list").Parse(`
<html>
<body>
{{ range $key, $value := .ItemMap }}
<p>{{$key}}: {{$value}}</p>
{{ end }}
</body>
</html>
`))

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type TemplateData struct {
    ItemMap database
}


func (db database) list(w http.ResponseWriter, req *http.Request) {
//	for item, price := range db {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
    if err := listTemp.Execute(w, &TemplateData{db}); err != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "failed to execute template %q\n", err)
    }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    if  _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
        return
    }

    price := req.URL.Query().Get("price")
    if dollar, err := strconv.ParseFloat(price, 32); err == nil {
        db[item] = dollars(dollar)
        fmt.Fprintf(w, "The price updated to %f\n", dollar)
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "Invalid price, and error is ---", err)
    }
}


