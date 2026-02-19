package types

type City struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Elevation float32 `json:"elevation"`
}
