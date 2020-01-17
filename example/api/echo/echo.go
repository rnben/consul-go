package echo

import (
	"context"
	echo_pb "github.com/ruben-zhi/consul-go/example/echo"
)

type EchoServer struct{}

func (e EchoServer) Say(ctx context.Context, req *echo_pb.Message) (*echo_pb.Message, error) {
	return &echo_pb.Message{
		Msg: req.Msg,
	}, nil
}
