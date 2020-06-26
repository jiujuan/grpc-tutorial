package server

import (
    grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func RecoveryInterceptor() grpc_recovery.Option {
    return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
        return status.Errorf(codes.NotFound, "panic : %v", p)
    })
}
