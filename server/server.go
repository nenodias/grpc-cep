package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	proto "github.com/nenodias/grpc-cep/protobuf/gen/cep"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CepBody struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Service struct {
	proto.CepServiceServer
}

func (s *Service) GetCep(ctx context.Context, cep *proto.CepRequest) (*proto.CepResponse, error) {
	retorno, err := GetCepBody(cep.GetCep())
	if err != nil {
		return nil, err
	}
	log.Println(retorno)
	return &proto.CepResponse{
		Cep:         retorno.Cep,
		Logradouro:  retorno.Logradouro,
		Complemento: retorno.Complemento,
		Bairro:      retorno.Bairro,
		Localidade:  retorno.Localidade,
		Uf:          retorno.Uf,
		Ibge:        retorno.Ibge,
		Gia:         retorno.Gia,
		Ddd:         retorno.Ddd,
		Siafi:       retorno.Siafi,
		Atualizado:  timestamppb.Now(),
	}, nil
}

func GetCepBody(cep string) (*CepBody, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	retorno := new(CepBody)
	err = json.NewDecoder(req.Body).Decode(retorno)
	if err != nil {
		return nil, err
	}
	return retorno, nil
}

func Run() {
	port := 8080
	serverInfo := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", serverInfo)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service := new(Service)
	var myService proto.CepServiceServer = service
	proto.RegisterCepServiceServer(grpcServer, myService)
	log.Printf("Listening: %s", serverInfo)
	log.Fatal(grpcServer.Serve(lis))
}
