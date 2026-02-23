package msgs

import "github.com/local-interloper/cli-weather/app/types"

type CitiesQueryResultMsg struct {
	Ok     bool
	Cities []types.City
}
