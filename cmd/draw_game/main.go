package main

import (
	"fmt"
	"games/app/cron/steam_game"
	"games/app/initialize"
	"games/config"
	"os"
	"time"
)

func main() {
	config.LoadConfig()
	initialize.AppInit()

	funcName := ""
	if len(os.Args) > 1 {
		funcName = string(os.Args[1])
	}

	switch funcName {
	case "gameDetail":
		gameDetailHandle()
	case "gameList":
		gameListHandle()
	default:
		defaultHandle()
	}
}

func defaultHandle() {
	fmt.Println("func = ''")

}

func gameDetailHandle() {
	fmt.Println("func = ameDetail")
	steamGame := steam_game.SteamGameHandle{}
	var status bool
	for {
		status = steamGame.ConsumerGameDetailQueue()
		if !status {
			time.Sleep(2 * time.Second)
		}
	}
}

func gameListHandle() {
	fmt.Println("func = ameList")

}
