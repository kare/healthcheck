// healthcheck retrieves given url and exits with error code 0 when server
// responds with 2xx OK status code.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: healthcheck url\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	log.SetPrefix("healthcheck: ")
	flag.Usage = usage
	if len(os.Args) < 2 || os.Args[1] == "" {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()

	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("http get request to '%v' failed with error: %v", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Fatalf("error: unexpected http status code %v for '%v'", res.StatusCode, url)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("i/o error while reading response body from '%v': %v", url, err)
	}
}
