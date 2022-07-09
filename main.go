package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal("error loading .env file")
	}
	//get the slackbottoken from .env file
	slackBotToken, exists := os.LookupEnv("SLACK_BOT_OAUTH-TOKEN")
	//get the channelId from .env file
	channelId, isAvail := os.LookupEnv("CHANNEL_ID")
	//check if slackBottoken is available
	if !exists {
		log.Fatal("no env variable found for the SLACK_BOT_OAUTH-TOKEN")
	}
	//check if slackBottoken is available
	if !isAvail {
		log.Fatal("no env variable found for the CHANNEL_ID")
	}
	//builds a slack client from the provided token
	api := slack.New(slackBotToken)
	//slice of all channels
	channelsSlice := []string{channelId}
	//slice of all files
	fileSlice := []string{"WILLIE_resumer_max.docx"}

	for i := 0; i < len(fileSlice); i++ {
		// contains all the parameters necessary for a file upload request
		params := slack.FileUploadParameters{
			File:     fileSlice[i],
			Channels: channelsSlice,
		}

		//upload file
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Name: %s, url : %s\n", file.Name, file.URL)
	}
}
