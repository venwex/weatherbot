package handler

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/venwex/weatherbot/clients/openweather"
	"github.com/venwex/weatherbot/service"
)

type Handler struct {
	bot      *tgbotapi.BotAPI
	owClient *openweather.OpenWeatherClient
	service  *service.Service
}

func New(bot *tgbotapi.BotAPI, owClient *openweather.OpenWeatherClient, svc *service.Service) *Handler {
	return &Handler{
		bot:      bot,
		owClient: owClient,
		service: svc,
	}
}

func (h *Handler) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		h.handleUpdate(update)
	}
}

// message processing
func (h *Handler) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	ctx := context.Background()
	if update.Message.IsCommand() {
		if err := h.ensureUser(ctx, update); err != nil {
			log.Printf("Error during ensuring the user: %v", err)
			return
		}

		switch update.Message.Command() {
		case "city":
			h.HandleSetCity(ctx, update)
			return
		case "weather":
			h.HandleSendWeather(ctx, update)
			return
		default:
			h.HandleUnknownCommand(update)
			return
		}
	}
}

// handlers
func (h *Handler) HandleSetCity(ctx context.Context, update tgbotapi.Update) {
	city := update.Message.CommandArguments()
	if err := h.service.UpdateUserCity(ctx, update.Message.From.ID, city); err != nil {
		log.Printf("error h.service.UpdateUserCity: %v", err)
		h.SendMessage(update, "Error during setting/updating the city")
		return
	}

	h.SendMessage(update, fmt.Sprintf("The %s is saved", city))
}

func (h *Handler) HandleSendWeather(ctx context.Context, update tgbotapi.Update) {
	city, err := h.service.GetUserCity(ctx, update.Message.From.ID)
	if err != nil {
		log.Printf("error h.service.GetUserCity: %v", err)
		h.SendMessage(update, "Error during getting the city")
		return
	}

	if len(city) == 0 {
		h.SendMessage(update, "First, specify the city using the command: /city cityname")
		return
	}

	coordinates, err := h.GetCoordinates(city)
	if err != nil {
		log.Printf("error owClient.Coordinates: %v", err)
		h.SendMessage(update, "Unable to retrieve coordinates, error")
		return
	}

	weather, err := h.GetWeather(coordinates)
	if err != nil {
		log.Printf("error owClient.Weather: %v", err)
		h.SendMessage(update, "Unable to retrieve data, error")
		return
	}

	h.SendMessage(update, fmt.Sprintf(
		"It's currently %d°C in %s with %s.",
		openweather.Convert(weather.Temp),
		city,
		weather.Description,
	))
}

func (h *Handler) HandleUnknownCommand(update tgbotapi.Update) {
	log.Printf("Unknown command [%s] %s", update.Message.From.UserName, update.Message.Text)
	h.SendMessage(update, "That command is not available")
}

func (h *Handler) ensureUser(ctx context.Context, update tgbotapi.Update) error {
	user, err := h.service.GetUser(ctx, update.Message.From.ID)
	if err != nil {
		return fmt.Errorf("Error during getting the user: %v", err)
	}

	if user == nil {
		err := h.service.CreateUser(ctx, update.Message.From.ID)
		if err != nil {
			return fmt.Errorf("Error during creating the user: %v", err)
		}
	}

	return nil
}

// utils
func (h *Handler) SendMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	h.bot.Send(msg)
}

func (h *Handler) GetCoordinates(city string) (openweather.Coordinates, error) {
	return h.owClient.Coordinates(city)
}

func (h *Handler) GetWeather(coordinates openweather.Coordinates) (openweather.Weather, error) {
	return h.owClient.Weather(coordinates)
}
