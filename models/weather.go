package models

// CityWeatherData 城市数据
type CityWeatherData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total      int          `json:"total"`
		SourceName string       `json:"sourceName"`
		List       []DayWeather `json:"list"`
		LogoURL    string       `json:"logoUrl"`
	} `json:"data"`
}

// DayWeather 某日天气数据
type DayWeather struct {
	City           string  `json:"city"`
	LastUpdateTime string  `json:"lastUpdateTime"`
	Date           string  `json:"date"`
	Weather        string  `json:"weather"`
	Temp           float64 `json:"temp,omitempty"`
	Humidity       string  `json:"humidity"`
	Wind           string  `json:"wind"`
	Pm25           float64 `json:"pm25"`
	Pm10           float64 `json:"pm10,omitempty"`
	Low            float64 `json:"low"`
	High           float64 `json:"high"`
	AirData        string  `json:"airData"`
	AirQuality     string  `json:"airQuality"`
	DateLong       int64   `json:"dateLong"`
	WeatherType    int     `json:"weatherType"`
	WindLevel      int     `json:"windLevel"`
	Province       string  `json:"province"`
}
