package main

import (
	"fmt"

	"github.com/aaronArinder/recobot/config"
	recobot "github.com/aaronArinder/recobot/src"
)

var BotID string

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	recobot.Start()
	<-make(chan struct{})

	return
}
