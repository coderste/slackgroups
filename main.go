package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_CLIENT_TOKEN"), slack.OptionDebug(true))

	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}
