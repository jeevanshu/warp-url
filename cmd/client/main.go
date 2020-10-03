package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8088", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error in connecting to server: %v", err)
	}
	defer conn.Close()

}
