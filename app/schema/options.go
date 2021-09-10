package schema

type OptionsRequest struct {
	LED      int `json:"led"`
	FanN     int `json:"fanN"`
	FanM     int `json:"fanM"`
	Engine   int `json:"engine"`
	Smoke    int `json:"smoke"`
	Infrared int `json:"infrared"`
	Card     int `json:"card"`
}

type OptionsResponse struct {
	LED      int `json:"led"`
	FanN     int `json:"fanN"`
	FanM     int `json:"fanM"`
	Engine   int `json:"engine"`
	Smoke    int `json:"smoke"`
	Infrared int `json:"infrared"`
	Card     int `json:"card"`
}
