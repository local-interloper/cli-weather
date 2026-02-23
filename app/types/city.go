package types

type City struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Elevation float32 `json:"elevation"`
	Country   string  `json:"country"`
}

func (c City) Title() string       { return c.Name }
func (c City) Description() string { return c.Country }
func (c City) FilterValue() string { return c.Name }
