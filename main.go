package main

import (
	"encoding/json"
	"github.com/robfig/cron"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wx_morning/utools"
)

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

// DayWeather 天气数据
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

type Words struct {
	Data WordsData `json:"data"`
}

type WordsData struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// GetWeather 获取城市天气
func GetWeather(city string) DayWeather {
	var cityData DayWeather
	client := &http.Client{}
	weatherUrl := "http://autodev.openspeech.cn/csp/api/v2.1/weather?openId=aiuicus&clientType=android&sign=android&city=" + city
	request, err := http.NewRequest("GET", weatherUrl, nil)
	if err != nil {
		log.Fatal("create new request err:", err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return cityData
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return cityData
	}

	//file, err := os.Create("city.json")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer file.Close()
	//
	//_, err = file.WriteString(string(readAll))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	var wd CityWeatherData
	err = json.Unmarshal(readAll, &wd)
	if err != nil {
		log.Println(err)
		return cityData
	}
	return wd.Data.List[0]
}

// TimeStringToGoTime 字符串转时间
func TimeStringToGoTime(tm string) time.Time {
	t, err := time.ParseInLocation("2006-01-02", tm, time.Local)
	if nil == err && !t.IsZero() {
		return t
	}
	return time.Time{}
}

// GetDay 获取相恋天数
func GetDay(startTime time.Time) string {
	return strconv.Itoa(int(time.Now().Sub(startTime).Hours() / 24))
}

// GetBirthdayLeft 获取距离下次生日的时间
func GetBirthdayLeft(date string, birthdayType int) string {
	var next time.Time
	// 判断是农历还是公历
	if birthdayType == 1 {
		// 农历转阳历
		splitData := strings.Split(date, "-")
		month, _ := strconv.Atoi(splitData[0])
		day, _ := strconv.Atoi(splitData[1])
		solar := utools.ConvertLunarToSolar(time.Now().Year(), month, day, false)

		var solarStr [3]string
		for k, v := range solar {
			if v < 10 {
				solarStr[k] = "0" + strconv.Itoa(v)
			} else {
				solarStr[k] = strconv.Itoa(v)
			}
		}
		next = TimeStringToGoTime(solarStr[0] + "-" + solarStr[1] + "-" + solarStr[2])
		// 如果转换后的阳历时间小于当前时间，获取下一年的阳历时间
		if next.Sub(time.Now()).Hours()/24 < 0 {
			solar = utools.ConvertLunarToSolar(time.Now().Year()+1, month, day, false)
			for k, v := range solar {
				if v < 10 {
					solarStr[k] = "0" + strconv.Itoa(v)
				} else {
					solarStr[k] = strconv.Itoa(v)
				}
			}
			next = TimeStringToGoTime(solarStr[0] + "-" + solarStr[1] + "-" + solarStr[2])
		}
	} else {
		next = TimeStringToGoTime(strconv.Itoa(time.Now().Year()) + "-" + date)
	}

	subDay := next.Sub(time.Now()).Hours() / 24
	if subDay > 0 {
		return strconv.Itoa(int(subDay))
	}
	return strconv.Itoa(int(next.AddDate(1, 0, 0).Sub(time.Now()).Hours() / 24))
}

// GetWords 获取彩虹屁
func GetWords() string {
	client := &http.Client{}
	wordsUrl := "https://api.shadiao.pro/chp"
	request, err := http.NewRequest("GET", wordsUrl, nil)
	if err != nil {
		log.Fatal("create new request err:", err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ""
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	readAll, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	var wordsData Words
	err = json.Unmarshal(readAll, &wordsData)
	if err != nil {
		log.Println(err)
		return ""
	}
	return wordsData.Data.Text
}

// GetColor 生成随机颜色
func GetColor() string {
	colors := []byte("0123456789ABCDEF")

	color := "#"
	for i := 0; i < 6; i++ {
		color += string(colors[rand.Intn(16)])
	}
	return color
}

func main() {
	//userIds := strings.Split(os.Getenv("USER_ID"), "\n")
	//templateId := os.Getenv("TEMPLATE_ID")
	//appId := os.Getenv("APP_ID")
	//appSecret := os.Getenv("APP_SECRET")
	//city := os.Getenv("CITY")
	//startData := os.Getenv("START_DATE")
	//birthday := os.Getenv("BIRTHDAY")

	rand.Seed(time.Now().UnixNano())
	userIds := []string{"oAXn15v2cBLKKI_zhj2B8786uYJc"}
	templateId := "58cEgEMo30xWKivHUCQHjQIcgGmcE7ZFVpcWtqvZug0"
	appId := "wx190c5148b35ec8af"
	appSecret := "ff4aae1e42471d863d335dc4bc5d958a"
	city := "长沙"
	startData := "2022-07-09"
	birthdayType := 1 // 1: 农历   2: 阳历
	birthday := "07-15"

	c := cron.New()
	spec := "0 0 8 * * ?"
	err := c.AddFunc(spec, func() {
		for _, userId := range userIds {
			wc := wechat.NewWechat()
			memory := cache.NewMemory()

			oa := wc.GetOfficialAccount(&offConfig.Config{
				AppID:     appId,
				AppSecret: appSecret,
				Token:     "CATER123123",
				Cache:     memory,
			})

			weather := GetWeather("长沙")
			templateMsg := &message.TemplateMessage{
				ToUser:     userId,
				TemplateID: templateId,
				Color:      GetColor(),
				Data: map[string]*message.TemplateDataItem{
					"city": {
						Value: city,
						Color: GetColor(),
					},
					"data": {
						Value: time.Now().Format("2006年01月02日"),
						Color: GetColor(),
					},
					"weather": {
						Value: weather.Weather,
						Color: GetColor(),
					},
					"temperature": {
						Value: strconv.FormatFloat(weather.Temp, 'f', -1, 64),
						Color: GetColor(),
					},
					"highest": {
						Value: strconv.FormatFloat(weather.High, 'f', -1, 64),
						Color: GetColor(),
					},
					"lowest": {
						Value: strconv.FormatFloat(weather.Low, 'f', -1, 64),
						Color: GetColor(),
					},
					"love_days": {
						Value: GetDay(TimeStringToGoTime(startData)),
						Color: GetColor(),
					},
					"birthday_left": {
						Value: GetBirthdayLeft(birthday, birthdayType),
						Color: GetColor(),
					},
					"words": {
						Value: GetWords(),
						Color: GetColor(),
					},
				},
			}

			template := oa.GetTemplate()
			send, err := template.Send(templateMsg)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(send, " success")
		}
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	select {}
}
