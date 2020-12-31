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

	ret := []params.Host{}
	for i := 1; i <= 200; i++ {
		ret = append(ret, buildHost(fmt.Sprintf("HOST_%04d", i), "192.168.0.1"))
	}
	cmd.HostsOverwrite(c, ret)
}

func buildHost(n, a string) params.Host {
	return params.Host{
		Name:        n,
		Address:     a,
		Command:     "check-host-alive",
		ActiveCheck: false,
	}

}
