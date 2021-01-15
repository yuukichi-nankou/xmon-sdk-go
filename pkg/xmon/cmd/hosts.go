package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

// TODO クラス化したい

func HostsAppend(x *xmon.XmonClient, hs []params.Host) bool {

	for _, h := range hs {
		var ret bool
		if HostsExist(x, h) {
			// すでに存在するホストは追加しない
			continue
		}
		ret = HostsAdd(x, h)
		if !ret {
			return false
		}
	}
	return true
}

func HostsOverwrite(x *xmon.XmonClient, hs []params.Host) bool {

	for _, h := range hs {
		var ret bool
		if HostsExist(x, h) {
			// 存在する場合は上書き
			ret = HostsUpdate(x, h)
		} else {
			// 存在しない場合は追加
			ret = HostsAdd(x, h)
		}
		if !ret {
			return false
		}
	}

	return true
}

func HostsExist(x *xmon.XmonClient, h params.Host) bool {
	x.Request.Method = "hosts.exist"
	x.Request.Params = h
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

func HostsAdd(x *xmon.XmonClient, h params.Host) bool {
	x.Request.Method = "hosts.add"
	x.Request.Params = h
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

func HostsUpdate(x *xmon.XmonClient, h params.Host) bool {
	x.Request.Method = "hosts.update"
	x.Request.Params = h
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
