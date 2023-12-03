package reply

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Gamekeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002064Кости\U00002064"),
		tgbotapi.NewKeyboardButton("\U00002064Дартс\U00002064"),
		tgbotapi.NewKeyboardButton("\U00002064Футбол\U00002064"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002064Автомат\U00002064"),
		tgbotapi.NewKeyboardButton("\U00002064Баскетбол\U00002064"),
		tgbotapi.NewKeyboardButton("\U00002064Боулинг\U00002064"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002064Выход\U00002064"),
	),
)
