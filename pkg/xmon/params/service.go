package params

type ServicesGet struct {
	HostKey string `json:"host_key"`
	Key     string `json:"key"`
}

type Service struct {
	Host        string `json:"host_name"`
	Name        string `json:"service_description"`
	Command     string `json:"check_command"`
	ActiveCheck bool   `json:"active_checks_enabled"`
	CheckPeriod string `json:"check_period"`
}
