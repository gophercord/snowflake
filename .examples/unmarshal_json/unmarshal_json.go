package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercord/snowflake"
)

type Bot struct {
	Username      string              `json:"username"`
	Discriminator uint16              `json:"discriminator"`
	ID            snowflake.Snowflake `json:"id"`
}

func main() {
	response := `{
		"username": "SpaceAI",
		"discriminator": 6555,
		"id": "854409704704573450"
	}`

	var bot Bot
	json.Unmarshal([]byte(response), &bot)

	fmt.Println(bot.ID)
	fmt.Println(bot.ID.Time())
}
