package consts

// ngrok http --domain=blessed-expert-parakeet.ngrok-free.app 80

//youtubekey = "AIzaSyDMeRhspy2tqx3oepQQ3fVTtmLF-UK7yrY"

var World_tour = map[string]string{
	"в Болгарию": "AgACAgIAAxkBAAIEm2VkZAABIaUNvOUIXgF7bIx8jXUmzAACM9ExG94RIEtoSJhkEiDS9AEAAwIAA3MAAzME",
	"в Испанию":  "AgACAgIAAxkBAAIEmWVkY8-KMUpH1XecNcrG4kLwsj0aAAIw0TEb3hEgS_DtaCpesBvLAQADAgADcwADMwQ",
	"на Кубу":    "AgACAgIAAxkBAAIEn2VkZLtIhPeQgiwYMZAgGI6OfHgHAAI70TEb3hEgS1TZw6KLFPR0AQADAgADcwADMwQ",
	"в Кипр":     "AgACAgIAAxkBAAIEmmVkY-rQBC5yQGPe5Kdkgakc4aQzAAIx0TEb3hEgS6Af3-AiiDaPAQADAgADcwADMwQ",
	"в Таиланд":  "AgACAgIAAxkBAAIEo2VkZvS-yX3ob9SaQoxSw5VCdl3mAAJG0TEb3hEgS-7LKamvAvg-AQADAgADcwADMwQ",
	"во Вьетнам": "AgACAgIAAxkBAAIEnmVkZKckwwFyhuCHGXW-WnRMrYtCAAI60TEb3hEgS-5HzHLxCfLcAQADAgADcwADMwQ",
	"в Индию":    "AgACAgIAAxkBAAIEnWVkZJUzwWyYyGdYhnELgoffNc1yAAI50TEb3hEgSwTYKv2JGxkUAQADAgADcwADMwQ",
	"в Тунис":    "AgACAgIAAxkBAAIEnGVkZINhYbdlpgk4h2iGPvOb7myZAAI40TEb3hEgS3QJFVsWvzWpAQADAgADcwADMwQ",
	"в Армению":  "AgACAgIAAxkBAAIEmGVkY7WOizS9i_UnrzKHHvAVmvPMAAIu0TEb3hEgSwiuKTBbVAYbAQADAgADcwADMwQ",
	"в Египет":   "AgACAgIAAxkBAAIEoGVkZM52GbawfInvx0OA-LI60JuqAAI80TEb3hEgSxDIyiU95iw3AQADAgADcwADMwQ",
	"в Турцию":   "AgACAgIAAxkBAAIEoWVkZOcCi8OVvbs1h9IpHL-bIXr2AAI90TEb3hEgSzZ58Zys_RnkAQADAgADcwADMwQ",
	"в Россию":   "AgACAgIAAxkBAAIEomVkZP18kBq7q-xvlEywvGB_Rz1TAAI-0TEb3hEgS-sWRzINA4h9AQADAgADcwADMwQ",
	"в Дубай":    "AgACAgIAAxkBAAIEkmVkX8vZ3mtgXjfLJnXJ1xEN0S0fAALy0DEb3hEgS1dFvgnzkKD7AQADAgADcwADMwQ",
}

const (
	TymenLat = 57.1522
	TymenLon = 65.5272
	StrLat   = 60.72411
	StrLon   = 77.58138
)

const WeatherEndpoinCurrent = `https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric&lang=ru`
const WeatherEndpoinDays = `https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&cnt=%d&appid=%s&units=metric&lang=ru`

var Transcript_dict = map[string]string{
	"♉":           "Taurus",
	"♈":           "Aries",
	"♎":           "Libra",
	"♑":           "Capricorn",
	"♏":           "Scorpio",
	"♒":           "Aquarius",
	"♓":           "Pisces",
	"♍":           "Virgo",
	"♌":           "Leo",
	"♊":           "Gemini",
	"♐":           "Sagittarius",
	"♋":           "Cancer",
	"Сегодня":     "Today",
	"Завтра":      "Tomorrow",
	"Послезавтра": "Tomorrow02",
}
