package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func main() {
	parallel := flag.Int("parallel", 10, "set number of parallel requests")
	flag.Parse()
	urls := flag.Args()
	doRequests(urls, *parallel, os.Stdout)
}

func doRequests(urls []string, parallel int, w io.Writer) {
	wg := sync.WaitGroup{}
	total := len(urls)
	for i := 0 ; i < total; i += parallel {
		max := total - i
		if parallel < max {
			max = parallel
		}
		urlSlice := urls[i:i + max]
		for _, url := range urlSlice {
			wg.Add(1)
			go doRequest(&wg, url, w)
		}
		wg.Wait()	
	}
}

func doRequest(wg *sync.WaitGroup, urlStr string, w io.Writer) {

	defer wg.Done()

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		// Errors behavior not specified
		return
	}
	if parsedURL.Scheme == "" {
		// Looks a requirement to allow no protocol, try to default http://
		parsedURL.Scheme = "http"
	}
	finalURL := parsedURL.String()

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		// Errors behavior not specified
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Errors behavior not specified
		return
	}
	defer res.Body.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, res.Body); err != nil {
		log.Fatal(err)
		// Errors behavior not specified
		return
	}
	fmt.Fprintf(w, "%s %x\n", finalURL, hash.Sum(nil))
}
