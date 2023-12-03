package horo

type days struct {
	Today      string `xml:"today"`
	Tomorrow   string `xml:"tomorrow"`
	Tomorrow02 string `xml:"tomorrow02"`
}

type Horo struct {
	Aries       days `xml:"aries"`
	Taurus      days `xml:"taurus"`
	Gemini      days `xml:"gemini"`
	Cancer      days `xml:"cancer"`
	Leo         days `xml:"leo"`
	Virgo       days `xml:"virgo"`
	Libra       days `xml:"libra"`
	Scorpio     days `xml:"scorpio"`
	Sagittarius days `xml:"sagittarius"`
	Capricorn   days `xml:"capricorn"`
	Aquarius    days `xml:"aquarius"`
	Pisces      days `xml:"pisces"`
}
