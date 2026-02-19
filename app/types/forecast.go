package types

type Forecast struct {
	Latitude             float32      `json:"latittude"`
	Longtitude           float32      `json:"longtitude"`
	GenerationTime       float32      `json:"generationtime_ms"`
	UtcOffsetSeconds     float32      `json:"utc_offset_seconds"`
	Timezone             string       `json:"timezone"`
	TimezoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            float32      `json:"elevation"`
	DailyUnits           Units        `json:"daily_units"`
	Daily                ForecastData `json:"daily"`
}
