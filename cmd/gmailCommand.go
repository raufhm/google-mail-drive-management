package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/raufhm/google-mail-drive-management/helper"
	"github.com/raufhm/google-mail-drive-management/provider"
	"github.com/raufhm/google-mail-drive-management/repo/gmailRepo"
	"github.com/spf13/cobra"
)

func getGmailContent(cmd *cobra.Command, args []string) {
	gmailer := provider.NewGmailArgs()
	email := gmailer.GetEmail(cmd)
	accountInfo := gmailer.GetAccountInfo(cmd)
	query := gmailer.GetQueryGmail(cmd)

	service := provider.GetGmailService(email)
	handler := gmailRepo.NewHandler(service)

	messagesResponse := handler.GetMessages(accountInfo, query)
	if messagesResponse != nil {
		for _, message := range messagesResponse {
			if message.Messages != nil && len(message.Messages) > 0 {
				for _, m := range message.Messages {
					// get a RAW message
					msg := handler.GetMessageDetails(m.Id, email)
					// if download args is specified, then download email
					if gmailer.GetDownload(cmd) {
						helper.SaveEmail(msg, m.Id)
					}
					// if purge args is specified, then purge
					if gmailer.GetPurge(cmd) {
						if err := handler.DeleteMessage(m.Id, email); err != nil {
							log.Errorf("error deleting message: %s, account: %s", m.Id, email)
						}
					}
				}
			}
		}
	}
}
