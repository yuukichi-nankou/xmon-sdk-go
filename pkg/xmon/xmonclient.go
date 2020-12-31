package xmon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type XmonClient struct {
	Address string
	Request XmonRequest
	Debug   bool
}

func NewXmonClient(a, u, t string, d bool) *XmonClient {
	return &XmonClient{
		Address: a,
		Request: XmonRequest{
			UserId:    u,
			AuthToken: t,
			Jsonrpc:   "2.0",
			Method:    "",
			Id:        1,
			Lang:      "ja",
		},
		Debug: d,
	}
}

// jsonデータ型定義
type XmonRequest struct {
	UserId    string      `json:"user_id"`
	AuthToken string      `json:"auth_token"`
	Jsonrpc   string      `json:"jsonrpc"`
	Method    string      `json:"method"`
	Params    interface{} `json:"params"`
	Id        int         `json:"id"`
	Lang      string      `json:"lang"`
}

type XmonResult struct {
	Result  interface{} `json:"result"`
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Error   XmonError   `json:"error"`
}

type XmonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (x *XmonClient) Send() (XmonResult, error) {
	rj, _ := json.Marshal(x.Request)
	if x.Debug {
		fmt.Println(string(rj))
	}
	var c http.Client
	resp, err := c.Post(
		"http://"+x.Address+"/1.0/xmon_api.php",
		"application/json",
		strings.NewReader(string(rj)),
	)
	if err != nil {
		return XmonResult{}, err
	}
	if resp.StatusCode != 200 {
		return XmonResult{}, errors.New("ERROR: return code not 200 ")
	}
	// if failed in reading response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return XmonResult{}, errors.New("ERROR: Cannot read responce body! ")
	}

	if x.Debug {
		fmt.Println(string(respBody))
	}

	// decode json result
	var data XmonResult
	err = json.Unmarshal(respBody, &data)
	// if return data from API is not a json, the IP might not be X-MON
	if err != nil {
		return XmonResult{}, errors.New("ERROR: Json parse error! ")
	}

	// close the response body when finished with it
	defer resp.Body.Close()
	return data, nil
}
