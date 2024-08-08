package provider

import (
	"github.com/charmbracelet/log"
	"github.com/raufhm/google-mail-drive-management/env"
	"github.com/raufhm/google-mail-drive-management/repo/authRepo"
	"os"
)

func NewConfig() (*env.Config, error) {
	return env.NewLoadConfig()
}

func NewAuth(config *env.Config) *authRepo.Authenticator {
	credentials, err := os.ReadFile(config.CredentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	return authRepo.NewAuthenticator(credentials, config.Scopes...)
}
