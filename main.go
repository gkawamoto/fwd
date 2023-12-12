package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("%s <from> <to>", os.Args[0])
		return
	}

	from := os.Args[1]
	to := os.Args[2]
	ln, err := net.Listen("tcp", from)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn, from, to)
	}
}

func handleRequest(conn net.Conn, from, to string) {
	log.Println("client", from, "=>", to)

	proxy, err := net.Dial("tcp", to)
	if err != nil {
		panic(err)
	}

	log.Println("connected", from, "=>", to)
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
