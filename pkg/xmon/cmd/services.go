package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func ServicesAppend(x *xmon.XmonClient, ps []params.Service) bool {

	for _, p := range ps {
		// TODO 存在確認して、なければUpdate
		ret := ServicesAdd(x, p)
		if !ret {
			return false
		}
	}

	return true
}

func ServicesAdd(x *xmon.XmonClient, p params.Service) bool {
	x.Request.Method = "services.add"
	x.Request.Params = p
	ret, err := x.Send()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if ret.Error.Message != "" {
		fmt.Println(ret.Error.Message)
		return false
	}
	return ret.Result.(bool)
}
