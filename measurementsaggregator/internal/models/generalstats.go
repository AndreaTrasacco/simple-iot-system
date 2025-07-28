package models

type GeneralStats struct {
	DeviceID string  `json:"deviceId" bson:"deviceId"`
	Metric   string  `json:"metric" bson:"metric"`
	Avg      float64 `json:"avg" bson:"avg"`
	Min      float64 `json:"min" bson:"min"`
	Max      float64 `json:"max" bson:"max"`
}
