// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Fetch wyświetla zawartość znalezioną pod adresem URL.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// go run main.go https://cs.lipsum.com/
// go run main.go http://badgolang.pl

func main() {
	// ioutil_ReadAll()

	// io_copy()

	http_prefix()

	resp_status()
}

func ioutil_ReadAll() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		log.Printf("resp: %v\n", resp)
		fmt.Println("----------------------")

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // aby nie wyciekały zasoby
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("----------------------")
		fmt.Printf("Bytes: %s\n", b)
	}
}

func io_copy() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		log.Printf("resp: %v\n", resp)
		fmt.Println("----------------------")

		// Defining source
		src := resp.Body

		// Defining destination using Stdout
		dst := os.Stdout
		b, err := io.Copy(dst, src)
		resp.Body.Close() // aby nie wyciekały zasoby
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("----------------------")
		fmt.Printf("Bytes: %v\n", b)
	}
}

func http_prefix() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		log.Printf("url: %v\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		log.Printf("resp: %v\n", resp)
		fmt.Println("----------------------")

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // aby nie wyciekały zasoby
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("----------------------")
		fmt.Printf("Bytes: %s\n", b)
	}
}

func resp_status() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		log.Printf("url: %v\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		log.Printf("resp: %v\n", resp)
		fmt.Println("----------------------")

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // aby nie wyciekały zasoby
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("----------------------")
		fmt.Printf("Bytes: %d\n", len(b))
		fmt.Printf("resp.Status: %s\n", resp.Status)
	}
}

//!-
