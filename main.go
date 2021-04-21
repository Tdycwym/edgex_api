package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/edgexfoundry/edgex-ui-go/configs"
	"github.com/edgexfoundry/edgex-ui-go/proxy"
)

func main() {

	var confFilePath string
	flag.StringVar(&confFilePath, "conf", "", "Specify local configuration file path")
	flag.Parse()

	err := configs.LoadConfig(confFilePath)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      proxy.ReverseProxy(),
		Addr:         ":" + strconv.FormatInt(configs.ServerConf.Port, 10),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
