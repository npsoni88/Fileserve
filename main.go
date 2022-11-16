package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const usage = `Usage: fserve $PATH $PORT
-- PATH is the path to serve the files from
-- PORT is the value to expose the server on`

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	path := args[1]
	port := args[2]
	fmt.Println(path)

	fileHandler := http.FileServer(http.Dir(path))
	log.Fatal(http.ListenAndServe(":"+port, fileHandler))
}
