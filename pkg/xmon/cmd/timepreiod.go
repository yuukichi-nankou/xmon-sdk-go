package cmd

import (
	"fmt"

	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon"
	"github.com/yuukichi-nankou/xmon-sdk-go/pkg/xmon/params"
)

func TimePeriodGet(x *xmon.XmonClient, s params.TimePeriodGet) ([]params.TimePeriod, error) {
	var timePeriods []params.TimePeriod

	x.Request.Method = "timeperiod.get"
	x.Request.Params = s
	ret, err := x.Send()
	if err != nil {
		return timePeriods, err
	}
	if ret.Error.Message != "" {
		return timePeriods, err
	}

	fmt.Println("test")
	for _, v := range ret.Result.([]interface{}) {
		tmp := params.TimePeriod{}
		for key, value := range v.(map[string]interface{}) {
			switch key {
			case "timeperiod_name":
				tmp.Name = value.(string)
			case "sunday":
				tmp.Sunday = value.(string)
			case "monday":
				tmp.Monday = value.(string)
			case "tuesday":
				tmp.Tuesday = value.(string)
			case "wednesday":
				tmp.Wednesday = value.(string)
			case "thursday":
				tmp.Thursday = value.(string)
			case "friday":
				tmp.Friday = value.(string)
			case "saturday":
				tmp.Saturday = value.(string)
			case "exclude":
				var tes []string
				for _, te := range value.([]interface{}) {
					tes = append(tes, te.(string))
				}
				tmp.Exclude = tes
			}
		}
		timePeriods = append(timePeriods, tmp)
	}
	return timePeriods, nil
}
