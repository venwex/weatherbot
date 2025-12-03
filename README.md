# 🌦️ WeatherBot — Telegram Bot in Go

A simple Telegram bot written in Go that retrieves weather data using the OpenWeather API.  
The bot uses the official `go-telegram-bot-api` library and is structured into clear, separate packages.

---

## 🚀 Features

- Fetches weather information from the OpenWeather API
- Uses `go-telegram-bot-api/v5` for interacting with Telegram
- Clean separation of logic:
  - `handler` — bot message/command handling
  - `clients/openweather` — HTTP client for weather API requests
  - `models` — data structures for JSON decoding
- Loads environment variables using `godotenv`
- Simple and easy-to-extend architecture

---

## 📂 Project Structure

```plaintext
.
├── handler/
│   └── handler.go             # Bot handler logic
│
├── clients/
│   └── openweather/
│       └── openweather.go     # Wrapper for OpenWeather API requests
│
├── models/
│   └── ...                    # Structs for decoding API responses
│
├── main.go                    # Bot initialization and wiring
├── .env                       # API keys (not included in repo)
├── go.mod
└── go.sum
