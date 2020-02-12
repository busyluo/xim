package configs

func initDevConf() {
	TcpServerConf = tcpServerConf{
		TcpListenAddr: "127.0.0.1:6060",
	}
}
