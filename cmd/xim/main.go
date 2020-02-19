package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"xim/configs"
	"xim/internal/tcpserver"
)

func main() {
	log.SetOutput(os.Stdout)
	tcpserver.Run(configs.TcpServerConf.TcpListenAddr)
}
