package main

import (
    "net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`<html>
    <head><title>Prosty serwer</title></head>
    <body>
        <h1>Witaj!</h1>
        <p> To prymitywna strona startowa bardzo prostego serwera HTTP<p>
        <p> Obsługuje jeszcze jeden <a href="/hello">zasób</a> </p>
        <body>

</html>`))

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`<html>
    <head><title>Hello</title></head>
    <body>
        <h1>Witaj!</h1>
        <p> Strona prezentuje zasób "hello"</p>
        <body>
</html>`))

}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8081", nil)
}