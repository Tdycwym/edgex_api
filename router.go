package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_api/handler"
	"github.com/nju-iot/edgex_api/handler/core_metadata"
	"github.com/nju-iot/edgex_api/wrapper"
)

func registerRouter(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	// your code

	// core_metadata
	r.GET("/api/v1/device", wrapper.JsonOutPutWrapper(core_metadata.GetAllDevice))

}
