package utils

func CodeToStatus(code int) string {
	if code < 20 {
		return "Clear"
	}

	if code < 30 {
		return "Rain"
	}

	if code < 40 {
		return "Sandstorm"
	}

	if code < 50 {
		return "Fog"
	}

	if code < 60 {
		return "Drizzle"
	}

	if code < 70 {
		return "Rain"
	}

	if code < 80 {
		return "Snow"
	}

	return "Thunderstorm"
}
