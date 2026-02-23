# ⛅ cli-weather

> A snappy terminal weather app — search for any city and get a 7-day forecast, right in your shell.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)
![No API key](https://img.shields.io/badge/API%20key-not%20required-brightgreen?style=flat-square)

```
  Date       Status                    Rain %   Min      Max
  Monday     Clear                     5.0      18.2     26.4
  Tuesday    Cloudy                    20.0     17.8     24.1
  Wednesday  Rain                      75.0     15.3     19.9
  Thursday   Clear                     10.0     16.1     25.0
  Friday     Clear                     0.0      17.4     27.2
  Saturday   Drizzle                   40.0     16.9     22.3
  Sunday     Thunderstorm              85.0     14.5     18.8
```

---

## Installation

**Via `go install`:**

```sh
go install github.com/local-interloper/cli-weather@latest
```

**Or build from source:**

```sh
git clone https://github.com/local-interloper/cli-weather
cd cli-weather
go build -o cli-weather .
```

---

## Usage

```sh
cli-weather
```

The app is fully interactive. On first run it walks you through three screens:

1. **City search** — type a city name and press `Enter`
2. **City picker** — select the matching city from a list, press `Enter` to confirm
3. **Forecast** — view the 7-day forecast table

Your selected city is saved to `~/.config/cli-weather/city.json` and loaded automatically on future runs, so you skip straight to the forecast.

---

## Keybindings

| Screen | Key | Action |
|---|---|---|
| **City search** | `Enter` | Search for the typed city |
| | `Ctrl+C` | Quit |
| **City picker** | `Enter` | Select city and show forecast |
| | `↑` / `↓` | Navigate list |
| | `/` | Filter list |
| | `Ctrl+C` | Back to city search |
| **Forecast** | `S` | Search a new city (clears saved default) |
| | `Q` / `Ctrl+C` | Quit |

---

## Data

Powered by [Open-Meteo](https://open-meteo.com/) — free, open, no account needed.

---

## Built with

| | |
|---|---|
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | TUI framework |
| [Bubbles](https://github.com/charmbracelet/bubbles) | Spinner, table, list & input components |
| [Lip Gloss](https://github.com/charmbracelet/lipgloss) | Terminal styling |
