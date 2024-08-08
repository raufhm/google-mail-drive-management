package helper

import (
	"encoding/base64"
	"fmt"
	"github.com/charmbracelet/log"
	"google.golang.org/api/gmail/v1"
	"io"
	"net/http"
	"os"
)

func SaveEmail(msg *gmail.Message, mId string) {
	decodedData, _ := base64.URLEncoding.DecodeString(msg.Raw)
	file := "output/gmail/" + mId + ".eml"
	err := os.WriteFile(file, decodedData, os.ModePerm)
	if err != nil {
		log.Error(err.Error())
	}
}

func SaveFile(fileName string, resp *http.Response) {
	outFile, err := os.Create(fmt.Sprintf("output/gdrive/%s", fileName))
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			return
		}
	}(outFile)

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Fatalf("Unable to write file: %v", err)
	}

	fmt.Println("File downloaded successfully")
}
