package reply

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Weatherkeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002063Сейчас\U00002063"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002063Сегодня\U00002063"),
		tgbotapi.NewKeyboardButton("\U00002063Завтра\U00002063"),
		tgbotapi.NewKeyboardButton("\U00002063Послезавтра\U00002063"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002063На 5 дней\U00002063"),
	),
)

var DayKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002063Стрежевой\U00002063"),
		tgbotapi.NewKeyboardButton("\U00002063Тюмень\U00002063"),
	),
)
