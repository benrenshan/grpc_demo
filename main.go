package main

import (
	"google.golang.org/grpc"
	"grpc_demo/service"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, service.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}
