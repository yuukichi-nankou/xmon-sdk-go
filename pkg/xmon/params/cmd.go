package params

type CmdProcessServiceCheckResultSet struct {
	Host    string `json:"host_name"`
	Service string `json:"service_description"`
	Code    string `json:"status_code"` // なぜか String で指定する
	Output  string `json:"plugin_output"`
}
