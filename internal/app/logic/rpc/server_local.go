package rpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"xim/internal/app/logic/service"
	"xim/internal/pkg/pb"
)

// fake rcp. will be called directly.
func SignIn(ctx context.Context, req *pb.SignInReq) (*empty.Empty, error) {
	panic("implement me")
}

func SendMessage(ctx context.Context, m *pb.UpMessage) (*empty.Empty, error) {
	err := service.MessageService.SendMessage(ctx, m)
	return nil, err
}
