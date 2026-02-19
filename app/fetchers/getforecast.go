package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/local-interloper/cli-weather/app/consts"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
)

func GetForecast(coords types.Coords) msgs.ForecastMsg {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/forecast", consts.ForecastApiUrl))

	if err != nil {
		panic(err.Error())
	}

	query := requestUrl.Query()
	query.Add("timezone", "auto")
	query.Add("latitude", fmt.Sprintf("%f", coords.Latitude))
	query.Add("longitude", fmt.Sprintf("%f", coords.Longitude))
	query.Add("daily", "temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,precipitation_probability_max,weather_code")

	requestUrl.RawQuery = query.Encode()

	res, err := http.Get(requestUrl.String())
	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var forecast msgs.ForecastMsg
	json.Unmarshal(body, &forecast)

	return forecast
}
