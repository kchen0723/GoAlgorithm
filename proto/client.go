package main

//https://www.cnblogs.com/davis12/p/18111154

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// client := NewAuthServiceClient(conn)
	// resp, err := client.Auth(context.Background(), &AuthRequest{Name: "admin", Password: "123456"})
	client := NewMathServiceClient(conn)
	resp, err := client.Sum(context.Background(), &MathRequest{Num1: 1, Num2: 2})
	if err != nil {
		panic(err)
	}
	log.Printf("auth grpc reponse: ", resp.Reply)
}
