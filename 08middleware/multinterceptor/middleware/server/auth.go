package server

import (
    "context"
    "errors"
    grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
)

type Auth struct {
    User string
    Password string
}

// AuthInterceptor 拦截器, 采用 `:authorization` header，头部的数据形式是固定的，形如 "basic/bearer token"
// 可以看文件 grpc-ecosystem/go-grpc-middleware/auth/metadata.go 定义
func AuthInterceptor(ctx context.Context) (context.Context, error) {
    token, err := grpc_auth.AuthFromMD(ctx, "bearer")
    log.Println("Req token val: ", token)
    if err != nil {
        return nil, err
    }
    auth, err := decodeToken(token)
    if err != nil {
        return nil, status.Errorf(codes.Unavailable, " %v ", err)
    }
    newCtx := context.WithValue(ctx, auth.User, auth)
    log.Println(newCtx.Value(auth.User))
    return newCtx, nil
}

// decodeToken 解码token函数
// 当然具体的解码过程没写，只是模拟一下
func decodeToken(token string) (Auth, error) {
    var auth Auth
    if token == "grpc-auth-user" {
        // 假设一下解码出来的数据
        auth.User = "test.token"
        auth.Password = "grpc.auth.test"
        return auth, nil
    }
    return auth, errors.New("decodeToken error: "+token)
}