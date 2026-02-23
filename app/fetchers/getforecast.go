package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/local-interloper/cli-weather/app/consts"
	"github.com/local-interloper/cli-weather/app/types"
)

func GetForecast(city types.City) (types.ForecastResponse, error) {
	var forecast types.ForecastResponse
	requestUrl, err := url.Parse(fmt.Sprintf("%s/forecast", consts.ForecastApiUrl))

	if err != nil {
		return forecast, err
	}

	query := requestUrl.Query()
	query.Add("timezone", "auto")
	query.Add("latitude", fmt.Sprintf("%f", city.Latitude))
	query.Add("longitude", fmt.Sprintf("%f", city.Longitude))
	query.Add("daily", "temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,precipitation_probability_max,weather_code")
	query.Add("current", "temperature_2m,apparent_temperature")

	requestUrl.RawQuery = query.Encode()

	res, err := http.Get(requestUrl.String())
	if err != nil {
		return forecast, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return forecast, err
	}

	if err := json.Unmarshal(body, &forecast); err != nil {
		return forecast, err
	}

	return forecast, nil
}
