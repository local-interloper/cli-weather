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

func GetCity(cityName string) *types.City {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/search", consts.GeocodingApiUrl))
	if err != nil {
		panic(err.Error())
	}

	query := requestUrl.Query()
	query.Add("name", cityName)

	requestUrl.RawQuery = query.Encode()

	res, err := http.Get(requestUrl.String())
	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var response types.SearchResponse
	json.Unmarshal(body, &response)

	if len(response.Results) == 0 {
		return nil
	}

	return &response.Results[0]
}
