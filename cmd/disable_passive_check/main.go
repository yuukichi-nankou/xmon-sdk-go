package main

import (
	"fmt"

	core "github.com/yuukichi-nankou/xmon-sdk-go/cmd"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/cmd"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func main() {
	c, err := core.LoadXmonClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 監視時間帯が 24*7 でないサービスリスト
	services, err := fetchServiceWithout247(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(services)

	// 監視時間帯のリスト
	timePeriods, err := cmd.TimePeriodGet(c, params.TimePeriodGet{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timePeriods)

	// 除外時間帯のリスト

	// パッシブチェックが無効なサービスリスト
	//status := cmd.StatusGet(c, params.StatusGet{ServiceProps: params.ServiceProps{AcceptPassiveChecks: false}})
	//fmt.Println(status)

	// 監視時間外ならパッシブチェックを停止
	// 監視時間内ならパッシブチェックを開始
}

// 監視時間帯が 24*7 でないサービスリスト
func fetchServiceWithout247(c *xmon.XmonClient) ([]params.Service, error) {
	var result []params.Service
	services, err := cmd.ServicesGet(c, params.ServicesGet{})
	if err != nil {
		return result, err
	}

	for _, service := range services {
		if service.CheckPeriod == "24x7" {
			continue
		}
		result = append(result, service)
	}
	return result, err

}
