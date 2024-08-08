package gmailRepo

import (
	"google.golang.org/api/gmail/v1"
	"log"
)

type Handler struct {
	service *gmail.Service
}

func NewHandler(service *gmail.Service) *Handler {
	return &Handler{service: service}
}

// GetMessages is getting email from the given account
func (h *Handler) GetMessages(accountInfo, query string) []*gmail.ListMessagesResponse {
	user := accountInfo
	var messages []*gmail.ListMessagesResponse
	var nextPageToken string
	for {
		req := h.service.Users.Messages.List(user).Q(query)
		if nextPageToken != "" {
			req.PageToken(nextPageToken)
		}

		resp, err := req.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve messages: %v", err)
		}

		messages = append(messages, resp)

		// Check if there are more pages of results
		// If there are no more pages, break out of the loop
		// and return the results
		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}

	return messages
}

func (h *Handler) GetMessageDetails(messageId, email string) *gmail.Message {
	messages, err := h.service.Users.Messages.Get(email, messageId).Format("RAW").Do()
	if err != nil {
		log.Fatalf("Failed to list Gmail messages: %v", err)
	}
	return messages
}

func (h *Handler) DeleteMessage(messageId, email string) error {
	err := h.service.Users.Messages.Delete(email, messageId).Do()
	if err != nil {
		log.Fatalf("Failed to list Gmail messages: %v", err)
	}
	return err
}
