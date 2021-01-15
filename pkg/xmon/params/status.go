package params

type StatusGet struct {
	ServiceProps ServiceProps `json:"serviceprops"`
}

type ServiceProps struct {
	AcceptPassiveChecks bool `json:"accept_passive_checks"`
}

type Status struct {
	Host        string `json:"host_name"`
	Service     string `json:"description"`
	ActiveCheck bool   `json:"active_checks_enabled"`
}
