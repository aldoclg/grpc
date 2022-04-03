package app

import (
    "context"
    "fmt"
    pb "github.com/aldoclg/grpc/package/proto/v1/service"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type GreeterServiceServer struct {
    pb.UnimplementedGreeterServiceServer
}

func (gs GreeterServiceServer) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
    if req.Msg == nil {
        err = status.New(codes.InvalidArgument, "Message connot  be empty").Err()
        return nil, err
    }

    helloMsg := fmt.Sprintf("%s, %s", req.Msg.Greeting.String(), req.Msg.Name)

    res = &pb.GreetResponse{
        Resp: helloMsg,
    }

    return res, nil
}

