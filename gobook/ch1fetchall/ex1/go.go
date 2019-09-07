/*


Find a website that produces a large amount of data.

Investigate caching by running fetchall twice in a row.
  Do the times differ much?
  Do you get the same content?

Modify fetchall to output to a file

*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	outfile, err := os.Create("./outfile.txt")
	defer outfile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintln(outfile, <-ch)
	}

	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(outfile, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) ///throw away the content but count the size in bytes
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url)
}
