package utils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sort"
	"telegram_bot/confis"
)

type send struct {
	name   string
	number int64
}

func Start() {
	//每天0点重置一些统计信息或同步数据到后台
	c := cron.New()

	c.AddFunc("@every 2m", func() {

		number := make([]send, 0)
		for _, info := range confis.GetServer() {
			n := Query(info[0], info[1])
			if n > 0 {
				number = append(number, send{
					name:   info[1],
					number: n,
				})
			}
		}

		sort.Slice(number, func(i, j int) bool {
			return number[i].number >= number[j].number
		})

		var context string
		for _, value := range number {
			context += fmt.Sprintf("%v：%v\n", value.name, value.number)
		}
		SendTG(context)
	})

	c.Start()
}
