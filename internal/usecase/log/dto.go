package log

type ResponseLog struct {
	TownName    string `json:"town_name"`
	TotalTime   int    `json:"total_time"`
	TimeIn      string `json:"time_in"`
	TimeOut     string `json:"time_out"`
	DateAt      string `json:"date_at"`
	PlatNumber  string `json:"plat_number"`
	VehicleType string `json:"vehicle_type"`
	DateOutAt   string `json:"date_out_at"`
	TownCode    string `json:"town_code"`
}
