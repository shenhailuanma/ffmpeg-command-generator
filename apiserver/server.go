package main

import (
	"flag"
	"fmt"
	"github.com/shenhailuanma/ffmpeg-command-generator/apiserver/routers"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	port        = flag.Int("p", 9090, "service port")
	env         = flag.String("e", "prd", "env support: prd,dev, default: prd")
	loglevel    = flag.String("l", "info", "log level: error, warning, info, debug default: info")
	versionFlag = flag.Bool("v", false, "version")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("v1.0.0")
		return
	}

	// init log
	logrus.SetOutput(os.Stdout)
	logrus.Info("service start env:", *env)
	logrus.Info("service start port:", *port)

	// prepare

	// start service
	var listenPort = fmt.Sprintf(":%d", *port)
	err := routers.Run(listenPort, *env)
	if err != nil {
		logrus.Error("service catch error:", err.Error())
		os.Exit(1)
	}
}
