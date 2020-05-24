package cloudauth

import (
	"context"
	"errors"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var (
	ContextKeyUserID = "USER_ID"
	// APP Error
	ErrRetrieveMetadataFromContext = errors.New("failed retrieve metadata from context")
	// GRPC Error
	ErrRetrieveTokenFromContext = status.Error(codes.Internal, "failed retrieve token from context")
)

func GetValueFromMetadata(ctx context.Context, key string) ([]string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrRetrieveMetadataFromContext
	}
	return md.Get(key), nil
}

func AuthFuncOnGRPC(ctx context.Context) (context.Context, error) {
	_, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, ErrRetrieveTokenFromContext
	}
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, nil
	}
	// Authentication
	newCtx := context.WithValue(ctx, cloudutil.ContextKeyRequestId, "ok")
	newCtx = context.WithValue(ctx, cloudutil.ContextKeyUserIp, p.Addr.String())
	newCtx = context.WithValue(ctx, "result", "ok")
	return newCtx, nil
}
