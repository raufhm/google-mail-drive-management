package gdriveRepo

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/raufhm/google-mail-drive-management/helper"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"net/http"
)

type Handler struct {
	service *drive.Service
}

func NewHandler(service *drive.Service) *Handler {
	return &Handler{service: service}
}

// ListFiles lists files in the user's Google Drive account based on the given query
func (h *Handler) ListFiles(folderId, query string) []*drive.FileList {
	var files []*drive.FileList
	//var nextPageToken string
	for {
		req := h.service.Files.List().Q(query)
		//if nextPageToken != "" {
		//	req.PageToken(nextPageToken)
		//}

		resp, err := req.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		files = append(files, resp)
		break

		// Check if there are more pages of results
		// If there are no more pages, break out of the loop
		// and return the results
		//if resp.NextPageToken == "" {
		//	break
		//}
		//nextPageToken = resp.NextPageToken
	}

	return files
}

// GetFileDetails gets the details of a specific file in the user's Google Drive account
func (h *Handler) GetFileDetails(fileId string) *drive.File {
	file, err := h.service.Files.Get(fileId).Do()
	if err != nil {
		log.Fatalf("Failed to retrieve file details: %v", err)
	}
	return file
}

func (h *Handler) DownloadFile(fileId, fileName string) *http.Response {
	deferFunc := func(resp *http.Response) {
		if resp != nil && resp.Body != nil {
			err := resp.Body.Close()
			if err != nil {
				return
			}
		}
	}
	resp, err := h.service.Files.Get(fileId).Download()
	if err != nil {
		var gErr *googleapi.Error
		if errors.As(err, &gErr) && gErr.Code == 403 {
			// If the resp is not directly downloadable, try exporting it
			log.Info("File not directly downloadable, attempting to export...")
			exportMimeType := "application/pdf"
			exportResp, err := h.service.Files.Export(fileId, exportMimeType).Download()
			if err != nil {
				log.Fatalf("Unable to export resp: %v", err)
			}
			helper.SaveFile(fmt.Sprintf("%s.pdf", fileName), exportResp)
			deferFunc(exportResp)
		}
	} else {
		helper.SaveFile(fileName, resp)
		deferFunc(resp)
	}

	return resp
}

// DeleteFile deletes a specific file from the user's Google Drive account
func (h *Handler) DeleteFile(fileId string) error {
	err := h.service.Files.Delete(fileId).Do()
	if err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	return err
}
