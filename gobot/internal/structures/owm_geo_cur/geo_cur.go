package owm_geo_cur

type OwmCurrent struct {
	Weather    []weather `json:"weather"`
	Main       main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       wind      `json:"wind"`
	Clouds     clouds    `json:"clouds"`
}
type weather struct {
	Description string `json:"description"`
}
type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}
type wind struct {
	Speed float64 `json:"speed"`
}
type clouds struct {
	All int `json:"all"`
}
