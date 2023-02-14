package models

type CityWeatherData struct {
	CityDId    string       `json:"cityid"`
	City       string       `json:"city"`
	CityEn     string       `json:"cityEn"`
	Country    string       `json:"country"`
	CountryEn  string       `json:"countryEn"`
	UpdateTime string       `json:"update_time"`
	Data       []DayWeather `json:"data"`
	Nums       int          `json:"nums"`
}

// DayWeather 某日天气数据
type DayWeather struct {
	Day           string        `json:"day"`
	Date          string        `json:"date"`
	Week          string        `json:"week"`
	Wea           string        `json:"wea"`
	WeaImg        string        `json:"wea_img"`
	WeaDay        string        `json:"wea_day"`
	WeaDayImg     string        `json:"wea_day_img"`
	WeaNight      string        `json:"wea_night"`
	WeaNightImg   string        `json:"wea_night_img"`
	Tem           string        `json:"tem"`
	Tem1          string        `json:"tem1"`
	Tem2          string        `json:"tem2"`
	Humidity      string        `json:"humidity"`
	Visibility    string        `json:"visibility"`
	Pressure      string        `json:"pressure"`
	Win           []string      `json:"win"`
	WinSpeed      string        `json:"win_speed"`
	WinMeter      string        `json:"win_meter"`
	Sunrise       string        `json:"sunrise"`
	Sunset        string        `json:"sunset"`
	Air           string        `json:"air"`
	AirLevel      string        `json:"air_level"`
	AirTips       string        `json:"air_tips"`
	Phrase        string        `json:"phrase"`
	Narrative     string        `json:"narrative"`
	Moonrise      string        `json:"moonrise"`
	Moonset       string        `json:"moonset"`
	MoonPhrase    string        `json:"moonPhrase"`
	Rain          string        `json:"rain"`
	UvIndex       string        `json:"uvIndex"`
	UvDescription string        `json:"uvDescription"`
	Alarm         []interface{} `json:"alarm,omitempty"`
}
