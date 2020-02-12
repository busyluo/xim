package configs

var TcpServerConf tcpServerConf

type tcpServerConf struct {
	TcpListenAddr string
}

func init() {
	initDevConf()
}
