package main

import "github.com/Garg-Shashwat/Go-User-Authentication/root/transport"

func main() {
	server := transport.RestServer{}
	server.Setup()
	server.Start()
}
