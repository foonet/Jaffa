// Package main is the entry point of whole program.
package main

import (
	"flag"
	"io"
	"os"
	"runtime"

	log "./logrus"
	"./core"
)

// For auto version building
//  go build -ldflags "-X main.version=version"
var version string

func main() {

	var (
		configPath      string
		logPath         string
		isLogVerbose    bool
		processorNumber int
	)

	flag.StringVar(&configPath, "c", "./config.json", "config file path")
	flag.StringVar(&logPath, "l", "", "log file path")
	flag.BoolVar(&isLogVerbose, "v", false, "verbose mode")
	flag.IntVar(&processorNumber, "p", runtime.NumCPU(), "number of processor to use")
	flag.Parse()

	if isLogVerbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if logPath != "" {
		lf, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			println("Logfile error: Please check your log file path")
		} else {
			log.SetOutput(io.MultiWriter(lf, os.Stdout))
		}
	}

	//log.Info("Jaffa " + version)
	log.Info("Jaffa v0.2 - DNS server/forwarder/dispatcher written in Go.")
	//log.Info("If you need any help, please visit the project repository: https://github.com/foonet/jaffa")
	log.Info("ZmEu (zmeu_at_whitehat.ro)")

	runtime.GOMAXPROCS(processorNumber)

	core.InitServer(configPath)
}
