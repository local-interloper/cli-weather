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

func GetCity(cityName string) (types.City, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/search", consts.GeocodingApiUrl))
	if err != nil {
		return types.City{}, err
	}

	query := requestUrl.Query()
	query.Add("name", cityName)

	requestUrl.RawQuery = query.Encode()

	res, err := http.Get(requestUrl.String())
	if err != nil {
		return types.City{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.City{}, err
	}

	var response types.SearchResponse
	json.Unmarshal(body, &response)

	if len(response.Results) == 0 {
		return types.City{}, fmt.Errorf("Failed to find city with name %s", cityName)
	}

	return response.Results[0], nil
}
