package service

import (
	"context"
	"github.com/fun-dev/fun-cloud-protobuf/auth/rpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type AuthService struct {}

func (a AuthService) GetUserInfo(ctx context.Context, empty *empty.Empty) (*rpc.GetUserInfoResponse, error) {
	panic("implement me")
}

func (a AuthService) SignIn(ctx context.Context, request *rpc.SignInRequest) (*rpc.SignInResponse, error) {
	authType := request.AuthType
	switch authType {
	// Add Password Verification
	case rpc.AuthType_BASIC:
	case rpc.AuthType_GOOGLE_OAUTH:
	default:

	}
	panic("implement me")
}

func (a AuthService) SignOut(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
	panic("implement me")
}

func (a AuthService) SignUp(ctx context.Context, request *rpc.SignUpRequest) (*rpc.SignUpResponse, error) {
	panic("implement me")
}

func NewAuthService() rpc.AuthServiceServer {
	return &AuthService{}
}