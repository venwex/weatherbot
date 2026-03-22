package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/venwex/weatherbot/clients/openweather"
	"github.com/venwex/weatherbot/handler"
	"github.com/venwex/weatherbot/repository"
	"github.com/venwex/weatherbot/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := InitDB()
	bot := InitBot()
	
	owClient := openweather.New(os.Getenv("OPENWEATHERAPI_KEY"))

	repo := repository.NewUserRepo(db)
	svc := service.NewService(repo)
	handler := handler.New(bot, owClient, svc)

	handler.Start()
}

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error during connecting to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("error pinging the database")
	}

	return db
}

func InitBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}
