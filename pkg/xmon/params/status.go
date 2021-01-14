package params

type StatusGet struct {
}

type Status struct {
	Host         string `json:"host_name"`
	Service      string `json:"description"`
	ActiveCheck  bool   `json:"active_checks_enabled"`
	PassiveCheck bool   `json:"passive_checks_enabled"`
}
