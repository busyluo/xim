package rpc

import (
	svr "xim/internal/logic/rpc"
	"xim/internal/pb"
)

func SignIn(r pb.SignInReq) error {
	_, err := svr.SignIn(nil, &r)
	return err
}

func SendMessage(m pb.UpMessage) error {
	_, err := svr.SendMessage(nil, &m)
	return err
}
