package api

import (
	"github.com/jinzhu/now"
	"log"
	"time"
)

func NowDemo() {
	t := now.With(time.Now()).BeginningOfHour()
	log.Println(t)
}
