package lightning

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func Lightning(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	time := strings.Split(u.Message.CommandArguments(), " ")
	t := time[len(time)-1]
	r, err := strconv.Atoi(t)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(u.Message.Chat.ID, "Некорректное время\nВведите команду и через пробел время, прошедшее со вспышки до грома"))
		return
	}
	r = r * 340
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, fmt.Sprintf("Молния ударила на расстоянии %d метров", r))
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
