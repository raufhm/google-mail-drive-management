package main

import (
	"github.com/raufhm/google-mail-drive-management/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
