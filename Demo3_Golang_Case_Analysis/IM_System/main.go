package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
