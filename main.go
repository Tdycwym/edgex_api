package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/nju-iot/edgex_api/configs"
	"github.com/nju-iot/edgex_api/proxy"
)

func main() {

	var confFilePath string
	flag.StringVar(&confFilePath, "conf", "", "Specify local configuration file path")
	flag.Parse()

	err := configs.LoadConfig(confFilePath)
	if err != nil {
		panic(err)
	}
	r := registerRouter()
	server := &http.Server{
		Handler:      proxy.ReverseProxy(r),
		Addr:         ":" + strconv.FormatInt(configs.ServerConf.Port, 10),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
