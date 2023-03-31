package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hello_world/api"
	"time"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Time("sdafs{}", time.Now()).Send()

	//api.GormDemoMain()
	//api.ErrorsDemo()
	//api.ExcelDemo()
	//api.GinDemo()
	api.DigDemo()
}
