package main

import (
	"fmt"

	bot "github.com/GuruDev1736/Golang_Discord_Bot/Bot"
	config "github.com/GuruDev1736/Golang_Discord_Bot/Config"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return

}
