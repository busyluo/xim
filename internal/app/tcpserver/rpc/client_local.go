package rpc

import (
	"context"
	"xim/internal/pkg/local_call"
	"xim/internal/pkg/pb"
)

func SignIn(_ context.Context, r pb.SignInReq) error {
	_, err := local_call.LogicServer.SignIn(nil, &r)
	return err
}

func SendMessage(_ context.Context, m pb.UpMessage) error {
	_, err := local_call.LogicServer.SendMessage(nil, &m)
	return err
}
