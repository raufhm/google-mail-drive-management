package authRepo

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"time"
)

type Authenticator struct {
	config *oauth2.Config
	ctx    context.Context
}

func NewAuthenticator(credentials []byte, scopes ...string) *Authenticator {
	config, err := google.ConfigFromJSON(credentials, scopes...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return &Authenticator{
		config: config,
		ctx:    context.Background(),
	}
}

func (a *Authenticator) GetClient(tokenFile string) *http.Client {
	tok, err := a.tokenFromFile(tokenFile)
	if err != nil {
		tok = a.getTokenFromWeb()
		a.saveToken(tokenFile, tok)
	} else if tok.Expiry.Before(time.Now()) {
		tok = a.getTokenFromWeb()
		a.saveToken(tokenFile, tok)
	}
	return a.config.Client(a.ctx, tok)
}

func (a *Authenticator) GetContext() context.Context {
	return a.ctx
}

func (a *Authenticator) getTokenFromWeb() *oauth2.Token {
	authURL := a.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := a.config.Exchange(a.ctx, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func (a *Authenticator) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func (a *Authenticator) saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
