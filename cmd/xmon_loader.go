package cmd

import (
	"fmt"
	"os"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
)

func LoadXmonClient() (*xmon.XmonClient, error) {

	a := os.Getenv("ADDRESS")
	u := os.Getenv("USER")
	t := os.Getenv("TOKEN")
	if a == "" || u == "" || t == "" {
		return &xmon.XmonClient{}, fmt.Errorf("Faild to get Enviroment")
	}

	debug := true
	d := os.Getenv("DEBUG")
	if d == "" {
		debug = false
	}

	return xmon.NewXmonClient(a, u, t, debug), nil
}
