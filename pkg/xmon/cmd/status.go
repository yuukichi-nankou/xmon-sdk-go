package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func StatusGet(x *xmon.XmonClient, s params.StatusGet) []params.Status {
	var status []params.Status

	x.Request.Method = "status.get"
	x.Request.Params = s
	ret, err := x.Send()
	if err != nil {
		fmt.Println(err)
		return status
	}
	if ret.Error.Message != "" {
		fmt.Println(ret.Error.Message)
		return status
	}

	fmt.Println("test")
	for _, v := range ret.Result.([]interface{}) {
		tmp := params.Status{}
		for key, value := range v.(map[string]interface{}) {
			switch key {
			case "host_name":
				tmp.Host = value.(string)
			case "description":
				tmp.Service = value.(string)
			case "active_checks_enabled":
				if value.(string) == "1" {
					tmp.ActiveCheck = true
				} else {
					tmp.ActiveCheck = false
				}
			}
		}
		status = append(status, tmp)
	}
	return status
}
