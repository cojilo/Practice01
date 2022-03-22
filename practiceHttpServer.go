package main

import (

    // HTTP Client , Server Package
    "net/http"

    // JSON Encode , Decode Package
    "encoding/json"

    // ※Go1.6 context / Go 1.12+
      "google.golang.org/appengine/v2"
      "google.golang.org/appengine/v2/memcache"

    "os"
)

// Response Parameter Struct
type Response struct {
    Status          string `json:"status"`
    Message         string `json:"message"`
}

// main関数は実行可能プログラムのエントリポイント
func main() {

    // Args1：URL,Args2:Handler
    http.HandleFunc("/", handle)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // RUN HTTP Server
    appengine.Main()
}

// JSON Response Function
func handle(w http.ResponseWriter, r *http.Request) {
    //Writing w,Response HTTP Request
    json.NewEncoder(w).Encode(Response{Status: "ok", Message: "Test Message"})
}
