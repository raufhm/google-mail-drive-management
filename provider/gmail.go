package provider

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/raufhm/google-mail-drive-management/env"
	"github.com/raufhm/google-mail-drive-management/repo/authRepo"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"path/filepath"
)

func NewGmailService(authenticator *authRepo.Authenticator, config *env.Config, email string) (*gmail.Service, error) {
	tokenFile := filepath.Join(filepath.Dir(config.TokenFile), fmt.Sprintf("token_%s.json", email))
	client := authenticator.GetClient(tokenFile)
	return gmail.NewService(authenticator.GetContext(), option.WithHTTPClient(client))
}

func GetGmailService(email string) *gmail.Service {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	auth := NewAuth(config)
	service, err := NewGmailService(auth, config, email)
	if err != nil {
		log.Fatalf("Error loading gmail service: %v", err)
	}
	return service
}
