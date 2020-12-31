package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func CmdProcessServiceCheckResultSet(x *xmon.XmonClient, s params.CmdProcessServiceCheckResultSet) bool {
	x.Request.Method = "cmd_process_service_check_result.set"
	x.Request.Params = s
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
