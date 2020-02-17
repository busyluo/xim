package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"xim/internal/pb"
)

type LogicRpcServer struct {
	pb.UnimplementedLogicServer
}

func (LogicRpcServer) SignIn(context.Context, *pb.SignInReq) (*empty.Empty, error) {
	panic("implement me")
}

func (LogicRpcServer) SendMessage(context.Context, *pb.UpMessage) (*empty.Empty, error) {
	panic("implement me")
}
