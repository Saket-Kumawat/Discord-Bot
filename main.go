package main

import (
	"fmt"

	bot "github.com/Saket-Kumawat/Discord-Bot/bot"
	"github.com/Saket-Kumawat/Discord-Bot/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{}) // block forever

	return

}
