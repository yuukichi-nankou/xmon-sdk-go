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
	cmd.ServicesAppend(
		c,
		loadServices(),
	)
}

func loadServices() []params.Service {
	ret := []params.Service{}

	for i := 1; i < 2000; i++ {
		ret = append(ret, buildService(fmt.Sprintf("HOST_%04d", i), "PING", "check_xmon3_ping"))
	}
	return ret
}

func buildService(h, n, c string) params.Service {
	return params.Service{
		Host:        h,
		Name:        n,
		Command:     c,
		ActiveCheck: true,
	}

}
