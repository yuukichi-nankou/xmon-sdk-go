package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func ServicesGet(x *xmon.XmonClient, p params.ServicesGet) ([]params.Service, error) {
	var services []params.Service

	x.Request.Method = "services.get"
	x.Request.Params = p
	ret, err := x.Send()
	if err != nil {
		return services, err
	}
	if ret.Error.Message != "" {
		fmt.Println(ret.Error.Message)
		return services, err
	}

	for _, v := range ret.Result.([]interface{}) {
		tmp := params.Service{}
		for key, value := range v.(map[string]interface{}) {
			switch key {
			case "host_name":
				tmp.Host = value.(string)
			case "service_description":
				tmp.Name = value.(string)
			case "check_command":
				tmp.Command = value.(string)
			case "active_checks_enabled":
				tmp.ActiveCheck = value.(bool)
			case "check_period":
				tmp.CheckPeriod = value.(string)
			}
		}
		services = append(services, tmp)
	}
	return services, nil

}

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
