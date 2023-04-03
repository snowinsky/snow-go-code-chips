package api

import (
	"log"
)
import "github.com/shopspring/decimal"

func DecimalDemo() {
	num1, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	num2, err := decimal.NewFromString("10")
	if err != nil {
		panic(err)
	}

	log.Println(num1.Add(num2))
	log.Println(num1.Sub(num2))
	log.Println(num1.Mul(num2))
	log.Println(num1.Div(num2))
}
