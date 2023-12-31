package publisher

const (
	TopicLogIn  = "type-log-in"
	TopicLogOut = "type-log-out"

	ChannelMetrics = "metrics"

	ListenLogOut = "log-out"
	ListenLogIn  = "log-in"
)

type MessageLogIn struct {
	PlatNumber  string `json:"plat_number"`
	TownCode    string `json:"town_code"`
	VehicleType string `json:"vehicle_type"`
}

type MessageLogOut struct {
	PlatNumber  string `json:"plat_number"`
	TownCode    string `json:"town_code"`
	VehicleType string `json:"vehicle_type"`
}
