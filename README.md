# 🌦 WeatherBot — Telegram Bot in Go

A Telegram bot written in Go that returns real-time weather data using the OpenWeather API.

The bot integrates with Telegram via `go-telegram-bot-api` and uses a layered architecture for handling updates and external API requests.

---

## Features

- Retrieve real-time weather data for any city
- Integration with Telegram Bot API
- OpenWeather API client for weather requests
- Modular project structure (handlers, API client, models)
- Environment-based configuration using `.env`
- Dockerized application setup for consistent development and deployment
- PostgreSQL integration for storing user queries and bot data
- Container orchestration using Docker Compose

---

## Tech Stack

- Go
- PostgreSQL
- Docker
- Telegram Bot API (`go-telegram-bot-api/v5`)
- OpenWeather API
- godotenv

---

## Example

```

User: /city Almaty
Bot: The Almaty is saved
User: /weather
Bot: Weather in Almaty — 18°C, clear sky

```
