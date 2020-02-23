package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"xim/internal/app/tcpserver/logic"
	"xim/internal/pkg/pb"
)

type TcpServer struct{}

func NewTcpServer() *TcpServer {
	return &TcpServer{}
}

func (TcpServer) DispatchMessage(_ context.Context, req *pb.DispatchReq) (*empty.Empty, error) {
	err := logic.DispatchMessage(req.Id, req.Msg)
	return nil, err
}
