package tcpserver

import (
	_ "github.com/golang/protobuf/proto"
	"net"
)

type TcpServer struct {
}

func Run(addr string) {
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

		go ServeConn(conn)
	}
}

//func Run(addr string) {
//	log.Info("tcp server started.")
//	server := framedserver.NewFramedServer(addr)
//	log.Info("listening...")
//
//	for {
//		conn, err := server.Accept()
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
