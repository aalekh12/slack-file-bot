package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4584469582227-4598624339265-5yPVHXEzrYLFyXc2V27xlYHX")
	os.Setenv("CHANNEL_ID", "C04HW88L80G")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelarr := []string{os.Getenv("CHANNEL_ID")}
	filearr := []string{"main.pdf"}

	for i := 0; i < len(filearr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelarr,
			File:     filearr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("%s\n", err)
			return
		}
		fmt.Println("Name: %s URL %s \n", file.Name, file.URL)
	}
}
