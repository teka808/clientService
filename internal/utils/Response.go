package utils

type Response struct {
	Device_id   string `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Timezone    string `json:"timezone"`
	Coordinates Coordinates
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
