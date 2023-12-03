package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var weather = tgbotapi.BotCommand{Command: "/weather", Description: "Погода"}
var game = tgbotapi.BotCommand{Command: "/game", Description: "Игры"}
var horo = tgbotapi.BotCommand{Command: "/horo", Description: "Гороскопы"}
var lightning = tgbotapi.BotCommand{Command: "/calc", Description: "Калькулятор молний"}
var world_tour = tgbotapi.BotCommand{Command: "/world_tour", Description: "Путешествие по миру"}

func CommandsReq() tgbotapi.SetMyCommandsConfig {
	commands := tgbotapi.NewSetMyCommands(
		weather,
		game,
		horo,
		lightning,
		world_tour,
	)
	return commands
}
