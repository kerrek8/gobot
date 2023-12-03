package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsCommand(u tgbotapi.Update) bool {
	return u.Message != nil && u.Message.IsCommand()
}

func IsText(u tgbotapi.Update) bool {
	return !u.Message.IsCommand() && u.Message != nil
}

func IsNil(u tgbotapi.Update) bool {
	return u.Message == nil
}

func IsWeather(u tgbotapi.Update) bool {
	r := []rune(u.Message.Text)
	return string(r[0]) == "⁣" && string(r[len(r)-1]) == "⁣"
}

func IsGame(u tgbotapi.Update) bool {
	r := []rune(u.Message.Text)
	return string(r[0]) == "⁤" && string(r[len(r)-1]) == "⁤"
}

func IsHoro(u tgbotapi.Update) bool {
	r := []rune(u.Message.Text)
	return string(r[0]) == "⁢" && string(r[len(r)-1]) == "⁢"
}
