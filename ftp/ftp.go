package ftp

import (
	"bufio"
	"log"
	"strings"
)

// Serve scans incoming request for a valid command and routes then to the handle functions
func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)

	for s.Scan() {
		input := strings.Fields(s.Text())

		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]

		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd

		case "LIST": // ls

		case "PORT":

		case "USER":

		case "QUIT": //Quit
			// close the connection

		case "RETR": //get

		case "TYPE":

		default:

		}
	}

	if s.Err() != nil {
		log.Println(s.Err)
	}
}
