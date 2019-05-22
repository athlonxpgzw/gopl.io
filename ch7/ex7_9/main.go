// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"sort"
    "net/http"
    "log"
    "html/template"
	"time"
    "io"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style media="screen" type="text/css">
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=Title">Title</a></th>
					<th><a href="/?sort=Artist">Artist</a></th>
					<th><a href="/?sort=Album">Album</a></th>
					<th><a href="/?sort=Year">Year</a></th>
					<th><a href="/?sort=Length">Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

//!+printTracks
func printTracks(writer io.Writer, tracks []*Track) {
    if err := trackTable.Execute(writer, tracks); err != nil {
        log.Fatal(err)
    }
}

//!-printTracks

//!+StateSort
type StateSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x StateSort) Len() int           { return len(x.t) }
func (x StateSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x StateSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
//!-StateSort

func handler(r http.ResponseWriter, req *http.Request) {
    sortBy := req.FormValue("sort")
    sort.Sort(StateSort{tracks, func(x, y *Track) bool {
        switch sortBy {
        case "Title":
            return x.Title < y.Title
        case "Artist":
            return x.Artist < y.Artist
        case "Year":
            return x.Year < y.Year
        case "Album":
            return x.Album < y.Album
        case "Length":
            return x.Length < y.Length
        }
        return false
    }})
    printTracks(r, tracks)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}


/*
//!+artistoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Go          Delilah         From the Roots Up  2012  3m38s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Moby            Moby               1992  3m37s
//!-artistoutput

//!+artistrevoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
//!-artistrevoutput

//!+yearoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
//!-yearoutput

//!+customout
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//!-customout
*/

////!+customcode
//type customSort struct {
//	t    []*Track
//	less func(x, y *Track) bool
//}
//
//func (x customSort) Len() int           { return len(x.t) }
//func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
//func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
//
//
//
////!-customcode

//func init() {
//	//!+ints
//	values := []int{3, 1, 4, 1}
//	fmt.Println(sort.IntsAreSorted(values)) // "false"
//	sort.Ints(values)
//	fmt.Println(values)                     // "[1 1 3 4]"
//	fmt.Println(sort.IntsAreSorted(values)) // "true"
//	sort.Sort(sort.Reverse(sort.IntSlice(values)))
//	fmt.Println(values)                     // "[4 3 1 1]"
//	fmt.Println(sort.IntsAreSorted(values)) // "false"
//	//!-ints
//}
