package main

import (
	"github.com/Garg-Shashwat/Go-User-Authentication/config"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/transport"
)

func main() {
	config.InitiateConfig()

	server := transport.RestServer{}
	server.Setup()
	server.Start()
}
