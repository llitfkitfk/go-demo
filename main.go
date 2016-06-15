package main

import (
	"os"
	"os/signal"
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
)

func main() {

	concurrentReq()

	handleSignals()
}

func concurrentReq() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go MakeReq(url, ch)
	}
	for range os.Args[1:] {
		log.Printf(<-ch)
	}

	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func MakeReq(url string, ch chan <-string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}