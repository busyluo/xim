package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"xim/internal/app/httpserver/router"
	"xim/internal/app/logic/repository"
	"xim/internal/app/tcpserver"
	"xim/internal/pkg/config"
)

func main() {
	log.SetOutput(os.Stdout)

	db, err := gorm.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}

	// init repository
	repository.InitOrmRepository(db)

	// start tcp local
	go func() {
		tcpserver.Run(config.TcpServerConf.TcpListenAddr)
	}()

	// Echo instance
	e := echo.New()
	// Routes
	router.Route(e)

	// Start local
	err = e.Start(":8080")
	log.Fatal(err)
}
