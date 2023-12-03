package reply

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Horokeyboard1 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002062♉Телец♉\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♋Рак♋\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♐Стрелец♐\U00002062"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002062♑Козерог♑\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♏Скорпион♏\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♒Водолей♒\U00002062"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002062♓Рыбы♓\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♍Дева♍\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♈Овен♈\U00002062"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002062♌Лев♌\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♎Весы♎\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062♊Близнецы♊\U00002062"),
	),
)

var Horokeyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U00002062Сегодня\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062Завтра\U00002062"),
		tgbotapi.NewKeyboardButton("\U00002062Послезавтра\U00002062"),
	),
)
