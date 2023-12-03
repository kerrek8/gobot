package owm_days

type OwmDays struct {
	Cnt  int    `json:"cnt"`
	List []list `json:"list"`
}
type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}
type weather struct {
	Description string `json:"description"`
}
type clouds struct {
	All int `json:"all"`
}
type list struct {
	Main    main      `json:"main"`
	Weather []weather `json:"weather"`
	Clouds  clouds    `json:"clouds"`
	DtTxt   string    `json:"dt_txt"`
}
