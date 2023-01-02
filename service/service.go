package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wx_morning/models"
	"wx_morning/utools"
)

// GetWeather 获取城市天气
func GetWeather(city string) models.DayWeather {
	var cityData models.DayWeather
	client := &http.Client{}
	weatherUrl := "http://autodev.openspeech.cn/csp/api/v2.1/weather?openId=aiuicus&clientType=android&sign=android&city=" + city
	request, err := http.NewRequest("GET", weatherUrl, nil)
	if err != nil {
		log.Println("create new request err:", err)
		return cityData
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return cityData
	}
	defer response.Body.Close()

	readAll, err := io.ReadAll(response.Body)
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

	var wd models.CityWeatherData
	err = json.Unmarshal(readAll, &wd)
	if err != nil {
		log.Println(err)
		return cityData
	}
	return wd.Data.List[0]
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
		next = utools.TimeStringToGoTime(solarStr[0] + "-" + solarStr[1] + "-" + solarStr[2])
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
			next = utools.TimeStringToGoTime(solarStr[0] + "-" + solarStr[1] + "-" + solarStr[2])
		}
	} else {
		next = utools.TimeStringToGoTime(strconv.Itoa(time.Now().Year()) + "-" + date)
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
	wordsUrl := "https://api.lovelive.tools/api/SweetNothings"
	request, err := http.NewRequest("GET", wordsUrl, nil)
	if err != nil {
		log.Fatal("create new request err:", err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer response.Body.Close()

	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	//var wordsData models.Words
	//err = json.Unmarshal(readAll, &wordsData)
	//if err != nil {
	//	log.Println(err)
	//	return ""
	//}
	return string(readAll)
}
