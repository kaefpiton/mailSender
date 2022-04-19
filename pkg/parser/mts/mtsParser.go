package mts

import (
	"errors"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"mailSender/pkg/parser"
	"strconv"
	"strings"
	"time"
)
//todo сделать интерфейс и убрать M

const websiteM = "https://www.mtsbank.ru/"

const emptyDataErrorM  = "Parsed data was empty"

var currentUSDCourseM = parser.Course{}
var currentEURCourseM = parser.Course{}

var usdErrorM error
var eurErrorM error


func GetUSD() (parser.Course,error) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{websiteM},
		ParseFunc: parceUSD,
	}).Start()

	return currentUSDCourseM, usdErrorM
}

func GetEUR() (parser.Course,error){
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{websiteM},
		ParseFunc: parceEUR,
	}).Start()

	return currentEURCourseM, eurErrorM
}


func parceUSD(g *geziyor.Geziyor, r *client.Response) {
	currentUSDCourseM.Bankname = "МТС Банк"
	datetime := time.Now()
	currentUSDCourseM.Date = datetime.Format("2006-01-02 15:04:05")

	var parsedUSDBuy = r.HTMLDoc.Find("#__next > div.styled__Container-sc-tzc9s2-0.hinLwq > div.styled__GridContainerDesktop-sc-tuvh3p-0.styled__ContentContainerDesktop-sc-tuvh3p-1.gaQver.esnlob > div.styled__Cell-sc-dvqv7w-0.YhGdp > div > div > div:nth-child(1) > div:nth-child(5)").Text()
	parsedUSDBuy = strings.Replace(parsedUSDBuy, ",", ".", 1)

	if buy, err := strconv.ParseFloat(parsedUSDBuy, 64); err == nil {
		currentUSDCourseM.Buy = buy
	}else{
		usdErrorM = errors.New(emptyDataErrorM)
	}

	var parsedUSDSell = r.HTMLDoc.Find("#__next > div.styled__Container-sc-tzc9s2-0.hinLwq > div.styled__GridContainerDesktop-sc-tuvh3p-0.styled__ContentContainerDesktop-sc-tuvh3p-1.gaQver.esnlob > div.styled__Cell-sc-dvqv7w-0.YhGdp > div > div > div:nth-child(1) > div:nth-child(6)").Text()
	parsedUSDSell = strings.Replace(parsedUSDSell, ",", ".", 1)
	if sell, err := strconv.ParseFloat(parsedUSDSell, 64); err == nil {
		currentUSDCourseM.Sell = sell
	}else{
		usdErrorM = errors.New(emptyDataErrorM)
	}

}

func parceEUR(g *geziyor.Geziyor, r *client.Response) {
	currentEURCourseM.Bankname = "МТС Банк"
	datetime := time.Now()
	currentEURCourseM.Date = datetime.Format("2006-01-02 15:04:05")

	var parsedEURBuy = r.HTMLDoc.Find("#__next > div.styled__Container-sc-tzc9s2-0.hinLwq > div.styled__GridContainerDesktop-sc-tuvh3p-0.styled__ContentContainerDesktop-sc-tuvh3p-1.gaQver.esnlob > div.styled__Cell-sc-dvqv7w-0.YhGdp > div > div > div:nth-child(2) > div:nth-child(5)").Text()
	parsedEURBuy = strings.Replace(parsedEURBuy, ",", ".", 1)

	if buy, err := strconv.ParseFloat(parsedEURBuy, 64); err == nil {
		currentEURCourseM.Buy = buy
	}else{
		usdErrorM = errors.New(emptyDataErrorM)
	}

	var parsedEURSell = r.HTMLDoc.Find("#__next > div.styled__Container-sc-tzc9s2-0.hinLwq > div.styled__GridContainerDesktop-sc-tuvh3p-0.styled__ContentContainerDesktop-sc-tuvh3p-1.gaQver.esnlob > div.styled__Cell-sc-dvqv7w-0.YhGdp > div > div > div:nth-child(2) > div:nth-child(6)").Text()
	parsedEURSell = strings.Replace(parsedEURSell, ",", ".", 1)

	if sell, err := strconv.ParseFloat(parsedEURSell, 64); err == nil {
		currentEURCourseM.Sell = sell
	}else{
		usdErrorM = errors.New(emptyDataErrorM)
	}
}
