package worl_tour

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gobot/internal/consts"
	"math/rand"
	"reflect"
	"time"
)

func WorldHandler(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	keys := reflect.ValueOf(consts.World_tour).MapKeys()
	key := fmt.Sprintf("%s", keys[r.Intn(len(keys))])
	file := tgbotapi.FileID(fmt.Sprintf("%s", consts.World_tour[key]))
	m := tgbotapi.NewPhoto(u.Message.Chat.ID, file)
	m.Caption = fmt.Sprintf("Вы попали %s", key)
	bot.Send(m)
}
