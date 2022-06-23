package main

import (
    "fmt"
    "time"
    "net/http"
)
import _ "net/http/pprof"

func compute(value int) {
    for i := 0; i < value; i++ {
        time.Sleep(time.Second)
        fmt.Println(i)
    }
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    go compute(5)
    fmt.Fprintf(w, "Hello, World zero !!!!!")
}
func HelloServer1(w http.ResponseWriter, r *http.Request) {
    go compute(6)
    fmt.Fprintf(w, "Hello, W orlds one !!!!!")
}
func HelloServer2(w http.ResponseWriter, r *http.Request) {
    go compute(7)
    fmt.Fprintf(w, "Hello, W orlds two !!!!!")
}
func HelloServer3(w http.ResponseWriter, r *http.Request) {
    go compute(3)
    fmt.Fprintf(w, "Hello, W orlds three !!!!!")
}
func HelloServer4(w http.ResponseWriter, r *http.Request) {
    go compute(4)
    fmt.Fprintf(w, "Hello, W orlds four !!!!!")
}
func main() {

    go compute(10)
    
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/1", HelloServer1)
    http.HandleFunc("/2", HelloServer2)
    http.HandleFunc("/3", HelloServer3)
    http.HandleFunc("/4", HelloServer4)
    http.ListenAndServe(":9092", nil)


}
