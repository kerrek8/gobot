package game

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gobot/internal/keyboards/reply"
)

func Game(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите игру")
	msg.ReplyMarkup = reply.Gamekeyboard
	bot.Send(msg)
}

func GameHandler(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	switch u.Message.Text {
	case "⁤Выход⁤":
		gameExit(bot, u)
	case "⁤Кости⁤":
		dice(bot, u)
	case "⁤Автомат⁤":
		automat(bot, u)
	case "⁤Боулинг⁤":
		bouling(bot, u)
	case "⁤Футбол⁤":
		football(bot, u)
	case "⁤Баскетбол⁤":
		basketball(bot, u)
	case "⁤Дартс⁤":
		darts(bot, u)
	}
}

func gameExit(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Хорошего настроения")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func automat(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "🎰")
	bot.Send(msg)
}

func dice(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "🎲")
	bot.Send(msg)
}

func football(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "⚽")
	bot.Send(msg)
}

func basketball(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "🏀")
	bot.Send(msg)
}

func darts(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "🎯")
	bot.Send(msg)
}

func bouling(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewDiceWithEmoji(u.Message.Chat.ID, "🎳")
	bot.Send(msg)
}
