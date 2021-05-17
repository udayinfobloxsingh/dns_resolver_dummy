package main

import (
	dnspb "dns_resolver_dummy/dnspb"
	"dns_resolver_dummy/server"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50551...")
	listener, err := net.Listen("tcp", ":50551")
	if err != nil {
		log.Fatalf("Unable to listen on port :50551: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := server.DnsServiceServer{}
	dnspb.RegisterDnsServiceServer(s, &srv)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50551")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Done.")
}
