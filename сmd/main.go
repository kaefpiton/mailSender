package main

import (
	"fmt"
	sender "mailSender/pkg/mailsend"
	"mailSender/pkg/parser/mts"
	"mailSender/сmd/config"
	"time"
)

func main() {

	myconfig := config.LoadConfig("C:/Projects/Go/src/mailSender/сmd/config/config.json")

	adresses := []string{"dimafed7@gmail.com"}
							//"1991goliaf@gmail.com"}

	//Порог еврового курса
	var thresholdSellEur float64 = 0

	for{
		usd, usderr := mts.GetUSD()
		eur, eurerr := mts.GetEUR()

		if usderr == nil && eurerr == nil{
			if eur.Sell > thresholdSellEur{
				fmt.Printf("Текущий курс продажи %v\n", thresholdSellEur)
				err := sender.SendCourse(myconfig,adresses,usd,eur)
				if err != nil {
					fmt.Println(err)
					break
				}
				thresholdSellEur = eur.Sell
			}

		}else {
			err := sender.SendError(myconfig,adresses,"Произошла ошибка при парсинге данных!")
			if err != nil {
				fmt.Println(err)
			}
			break
		}

		time.Sleep(time.Minute * 5)
	}


}


