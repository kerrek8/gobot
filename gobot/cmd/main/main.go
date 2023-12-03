package main

import (
	"gobot/internal/commands"
	"gobot/internal/handlers/game"
	"gobot/internal/handlers/horo"
	"gobot/internal/handlers/lightning"
	"gobot/internal/handlers/weather"
	"gobot/internal/handlers/worl_tour"
	"gobot/internal/server"
	"gobot/internal/setup"
	"log"
	"os"
)

func main() {

	bot := setup.Botstart(os.Getenv("BOT-TOKEN"))
	//bot.Debug = true
	setup.Webhooksstart(bot)

	com := commands.CommandsReq()
	_, err := bot.Request(com)
	if err != nil {
		log.Fatal(err)
	}
	updates := bot.ListenForWebhook("/" + os.Getenv("BOT-TOKEN"))

	go server.Server()

	for update := range updates {

		if setup.IsNil(update) { // ignore any non-Message updates
			continue
		}

		if setup.IsCommand(update) { // command updates
			switch update.Message.Command() {
			case "weather":
				weather.Weather(bot, update)
			case "game":
				game.Game(bot, update)
			case "horo":
				horo.Horo(bot, update)
			case "calc":
				lightning.Lightning(bot, update)
			case "world_tour":
				worl_tour.WorldHandler(bot, update)
			}
		}

		if setup.IsText(update) { //text update

			//log.Println(update.Message.Text[0], update.Message.Text[len(update.Message.Text)-1])
			if setup.IsWeather(update) {
				weather.WeatherHandler(bot, update)
			}
			if setup.IsGame(update) {
				game.GameHandler(bot, update)
			}
			if setup.IsHoro(update) {
				horo.Horohandler(bot, update)
			}
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
