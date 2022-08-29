package main

import (
	"github.com/robfig/cron"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"strconv"
	"time"
	"wx_morning/service"
	"wx_morning/utools"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	utools.ViperInit()
	userIds := viper.GetStringSlice("user.userIds")
	templateId := viper.GetString("wx.templateId")
	appId := viper.GetString("wx.appId")
	appSecret := viper.GetString("wx.appSecret")
	city := viper.GetString("user.city")
	startData := viper.GetString("user.startData")
	birthdayType := viper.GetInt("user.birthdayType")
	birthday := viper.GetString("user.birthday")
	token := viper.GetString("wx.token")
	spec := viper.GetString("user.spec")

	c := cron.New()
	err := c.AddFunc(spec, func() {
		for _, userId := range userIds {
			wc := wechat.NewWechat()
			memory := cache.NewMemory()
			//redisOpts := &cache.RedisOpts{
			//	Host:        viper.GetString("redis.host"),
			//	Database:    0,
			//	MaxActive:   10,
			//	MaxIdle:     10,
			//	IdleTimeout: 60,
			//	Password:    viper.GetString("redis.passwd"),
			//}
			//redisCache := cache.NewRedis(context.Background(), redisOpts)

			oa := wc.GetOfficialAccount(&offConfig.Config{
				AppID:     appId,
				AppSecret: appSecret,
				Token:     token,
				Cache:     memory,
			})

			weather := service.GetWeather(city)
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
