package app

import (
	"log"
	"net"
	"net/rpc"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

func Run() {
	my_service := new(internal.Service)
	rpc.Register(my_service)

    // Creating a listener on a local machine on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
    defer listener.Close()

    log.Println("The server is running, listening to port 1234...")

    // Infinite loop for processing incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func() {
            serverCodec := models.GetServerCodec(conn)
			rpc.ServeCodec(serverCodec)
		}()
	}
}