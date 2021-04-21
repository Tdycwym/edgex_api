package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/edgexfoundry/edgex-ui-go/configs"
	"github.com/edgexfoundry/edgex-ui-go/proxy"
)

func main() {

	err := configs.LoadConfig()
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
