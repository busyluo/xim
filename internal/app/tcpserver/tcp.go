package tcpserver

import (
	_ "github.com/golang/protobuf/proto"
	"net"
	"xim/internal/app/tcpserver/logic"
	"xim/internal/app/tcpserver/rpc/server"
	"xim/internal/pkg/local_call"
)

type TcpServer struct {
}

func initRpc() {
	local_call.TcpServer = server.NewTcpServer()
}

func Run(addr string) {
	initRpc()

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := lis.AcceptTCP()
		conn.SetKeepAlive(true)

		go logic.ServeConn(conn)
	}
}

//func Run(addr string) {
//	log.Info("tcp local started.")
//	local := framedserver.NewFramedServer(addr)
//	log.Info("listening...")
//
//	for {
//		conn, err := local.Accept()
//		conn.SetKeepAlive(true)
//		conn.SetReadBuffer(4096)
//		conn.SetWriteBuffer(4096)
//		if err == nil {
//			log.Error(err)
//			continue
//		}
//		go ServeConn(conn)
//	}
//}
