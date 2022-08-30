package client

import (
	"context"
	"fmt"

	proto "github.com/nenodias/grpc-cep/protobuf/gen/cep"
	"google.golang.org/grpc"
)

func Run() {
	port := 8080
	serverInfo := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(serverInfo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	service := proto.NewCepServiceClient(conn)
	fmt.Println("Digite o cep: ")
	var cep string
	fmt.Scan(&cep)
	input := proto.CepRequest{Cep: cep}
	response, err := service.GetCep(context.Background(), &input)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
