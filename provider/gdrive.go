package provider

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/raufhm/google-mail-drive-management/env"
	"github.com/raufhm/google-mail-drive-management/repo/authRepo"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"path/filepath"
)

func NewGDriveService(authenticator *authRepo.Authenticator, config *env.Config, email string) (*drive.Service, error) {
	tokenFile := filepath.Join(filepath.Dir(config.TokenFile), fmt.Sprintf("token_%s.json", email))
	client := authenticator.GetClient(tokenFile)
	return drive.NewService(context.Background(), option.WithHTTPClient(client))
}

func GetDriveService(email string) *drive.Service {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	auth := NewAuth(config)
	service, err := NewGDriveService(auth, config, email)
	if err != nil {
		log.Fatalf("Error loading gmail service: %v", err)
	}
	return service
}
