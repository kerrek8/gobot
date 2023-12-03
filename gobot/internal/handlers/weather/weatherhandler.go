package weather

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gobot/internal/keyboards/reply"
)

func Weather(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите город")
	msg.ReplyMarkup = reply.DayKeyboard
	bot.Send(msg)
}

func WeatherHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Text {
	case "\U00002063Стрежевой\U00002063":
		strej(bot, update)
	case "\U00002063Тюмень\U00002063":
		tymen(bot, update)
	case "\U00002063Сейчас\U00002063":
		now(bot, update)
	case "\U00002063Сегодня\U00002063":
		today(bot, update)
	case "\U00002063Завтра\U00002063":
		tomor(bot, update)
	case "\U00002063Послезавтра\U00002063":
		afterTomorrow(bot, update)
	case "\U00002063На 5 дней\U00002063":
		fiveDays(bot, update)
	}
}
