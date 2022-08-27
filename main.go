package main

import (
	"github.com/robfig/cron"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"log"
	"math/rand"
	"strconv"
	"time"
	"wx_morning/service"
	"wx_morning/utools"
)

func main() {
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

			weather := service.GetWeather("长沙")
			templateMsg := &message.TemplateMessage{
				ToUser:     userId,
				TemplateID: templateId,
				Color:      utools.GetColor(),
				Data: map[string]*message.TemplateDataItem{
					"city": {
						Value: city,
						Color: utools.GetColor(),
					},
					"data": {
						Value: time.Now().Format("2006年01月02日"),
						Color: utools.GetColor(),
					},
					"weather": {
						Value: weather.Weather,
						Color: utools.GetColor(),
					},
					"temperature": {
						Value: strconv.FormatFloat(weather.Temp, 'f', -1, 64),
						Color: utools.GetColor(),
					},
					"highest": {
						Value: strconv.FormatFloat(weather.High, 'f', -1, 64),
						Color: utools.GetColor(),
					},
					"lowest": {
						Value: strconv.FormatFloat(weather.Low, 'f', -1, 64),
						Color: utools.GetColor(),
					},
					"love_days": {
						Value: service.GetDay(utools.TimeStringToGoTime(startData)),
						Color: utools.GetColor(),
					},
					"birthday_left": {
						Value: service.GetBirthdayLeft(birthday, birthdayType),
						Color: utools.GetColor(),
					},
					"words": {
						Value: service.GetWords(),
						Color: utools.GetColor(),
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
