package params

type TimePeriodGet struct {
	Key string `json:"key"`
}

type TimePeriod struct {
	Name      string   `json:"timeperiod_name"`
	Sunday    string   `json:"sunday"`
	Monday    string   `json:"monday"`
	Tuesday   string   `json:"tuesday"`
	Wednesday string   `json:"wednesday"`
	Thursday  string   `json:"thursday"`
	Friday    string   `json:"friday"`
	Saturday  string   `json:"saturday"`
	Exclude   []string `json:"exclude"`
}
