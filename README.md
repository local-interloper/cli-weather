# cli-weather

![Go version](https://img.shields.io/badge/go-1.25.7-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-GPL--3.0-blue)
![Built with Bubbletea](https://img.shields.io/badge/built%20with-Bubbletea-EE6FF8?logo=charm&logoColor=white)
![Open-Meteo](https://img.shields.io/badge/weather-Open--Meteo-orange)

> A sleek terminal weather app — no API key, no config, just forecasts.

## Features

- **Worldwide city search** — find any city on the planet
- **7-day forecast** — daily min/max temperatures, precipitation chance, and weather status at a glance
- **Remembers your city** — launches straight into your forecast after the first run
- **Zero setup** — powered by the free [Open-Meteo](https://open-meteo.com/) API

## Installation


<details>
<summary>Go install</summary>

```bash
go install github.com/local-interloper/cli-weather@latest
```

</details>

<details>
<summary>Build from source</summary>

```bash
git clone https://github.com/local-interloper/cli-weather
cd cli-weather
go build -o cli-weather .
```

</details>

## Usage

```bash
cli-weather
```

On first launch you'll be prompted to search for a city. Pick one from the results and your choice is saved — next time you run it, your forecast loads instantly.

To switch cities, press `/` from the forecast view.

## Configuration

Your default city is stored at `~/.config/cli-weather/city.json`. Delete it to reset.

## License

GPL-3.0 — see [LICENSE](LICENSE).
