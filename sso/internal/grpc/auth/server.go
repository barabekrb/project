package auth

import (
	"context"
	sso "sso1/protos/gen/go"

	"google.golang.org/grpc"
)

type ServerApi struct {
	sso.UnimplementedAuthServer
}

func Register(grpc *grpc.Server) {
	sso.RegisterAuthServer(grpc, &ServerApi{})
}

func (s *ServerApi) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	panic("www")
}

func (s *ServerApi) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	panic("www")
}
func (s *ServerApi) IsAdmin(ctx context.Context, req *sso.IsAdminRequest) (*sso.IsAdminResponse, error) {
	panic("www")
}
