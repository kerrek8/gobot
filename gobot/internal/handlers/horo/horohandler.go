package horo

import (
	"encoding/xml"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gobot/internal/consts"
	"gobot/internal/keyboards/reply"
	"gobot/internal/structures/horo"
	"io"
	"log"
	"net/http"
	"reflect"
)

var choose string
var day string

func Horo(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите знак зодиака")
	msg.ReplyMarkup = reply.Horokeyboard1
	bot.Send(msg)
}

func Horohandler(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	r := []rune(u.Message.Text)
	if 9800 <= r[1] && r[1] <= 9811 {
		horoDays(string(r[1]), bot, u)
	} else {
		day = string(r[1 : len(r)-1])

		resp, err := http.Get("https://ignio.com/r/export/utf/xml/daily/com.xml")
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var result horo.Horo
		err = xml.Unmarshal(body, &result)
		if err != nil {
			log.Fatalln(err)
		}
		s := ""
		s += fmt.Sprintf("Гороскоп для %s на %s:\n", choose, day)
		a := reflect.ValueOf(result).FieldByName(consts.Transcript_dict[choose]).FieldByName(consts.Transcript_dict[day])
		s += fmt.Sprintf("%s", a.Interface())
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}

func horoDays(r string, bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	choose = r
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите на какой день вы хотите гороскоп")
	msg.ReplyMarkup = reply.Horokeyboard2
	bot.Send(msg)
}

//func getresult() {
//
//}
