package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// NewScanner returns a new Scanner to read from r.
	// The split function defaults to ScanLines.
	scanner := bufio.NewScanner(conn)
	// Scan advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method.
	for scanner.Scan() {
		// Text returns the most recent token generated by a call to Scan
		// as a newly allocated string holding its bytes.
		ln := scanner.Text()
		// Fields splits the string s around each instance of one or more consecutive white space
		// characters, as defined by unicode.IsSpace, returning an array of substrings of s or an
		// empty list if s contains only white space.
		fs := strings.Fields(ln)
		// skip blank lines
		if len(fs) == 0 {
			continue
		}
		switch fs[0] {
		case "GET":
		case "SET":
		case "DEL":
		default:
			fmt.Println(ln)
			ln = fmt.Sprint("FROM SERVER - USAGE <GET | SET | DEL> <KEY> [VAL]")
			fmt.Fprintln(conn, ln)
			io.WriteString(conn, ln)
		}
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
}

/*
start main.go (go run main.go) then ...
telnet localhost 9000
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
