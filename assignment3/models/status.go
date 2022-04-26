package models

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusData struct {
	Status Status `json:"status"`
}

type GetStatus struct {
	WaterStatus float64
	WindStatus  float64
}
