package models

// Plants Data Model
type Plants struct {
	PlantID     int     `json:"plantID"`
	PlantType   string  `json:"plantType"`
	Temperature float64 `json:"temperature"`
	PhLevels    float64 `json:"phLevels"`
	Humidity    float64 `json:"humidity"`
	CreatedAt   string  `json:"createdAt"`
}
