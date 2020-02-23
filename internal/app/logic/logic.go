package logic

import (
	"xim/internal/app/logic/rpc"
	"xim/internal/pkg/local_call"
)

func initRpc() {
	local_call.LogicServer = rpc.NewServer()
}

func init() {
	initRpc()
}
