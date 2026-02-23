package consts

import (
	"os"
	"path"
)

var ConfigDirectoryPath string
var DefaultCityFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	ConfigDirectoryPath = path.Join(homeDir, ".config/cli-weather/")
	DefaultCityFile = path.Join(ConfigDirectoryPath, "city.json")
}
