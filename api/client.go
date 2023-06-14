package api

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(address string) (c HelloClient, err error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	c = NewHelloClient(conn)
	return
}

func NewFSC(address string) (c FSClient, err error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	c = NewFSClient(conn)
	return
}
