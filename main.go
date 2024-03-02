package main

import (
    "log"
    "net/http"
)

func main() {
    // Set the directory that contains the MP3 files
    musicDir := "./music"

    // Create a file server handler that serves files from the specified directory
    fileServer := http.FileServer(http.Dir(musicDir))

    // Use http.Handle instead of http.HandleFunc since we're serving static files
    // and not implementing any handler functions. Here, we're serving files at the
    // "/music" path. You can change this to "/" to serve at the root path.
    http.Handle("/music/", http.StripPrefix("/music", fileServer))

    // Start the server on port 8080 and log any errors
    log.Println("Serving MP3 files on http://localhost:8080/music/")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

