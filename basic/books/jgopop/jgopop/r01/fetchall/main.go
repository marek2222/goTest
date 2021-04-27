// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Fetchall pobiera równolegle zawartości kilku adresów URL i raportuje czasy pobierania
// oraz rozmiary odpowiedzi.
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
	list := []string{"https://lipsum.com/",
		"https://sq.lipsum.com/",
		"https://www.youtube.com/watch?v=C8y6B73Uk4w",
		"https://docs.microsoft.com/en-us/aspnet/core/blazor/file-uploads?view=aspnetcore-5.0"}
	go_fetch(list)

	// we save only 2 first item
	// we check how to doing bufor this site
	list = list[:2]
	go_fetch(list)

	fileName := "plik.txt"
	go_fetch_to_file(list, fileName)
}

func go_fetch(list []string) {
	start := time.Now()
	ch := make(chan string)
	fmt.Println()
	// i don't want use command line args
	//for _, url := range os.Args[1:] {
	for _, url := range list {
		go fetch(url, ch) // rozpoczęcie funkcji goroutine
	}
	for range list {
		fmt.Println(<-ch) // odbieranie z kanału ch
	}
	fmt.Printf("%.2fs upłynęło\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // wysyłanie do kanału ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // aby nie wyciekały zasoby
	if err != nil {
		ch <- fmt.Sprintf("podczas odczytywania %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func go_fetch_to_file(list []string, fileName string) {
	start := time.Now()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer file.Close()

	ch := make(chan string)
	fmt.Println()
	fmt.Println("Save to file: " + file.Name())
	// i don't want use command line args
	//for _, url := range os.Args[1:] {
	for _, url := range list {
		go fetch_to_file(url, file, ch) // rozpoczęcie funkcji goroutine
	}
	for range list {
		fmt.Println(<-ch) // odbieranie z kanału ch
	}
	fmt.Printf("%.2fs upłynęło\n", time.Since(start).Seconds())
}

func fetch_to_file(url string, file *os.File, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // wysyłanie do kanału ch
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // aby nie wyciekały zasoby
	if err != nil {
		ch <- fmt.Sprintf("podczas odczytywania %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
