package schema

import "time"

type InfoRequest struct {
	Illumination string `json:"illumination"`
	Temperature  string `json:"temperature"`
	Humidity     string `json:"humidity"`
}

type InfoResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"time"`
	Illumination string    `json:"illumination"`
	Temperature  string    `json:"temperature"`
	Humidity     string    `json:"humidity"`
}
