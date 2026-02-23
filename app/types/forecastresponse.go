package types

type ForecastResponse struct {
	Latitude             float32      `json:"latittude"`
	Longtitude           float32      `json:"longtitude"`
	GenerationTime       float32      `json:"generationtime_ms"`
	UtcOffsetSeconds     float32      `json:"utc_offset_seconds"`
	Timezone             string       `json:"timezone"`
	TimezoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            float32      `json:"elevation"`
	DailyUnits           DailyUnits   `json:"daily_units"`
	CurrentUnits         CurrentUnits `json:"current_units"`
	Daily                Daily        `json:"daily"`
	Current              Current      `json:"current"`
}
