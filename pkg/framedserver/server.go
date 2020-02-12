package framedserver

import "net"

type FramedServer struct {
	i        int
	listener *net.TCPListener
}

func NewFramedServer(addr string) *FramedServer {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}

	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	return &FramedServer{
		1,
		lis,
	}
}

func (svr *FramedServer) Accept() (*FramedConn, error) {
	conn, err := svr.listener.AcceptTCP()
	return &FramedConn{conn}, err
}
