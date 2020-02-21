package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"xim/internal/app/tcpserver"
	"xim/internal/pkg/config"
)

func main() {
	log.SetOutput(os.Stdout)
	tcpserver.Run(config.TcpServerConf.TcpListenAddr)
}
