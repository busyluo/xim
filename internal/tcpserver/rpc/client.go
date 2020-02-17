package rpc

import (
	svr "xim/internal/logic/rpc"
	"xim/internal/pb"
)

var server svr.LogicRpcServer

func SignIn(req pb.SignInReq) error {
	_, err := server.SignIn(nil, &req)
	return err
}

func SendMessage(message pb.UpMessage) error {
	return nil
}
