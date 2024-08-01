package steam_game

import (
	"fmt"
)

type SteamGameHandle struct{}

func (s *SteamGameHandle) ConsumerGameDetailQueue() bool {
	var appId string
	var err error

	appId, err = s.getAppIdFromGameDetailQueue()
	if err != nil {
		fmt.Println("Error getting app id from game detail queue")
		return false
	}

	s.GetGameDetail(appId)

	return true
}

func (s *SteamGameHandle) getAppIdFromGameDetailQueue() (string, error) {
	return "", nil
}

func (s *SteamGameHandle) GetGameDetail(appId string) {
	fmt.Println("SteamGameHandle GetGameDetail Begin")
	fmt.Printf("appId = %s \n", appId)

	fmt.Println("SteamGameHandle GetGameDetail End")
}

func (s *SteamGameHandle) ConsumerGamePriceQueue() {

}
