package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Webhooksstart(bot *tgbotapi.BotAPI) {
	wh, _ := tgbotapi.NewWebhook(os.Getenv("NGROK-URL") + "/" + os.Getenv("BOT-TOKEN"))
	_, err := bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
}
