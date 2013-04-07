// Diablo 3 Profiles in Go by mmorrell
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Account struct {
	BattleTag, BattleTagName, BattleTagNumber string
	Heroes                                    []Hero
	TimePlayed                                map[string]float64
	Kills                                     map[string]float64
}

type Hero struct {
	Name, Class                                      string
	Level, ParagonLevel, Gender                      int
	HeroId                                           int64 `json:"id"`
	Dead, Hardcore                                   bool
	Skills, Items, Followers, Stats, Kills, Progress interface{}
}

func main() {
	var d3Account Account = Account{BattleTagName: "Eurodance", BattleTagNumber: "1289"}

	d3Account.GetBattletag()
	d3Account.ShowTimePlayed()
	d3Account.ShowKills()

	d3Account.GetHeroes()
}

func (acc *Account) GetBattletag() {
	resp, err := http.Get(fmt.Sprintf("http://us.battle.net/api/d3/profile/%s-%s/",
		strings.ToLower(acc.BattleTagName), strings.ToLower(acc.BattleTagNumber)))
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal(body, &acc)

	if error != nil {
		fmt.Println(error)
	}
}

func (acc *Account) ShowTimePlayed() {
	fmt.Println("Time Played:\n")
	for k, v := range acc.TimePlayed {
		fmt.Printf("%s %.2f %s \n", k, v*100, "%")
	}
	fmt.Println()
}

func (acc *Account) ShowKills() {
	fmt.Println("Kills:\n")
	for k, v := range acc.Kills {
		fmt.Println(k, v)

	}
	fmt.Println()
}

func (acc *Account) GetHeroes() {
	for _, v := range acc.Heroes {
		if v.HeroId != 0 {
			resp, err := http.Get(fmt.Sprintf("http://us.battle.net/api/d3/profile/%s-%s/hero/%v",
				acc.BattleTagName, acc.BattleTagNumber, v.HeroId))
			defer resp.Body.Close()

			if err != nil {
				fmt.Println("Error:", err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			error := json.Unmarshal(body, &v)

			if error != nil {
				fmt.Println("Error:", error)
			}

			fmt.Printf("\n%s(%d)\n", v.Name, v.Level)

			for k, v2 := range v.Items.(map[string]interface{}) {
				fmt.Println(k+":", v2.(map[string]interface{})["name"])
			}

		}
	}
}
