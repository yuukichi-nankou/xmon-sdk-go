package main

import (
	"fmt"

	core "github.com/yuukichi-nankou/xmon-sdk-go/cmd"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/cmd"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func main() {
	c, err := core.LoadXmonClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 10; i++ {
		cmd.CmdProcessServiceCheckResultSet(c, buildParams(fmt.Sprintf("HOST_%04d", i), "PING", "0", "test ok"))
	}
}

func buildParams(h, s, c, o string) params.CmdProcessServiceCheckResultSet {
	return params.CmdProcessServiceCheckResultSet{
		Host:    h,
		Service: s,
		Code:    c,
		Output:  o,
	}
}
