package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"xim/internal/app/logic/service"
	"xim/internal/pkg/pb"
)

type Server struct {
	pb.UnimplementedLogicServer
}

func NewServer() *Server {
	return &Server{UnimplementedLogicServer: pb.UnimplementedLogicServer{}}
}

func (Server) SignIn(context.Context, *pb.SignInReq) (*empty.Empty, error) {
	panic("implement me")
}

func (Server) SendMessage(ctx context.Context, m *pb.UpMessage) (*empty.Empty, error) {
	// get device id by user id.
	err := service.MessageService.SendMessage(ctx, m)
	return nil, err
}
