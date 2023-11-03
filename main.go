/**
 * @Author: Cc
 * @Description: 描述
 * @File: main
 * @Version: 1.0.0
 * @Date: 2022/10/12 17:09
 * @Software : GoLand
 */

package main

import (
	"fmt"
	"log"
	"os"
	"telegram_bot/confis"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	confis.InitConfigStart()
}

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("robotToken"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	gameInfo := confis.GetGameAll()
	updates.Clear()
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Println(update.Message)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyToMessageID = update.Message.MessageID
			switch update.Message.Text {
			case "/start":
				msg.Text = "/help"
			case "/help":

				var msgText string
				for _, value := range confis.GetBotFunItem() {
					msgText += fmt.Sprintf("%v \n", value)
				}
				msg.Text = msgText
			case "/AllGame":

				var msgText = "GameAll \n"

				for _, value := range gameInfo[0] {
					msgText += fmt.Sprintf("%v\n", value)
				}

				msg.Text = msgText
			case "/Thailand":
				var msgText = "GameAll \n"
				for _, value := range gameInfo[1] {
					msgText += fmt.Sprintf("%v\n", value)
				}
				msg.Text = msgText

			case "/Indonesia":
				var msgText = "GameAll \n"
				for _, value := range gameInfo[2] {
					msgText += fmt.Sprintf("%v\n", value)
				}
				msg.Text = msgText
			////////// 泰国

			case "/Dummy":
			case "/SanGong":
			case "/Kaeng":

			/////////印尼
			case "/Domino":
			case "/RoyalWin":
			case "/BandarQQ":
				updates.Clear()
			case "/QiuQiu":
			case "/Texas":
			case "/Remi":

				//default:
				//	msg.Text = "暂时还不知道怎么回答您可以查看帮助 \n/help"
				//}

				send, errSend := bot.Send(msg)
				if errSend != nil {
					log.Println(errSend)
				}
				log.Println("ok:", send)
			}
		}
	}
}

/**

docker build -t repo.rdvalidasi.com/indonesia/game/chess/bandarqq:sendTG .
docker push repo.rdvalidasi.com/indonesia/game/chess/bandarqq:sendTG

cd /home/save/lcc/sendTG/
docker-compose down
docker rmi repo.rdvalidasi.com/indonesia/game/chess/bandarqq:sendTG
docker-compose up -d
*/
