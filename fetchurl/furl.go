package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func fetchurl(url string, ch chan<- string) {
	//	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	mybuf, _ := ioutil.ReadAll(resp.Body)

	//bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	//secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%s", string(mybuf))

}

func main() {
	runtime.GOMAXPROCS(20)

	urls := os.Args[1:]
	//start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetchurl(url, ch)
		//fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Println("")
	for _, url := range urls {
		_ = url
		fmt.Println(<-ch) // receive from channel ch
		//fmt.Fprintln(<-ch)
	}

	//secs := time.Since(start).Seconds()

	//fmt.Printf("%.2fs total elapsed\n", secs)

	//body, err := ioutil.ReadAll(resp.Body)

	//header := resp.Header

	//clen := resp.ContentLength

	//fmt.Println(resp.ContentLength)

	//fmt.Printf("%v", resp.Status)
	//fmt.Printf(" %v\n", resp.Proto)
	//fmt.Printf("%v\n", clen)

	// fmt.Printf("Connection: %v\n", resp.Close)

	// fmt.Printf("Content Length: %v\n", resp.ContentLength)

	// fmt.Printf("Date: %v\n", header.Get("Date"))
	// fmt.Println(header.Get("Proto"))
	//fmt.Println(string(body))

	//io.Copy(os.Stdout, resp.Body)

	// _ = header.

	//}
}
