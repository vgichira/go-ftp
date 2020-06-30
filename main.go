package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "public", "root directory")
	flag.Parse()
}

func main() {
	server := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", server)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		_, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func handleConnection(c net.Conn) {
	_, err := filepath.Abs(rootDir)

	if err != nil {
		log.Fatalln(err)
	}

	defer c.Close()

	// ftp.Serve(ftp.NewConn(c, absPath))
}
