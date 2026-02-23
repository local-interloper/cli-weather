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

func GetCities(cityQuery string) (types.SearchResponse, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/search", consts.GeocodingApiUrl))
	if err != nil {
		return types.SearchResponse{}, err
	}

	query := requestUrl.Query()
	query.Add("name", cityQuery)

	requestUrl.RawQuery = query.Encode()

	res, err := http.Get(requestUrl.String())
	if err != nil {
		return types.SearchResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.SearchResponse{}, err
	}

	var response types.SearchResponse
	json.Unmarshal(body, &response)

	if len(response.Results) == 0 {
		return types.SearchResponse{}, fmt.Errorf("Failed to find city with name %s", cityQuery)
	}

	return response, nil
}
