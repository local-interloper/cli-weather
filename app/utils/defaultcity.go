package utils

import (
	"encoding/json"
	"os"

	"github.com/local-interloper/cli-weather/app/consts"
	"github.com/local-interloper/cli-weather/app/types"
)

func SetDefaultCity(city types.City) {
	if _, err := os.Stat(consts.ConfigDirectoryPath); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(consts.ConfigDirectoryPath, 0o770)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}

	serializedCity, err := json.Marshal(city)
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(consts.DefaultCityFile, serializedCity, 0o770)
	if err != nil {
		panic(err.Error())
	}
}

func GetDefaultCity() (types.City, error) {
	var city types.City

	serialziedCity, err := os.ReadFile(consts.DefaultCityFile)
	if err != nil {
		return city, err
	}

	err = json.Unmarshal(serialziedCity, &city)
	if err != nil {
		os.Remove(consts.DefaultCityFile)
		return city, err
	}

	return city, nil
}

func ClearDefaultCity() error {
	return os.Remove(consts.DefaultCityFile)
}
