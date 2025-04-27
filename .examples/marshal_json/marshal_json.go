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
	bot := Bot{
		Username:      "SpaceAI",
		Discriminator: 6555,
		ID:            snowflake.Snowflake(854409704704573450),
	}

	var response []byte
	response, _ = json.Marshal(bot)

	fmt.Println(string(response))
}
