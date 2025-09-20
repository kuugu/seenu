package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Uint("port", 8090, "port to host HTTP server")
	directoryPath := flag.String("directory", "", "directory to host")

	flag.Parse()

	// we fail the program if directory is not specified
	if *directoryPath == "" {
		fmt.Println("--directory is a required parameter")
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir(*directoryPath))
	addr := fmt.Sprintf("localhost:%v", *port)

	http.Handle("/", fs)

	fmt.Printf("Listening on port %v\n", *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
