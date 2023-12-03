package weather

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gobot/internal/consts"
	"gobot/internal/keyboards/reply"
	"gobot/internal/structures/owm_days"
	"gobot/internal/structures/owm_geo_cur"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

var lat float64
var lon float64
var delta int

func tymen(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите интересующую выас опцию")
	msg.ReplyMarkup = reply.Weatherkeyboard
	bot.Send(msg)
	lat = consts.TymenLat
	lon = consts.TymenLon
	delta = 0
}

func strej(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите интересующую выас опцию")
	msg.ReplyMarkup = reply.Weatherkeyboard
	bot.Send(msg)
	lat = consts.StrLat
	lon = consts.StrLon
	delta = 2
}

func now(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	url := fmt.Sprintf(consts.WeatherEndpoinCurrent, lat, lon, os.Getenv("OWM-API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result owm_geo_cur.OwmCurrent
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}
	temp := math.Round(result.Main.Temp + 1)
	feelslike := math.Round(result.Main.FeelsLike + 1)
	speed := math.Round(result.Wind.Speed)
	s := ""
	s += fmt.Sprintf("🌤ПОГОДА СЕЙЧАС🌤\n")
	s += fmt.Sprintf("<b>Температура: %.0f°C</b>\n", temp)
	s += fmt.Sprintf("Ощущается: %.0f°C\n", feelslike)
	s += fmt.Sprintf("Погода: %s\n", result.Weather[0].Description)
	s += fmt.Sprintf("Облачность: %d%%\n", result.Clouds.All)
	s += fmt.Sprintf("Скорость ветра: %.0fм/с\n", speed)
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func today(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	cnt := 8
	url := fmt.Sprintf(consts.WeatherEndpoinDays, lat, lon, cnt, os.Getenv("OWM-API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result owm_days.OwmDays
	err = json.Unmarshal(body, &result)
	r := result.List
	if err != nil {
		log.Fatalln(err)
	}
	s := ""
	s += fmt.Sprintf("🌤ПОГОДА НА СЕГОДНЯ🌤\n")
	s += fmt.Sprintf("🌤%s", r[0].DtTxt[8:10]) + "-" + fmt.Sprintf("%s", r[0].DtTxt[5:7]) + "-" + fmt.Sprintf("%s🌤", r[0].DtTxt[0:4])
	var listappend [][]string
	for i := 0; i < len(r); i++ {
		timebreak, _ := strconv.Atoi(r[i].DtTxt[11:13])
		timebreak += 3 + delta
		if timebreak == 3 || timebreak >= 25 {
			break
		}
		temp := strconv.Itoa(int(math.Round(r[i].Main.Temp + 1)))
		feelslike := strconv.Itoa(int(math.Round(r[i].Main.FeelsLike + 1)))
		cloudness := strconv.Itoa(r[i].Clouds.All)
		dtt := strconv.Itoa(timebreak)
		h := []string{dtt, temp, feelslike, r[i].Weather[0].Description, cloudness}
		listappend = append(listappend, h)

	}
	for _, v := range listappend {
		s += fmt.Sprintf("\nВ %s:00 \n", v[0])
		s += fmt.Sprintf("<b>Температура: %s\n</b>", v[1])
		s += fmt.Sprintf("Ощущается как: %s\n", v[2])
		s += fmt.Sprintf("Погода: %s\n", v[3])
		s += fmt.Sprintf("Облачность: %s%%\n", v[4])
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func tomor(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	cnt := 16
	url := fmt.Sprintf(consts.WeatherEndpoinDays, lat, lon, cnt, os.Getenv("OWM-API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result owm_days.OwmDays
	err = json.Unmarshal(body, &result)
	r := result.List
	if err != nil {
		log.Fatalln(err)
	}
	s := ""
	for {
		date, _ := strconv.Atoi(r[0].DtTxt[8:10])
		datenow, _ := strconv.Atoi(time.Now().Format("02"))
		if date == datenow {
			r = r[1:]
		} else {

			break
		}
	}

	s += fmt.Sprintf("🌤ПОГОДА НА ЗАВТРА🌤\n")
	s += fmt.Sprintf("🌤%s", r[0].DtTxt[8:10]) + "-" + fmt.Sprintf("%s", r[0].DtTxt[5:7]) + "-" + fmt.Sprintf("%s🌤", r[0].DtTxt[0:4])
	var listappend [][]string

	for i := 0; i < len(r); i++ {
		timebreak, _ := strconv.Atoi(r[i].DtTxt[11:13])
		timebreak += 3 + delta
		datenow, _ := strconv.Atoi(time.Now().Format("02"))
		date, _ := strconv.Atoi(r[i].DtTxt[8:10])
		if timebreak >= 25 || date == datenow+2 {
			break
		}
		temp := strconv.Itoa(int(math.Round(r[i].Main.Temp + 1)))
		feelslike := strconv.Itoa(int(math.Round(r[i].Main.FeelsLike + 1)))
		cloudness := strconv.Itoa(r[i].Clouds.All)
		dtt := strconv.Itoa(timebreak)
		h := []string{dtt, temp, feelslike, r[i].Weather[0].Description, cloudness}
		listappend = append(listappend, h)

	}

	for _, v := range listappend {
		s += fmt.Sprintf("\nВ %s:00 \n", v[0])
		s += fmt.Sprintf("<b>Температура: %s\n</b>", v[1])
		s += fmt.Sprintf("Ощущается как: %s\n", v[2])
		s += fmt.Sprintf("Погода: %s\n", v[3])
		s += fmt.Sprintf("Облачность: %s%%\n", v[4])
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func afterTomorrow(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	cnt := 35
	url := fmt.Sprintf(consts.WeatherEndpoinDays, lat, lon, cnt, os.Getenv("OWM-API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result owm_days.OwmDays
	err = json.Unmarshal(body, &result)
	r := result.List
	if err != nil {
		log.Fatalln(err)
	}
	s := ""
	for {
		date, _ := strconv.Atoi(r[0].DtTxt[8:10])
		datenow, _ := strconv.Atoi(time.Now().Format("02"))
		if date == datenow || date == datenow+1 {
			r = r[1:]
		} else {
			break
		}
	}

	s += fmt.Sprintf("🌤ПОГОДА НА ПОСЛЕЗАВТРА🌤\n")
	s += fmt.Sprintf("🌤%s", r[0].DtTxt[8:10]) + "-" + fmt.Sprintf("%s", r[0].DtTxt[5:7]) + "-" + fmt.Sprintf("%s🌤", r[0].DtTxt[0:4])
	var listappend [][]string

	for i := 0; i < len(r); i++ {
		timebreak, _ := strconv.Atoi(r[i].DtTxt[11:13])
		timebreak += 3 + delta
		datenow, _ := strconv.Atoi(time.Now().Format("02"))
		date, _ := strconv.Atoi(r[i].DtTxt[8:10])
		if timebreak >= 25 || date == datenow+3 {
			break
		}
		temp := strconv.Itoa(int(math.Round(r[i].Main.Temp + 1)))
		feelslike := strconv.Itoa(int(math.Round(r[i].Main.FeelsLike + 1)))
		cloudness := strconv.Itoa(r[i].Clouds.All)
		dtt := strconv.Itoa(timebreak)
		h := []string{dtt, temp, feelslike, r[i].Weather[0].Description, cloudness}
		listappend = append(listappend, h)

	}

	for _, v := range listappend {
		s += fmt.Sprintf("\nВ %s:00 \n", v[0])
		s += fmt.Sprintf("<b>Температура: %s\n</b>", v[1])
		s += fmt.Sprintf("Ощущается как: %s\n", v[2])
		s += fmt.Sprintf("Погода: %s\n", v[3])
		s += fmt.Sprintf("Облачность: %s%%\n", v[4])
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func fiveDays(bot *tgbotapi.BotAPI, u tgbotapi.Update) {
	cnt := 40
	url := fmt.Sprintf(consts.WeatherEndpoinDays, lat, lon, cnt, os.Getenv("OWM-API"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result owm_days.OwmDays
	err = json.Unmarshal(body, &result)
	r := result.List
	if err != nil {
		log.Fatalln(err)
	}
	s := ""
	s += fmt.Sprintf("🌤ПОГОДА НА 5 ДНЕЙ🌤\n")
	s += fmt.Sprintf("🌤%s", r[0].DtTxt[8:10]) + "-" + fmt.Sprintf("%s", r[0].DtTxt[5:7]) + "-" + fmt.Sprintf("%s🌤", r[0].DtTxt[0:4])
	var listappend [][]string

	for i := 0; i < len(r); i++ {
		timebreak, _ := strconv.Atoi(r[i].DtTxt[11:13])
		timebreak += 3 + delta
		//datenow, _ := strconv.Atoi(time.Now().Format("02"))
		//date, _ := strconv.Atoi(r[i].DtTxt[8:10])
		if timebreak >= 25 || timebreak == 3 {
			continue
		}
		d := fmt.Sprintf("%s-%s-%s", r[i].DtTxt[8:10], r[i].DtTxt[5:7], r[i].DtTxt[:4])
		temp := strconv.Itoa(int(math.Round(r[i].Main.Temp + 1)))
		feelslike := strconv.Itoa(int(math.Round(r[i].Main.FeelsLike + 1)))
		cloudness := strconv.Itoa(r[i].Clouds.All)
		dtt := strconv.Itoa(timebreak)
		h := []string{d, dtt, temp, feelslike, r[i].Weather[0].Description, cloudness}
		listappend = append(listappend, h)

	}

	for _, v := range listappend {
		if v[1] == "6" || v[1] == "5" {
			s += fmt.Sprintf("\n🌤%s🌤\n", v[0])
			s += fmt.Sprintf("\nВ %s:00 \n", v[1])
			s += fmt.Sprintf("<b>Температура: %s\n</b>", v[2])
			s += fmt.Sprintf("Ощущается как: %s\n", v[3])
			s += fmt.Sprintf("Погода: %s\n", v[4])
			s += fmt.Sprintf("Облачность: %s%%\n", v[5])
		} else {
			s += fmt.Sprintf("\nВ %s:00 \n", v[1])
			s += fmt.Sprintf("<b>Температура: %s\n</b>", v[2])
			s += fmt.Sprintf("Ощущается как: %s\n", v[3])
			s += fmt.Sprintf("Погода: %s\n", v[4])
			s += fmt.Sprintf("Облачность: %s%%\n", v[5])
		}

	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, s)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}
