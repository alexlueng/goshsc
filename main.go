package main

import (
	"flag"
	"os"

	"github.com/alexlueng/goshsc/myhttp"
)

/*
	1. Fetch some "assets" we will need
	2. Check if the requested URL is a directory or a file
	3. If it is a dir
		- Read the directories content
		- Iterate over it and fill an items struct with the content
		- Sort the struct
		- Parse the template by providing a directory struct
	4. If it is a file
		- Read the files content
		- Write file to http.ResponseWriter
*/

var (
	port    = 8000
	webroot = "."
)

func init() {
	wd, _ := os.Getwd()

	flag.IntVar(&port, "p", port, "The port")
	flag.StringVar(&webroot, "d", wd, "Web root directory")

	flag.Parse()
}

func main() {
	server := &myhttp.FileServer{
		Port:    port,
		Webroot: webroot,
	}
	server.Start()
}
