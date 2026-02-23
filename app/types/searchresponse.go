package types

type SearchResponse struct {
	Results          []City  `json:"results"`
	GenerationTimeMs float32 `json:"generationtime_ms"`
}
