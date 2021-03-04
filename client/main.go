package main

import (
	"api_grpc/pb"
	"context"
	"log"

	"google.golang.org/grpc"
)

//======== criado conexão com um servidor grpc ======
func main() {
	//criando conexão insergura
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer connection.Close() // ao finalizar a troca de msg ele encerra conexão

	//instanciando o cliente
	client := pb.NewHelloServiceClient(connection)

	Hello(err, client)
}

//======= enviando nossa requisição ao servidor ===========
func Hello(err error, client pb.HelloServiceClient) {
	request := &pb.HelloRequest{
		Name: "Jp",
	}

	// esssa é a chamada da funçao que existe no servidor -> seria como o endpoint rest
	res, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Fatalf("Erro durante a execução: %v", err)
	}

	//exibição da mensagem
	log.Println(res.Msg)
}
