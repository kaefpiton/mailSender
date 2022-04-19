package parser

import (
	"errors"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"strconv"
	"strings"
)

const website = "https://mtcbank-adresa.ru/exchange.html"

const emptyDataError  = "Parsed data was empty"

var currentUSDCourse = Course{}
var currentEURCourse = Course{}

var usdError error
var eurError error


func GetUSD() (Course,error) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{website},
		ParseFunc: parceUSD,
	}).Start()

	return currentUSDCourse,usdError
}

func GetEUR() (Course,error){
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{website},
		ParseFunc: parceEUR,
	}).Start()

	return currentEURCourse,eurError
}



func parceUSD(g *geziyor.Geziyor, r *client.Response) {
	var data = []string{}
	var parsedUSD = strings.Split(r.HTMLDoc.Find("table.table.table-striped.respo.usdk").Find("tbody").Text(), "\n")

	//< 1 так как при сплите у нас выдает единицу (стоит быть аккуратней, где действительно один элемент)
	if len(parsedUSD) > 1 {
		for _, val := range(parsedUSD){
			if val != ""{
				data = append(data, val)
			}
		}

		currentUSDCourse.Bankname = data[0]

		if sell, err := strconv.ParseFloat(data[1], 64); err == nil {
			currentUSDCourse.Sell = sell
		}
		if buy, err := strconv.ParseFloat(data[2], 64); err == nil {
			currentUSDCourse.Buy = buy
		}

		currentUSDCourse.Date = data[3]

	}else {
		usdError = errors.New(emptyDataError)
	}
}

func parceEUR(g *geziyor.Geziyor, r *client.Response) {
	var data = []string{}
	var parsedEUR = strings.Split(r.HTMLDoc.Find("table.table.table-striped.respo.eurk").Find("tbody").Text(), "\n")

	//< 1 так как при сплите у нас выдает единицу (стоит быть аккуратней, где действительно один элемент)
	if len(parsedEUR) > 1 {
		for _, val := range(parsedEUR){
			if val != ""{
				data = append(data, val)
			}
		}

		currentEURCourse.Bankname = data[0]

		if sell, err := strconv.ParseFloat(data[1], 64); err == nil {
			currentEURCourse.Sell= sell
		}
		if buy, err := strconv.ParseFloat(data[2], 64); err == nil {
			currentEURCourse.Buy = buy
		}

		currentEURCourse.Date = data[3]

	}else {
		eurError = errors.New(emptyDataError)
	}
}






