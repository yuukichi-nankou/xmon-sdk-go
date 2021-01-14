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
	status := cmd.StatsuGet(c, params.StatusGet{})
	fmt.Println(status)
}
