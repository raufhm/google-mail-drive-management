package cmd

import (
	"github.com/raufhm/google-mail-drive-management/provider"
	"github.com/raufhm/google-mail-drive-management/repo/gdriveRepo"
	"github.com/spf13/cobra"
)

func getGDriveContent(cmd *cobra.Command, args []string) {
	gdriver := provider.NewGDriveArgs()
	email := gdriver.GetEmail(cmd)
	query := gdriver.GetQueryGDrive(cmd)

	service := provider.GetDriveService(email)
	handler := gdriveRepo.NewHandler(service)

	fileResponse := handler.ListFiles("", query)
	if fileResponse != nil {
		for _, fileList := range fileResponse {
			for _, f := range fileList.Files {
				file := handler.GetFileDetails(f.Id)
				if file != nil {
					if gdriver.GetDownload(cmd) {
						_ = handler.DownloadFile(file.Id, file.Name)

					}
					if gdriver.GetPurge(cmd) {
						//_ = handler.DeleteFile(file.Id)
					}
				}
			}
		}
	}
}
