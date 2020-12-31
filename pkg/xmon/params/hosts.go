package params

type Host struct {
	Name        string `json:"host_name"`
	Address     string `json:"address"`
	Command     string `json:"check_command,omitempty"`
	ActiveCheck bool   `json:"active_checks_enabled,omitempty"`
}
