package core_metadata

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_api/constdef"
	"github.com/nju-iot/edgex_api/logs"
	"github.com/nju-iot/edgex_api/resp"
	"github.com/nju-iot/edgex_api/wrapper"
	"github.com/parnurzeal/gorequest"
)

func GetAllDevice(c *gin.Context) *wrapper.JsonOutput {
	request := gorequest.New()
	deviceURL := fmt.Sprintf("http://%s:48082/api/v1/device", constdef.Hostname)
	_, bodyBytes, errs := request.Get(deviceURL).EndBytes()
	if len(errs) > 0 {
		logs.Error("[GetAllDevice] request Get failed: errs=%v")
		return wrapper.SampleJson(c, resp.RESP_CODE_RPC_ERROR, nil)
	}
	var data interface{}
	err := json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logs.Error("[GetAllDevice] json unmarshal failed: err=+%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_SEVER_EXCEPTION, nil)
	}
	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, data)
}
