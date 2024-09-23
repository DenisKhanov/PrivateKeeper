package interceptors

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/DenisKhanov/PrivateKeeper/pkg/auth"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// streamAuthMethods specifies the gRPC methods that require authentication.
var streamAuthMethods = map[string]struct{}{
	grpcHandlersPath + "UploadBinaryData": {},
}

// StreamPrivateAuthInterceptor is a gRPC interceptor that enforces authentication for specific streaming RPCs.
// It checks if the incoming context contains a valid token for accessing the specified methods.
// If the token is valid, it extracts the user ID from the token and adds it to the context.
// If the token is missing or invalid, it returns an error.
func StreamPrivateAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Infof("Method '%s'", info.FullMethod)
	if _, exist := streamAuthMethods[info.FullMethod]; !exist {
		return handler(srv, ss)
	}
	var tokenString string
	var err error
	var userID uuid.UUID
	md, ok := metadata.FromIncomingContext(ss.Context())
	if ok {
		values := md.Get("token")
		if len(values) > 0 {
			// ключ содержит слайс строк, получаем первую строку
			tokenString = values[0]
		}
	}
	if !ok || len(tokenString) == 0 {
		return status.Error(codes.InvalidArgument, `missing token`)
	}
	userID, err = auth.GetUUIDFromToken(tokenString)
	if err != nil {
		return status.Error(codes.Unauthenticated, `invalid token`)
	}

	newCtx := context.WithValue(ss.Context(), domain.UserIDKey, userID)
	wrappedStream := &wrappedServerStream{
		ServerStream: ss,
		ctx:          newCtx,
	}
	return handler(srv, wrappedStream)
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}
