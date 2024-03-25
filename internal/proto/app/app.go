package app

import (
	"context"
	"fmt"
	"log"
	"net"

	auth "github.com/xxlifestyle/auth_service/gen/go"
	"github.com/xxlifestyle/auth_service/pkg/config"
	"google.golang.org/grpc"
)

type myAuthServer struct {
	auth.UnimplementedAuthServer
}

func (s myAuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		Token: "test_token",
	}, nil
}
func (s myAuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return &auth.RegisterResponse{
		Token: "test_token",
	}, nil
}

func Run(cfg *config.ServiceConfig) {
	fmt.Println(cfg)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("can't listen hostname< error: %v", err)
	}
	serverRegistrar := grpc.NewServer()
	service := myAuthServer{}
	auth.RegisterAuthServer(serverRegistrar, service)

	serverRegistrar.Serve(lis)
}
