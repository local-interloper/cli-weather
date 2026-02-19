# ⛅ cli-weather

> A snappy terminal weather app — get a 7-day forecast for any city, right in your shell.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)
![No API key](https://img.shields.io/badge/API%20key-not%20required-brightgreen?style=flat-square)

```
$ weather Tokyo

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
go build -o weather .
```

---

## Usage

```sh
weather <city>
```

```sh
weather London
weather "New York"
weather Paris
```

A spinner appears while data loads, then the forecast table renders instantly. Press `q` or `Ctrl+C` to exit.

---

## Data

Powered by [Open-Meteo](https://open-meteo.com/) — free, open, no account needed.

---

## Built with

| | |
|---|---|
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | TUI framework |
| [Bubbles](https://github.com/charmbracelet/bubbles) | Spinner & table components |
| [Lip Gloss](https://github.com/charmbracelet/lipgloss) | Terminal styling |
