package main

import (
    "net/http"
    "log"
    "uppy-s3-presign/presigner"
)

func main() {

    index := http.FileServer(http.Dir("./uploader"))
    http.Handle("/", index)
    http.HandleFunc("/sign", presigner.Sign)

    log.Println("presign backend started")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
