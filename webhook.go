package main

import (
	"fmt"

	"github.com/gtuk/discordwebhook"
	"github.com/spf13/viper"
)

func SendWebhook(ServiceName string) {
	if !viper.GetBool("WebhookEnabled") {
		logger.Debug("Webhook is disabled, skipping...")
		return
	}
	var Title string = ":rotating_light: Alert! :rotating_light:"
	var Description string = fmt.Sprintf(":warning: Service `%s` on node `%s` is down!", ServiceName, viper.GetString("NodeName"))
	var Content string = fmt.Sprintf("<@&%s>", viper.GetString("RoleID"))
	var Color = "15548997"
	var Footer string = "github.com/yajtpg/ctlmon | " + Version

	embed := discordwebhook.Embed{
		Title:       &Title,
		Description: &Description,
		Color:       &Color,
		Footer:      &discordwebhook.Footer{Text: &Footer},
	}
	logger.Debug(Title, Description, Content, Color, Footer)
	err := discordwebhook.SendMessage(viper.GetString("WebhookURL"), discordwebhook.Message{
		Embeds:  &[]discordwebhook.Embed{embed},
		Content: &Content,
	})
	if err != nil {
		logger.Fatalf("Error while sending webhook: %s", err)
	}
	logger.Infof("Sent webhook for service: %s", ServiceName)
}
