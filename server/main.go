package main

import (
	"api_grpc/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

//========== Funçoes que serão usados pelo servidor =========
// grcp não trabalha com rotas!!
type server struct{}

//    __||__         func Hello(parm, ..) 														retorno da função
func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello " + request.GetName() // esse metodo foi gerado pelo protobuf

	//response
	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

//=========== temos aqui a criação do servidor ============
func main() {
	//criando conexão Tcp
	lis, errConnect := net.Listen("tcp", "0.0.0.0:50051")

	//tratando erro de conexão
	if errConnect != nil {
		log.Fatalf("Falha na conexão %v", errConnect)
	}

	//instancia do servidor -> parecido com o express
	grpcServer := grpc.NewServer()

	//pb -> faz uso dos metodos que foram gerados pelo protobuf
	//registra o servidor no protobuf
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	//tratando erro de disponibilidade do servidor
	errServer := grpcServer.Serve(lis)
	if errServer != nil {
		log.Fatalf("falha no servidor: %v", errServer)
	}

}
