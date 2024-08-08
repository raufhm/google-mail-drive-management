package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getGmailContentCmd)
	// getGmailContentCmd
	getGmailContentCmd.Flags().StringP("email", "e", "", "Email address for authentication")
	getGmailContentCmd.Flags().BoolP("download", "d", false, "Download flag to save email in local")
	getGmailContentCmd.Flags().BoolP("purge", "p", false, "Purge flag to remove email in gmail")
	getGmailContentCmd.Flags().StringP("size", "s", "", "Size flag to filter email with particular size")
	getGmailContentCmd.Flags().StringSliceP("range", "r", []string{},
		"Range flag to filter email with date range. \n"+
			"ex1: --range 0d,1d (read: older_than:0d newer_than:1d) or \n"+
			"ex2: --range 0d (read: older_than:0d)")
	getGmailContentCmd.Flags().StringP("category", "c", "primary", "Category flag to filter email in particular category")
	getGmailContentCmd.Flags().StringP("in", "i", "inbox", "In Folder flag to filter email in particular folder")
	getGmailContentCmd.Flags().StringP("filename", "f", "", "Filename flag to filter email with particular file name")

	err := getGmailContentCmd.MarkFlagRequired("email")
	if err != nil {
		return
	}

	rootCmd.AddCommand(getGDriveContentCmd)
	// getGDriveContentCmd
	getGDriveContentCmd.Flags().StringP("email", "e", "", "Email address for authentication")
	getGDriveContentCmd.Flags().BoolP("download", "d", false, "Download flag to save file in local")
	getGDriveContentCmd.Flags().BoolP("purge", "p", false, "Purge flag to remove file")
	getGDriveContentCmd.Flags().StringP("age", "a", "", "Age flag to filter file with older than")
	getGDriveContentCmd.Flags().StringP("size", "s", "1000", "Size flag to filter file with particular size")
	getGDriveContentCmd.Flags().StringP("text", "t", "", "Text flag to filter file with particular contains text")
	getGDriveContentCmd.Flags().StringP("filename", "f", "", "Filename flag to filter filename contains text")
	getGDriveContentCmd.Flags().StringP("mimeType", "m", "", "MimeType flag to filter mimeType contains")

	err = getGDriveContentCmd.MarkFlagRequired("email")
	if err != nil {
		return
	}

}

var getGmailContentCmd = &cobra.Command{
	Use:   "getGmailContent",
	Short: "Get Gmail content based on filters",
	Long:  `Fetch Gmail content based on provided filters such as age, size.`,
	Run:   getGmailContent,
}

var getGDriveContentCmd = &cobra.Command{
	Use:   "getGdriveContent",
	Short: "Get drive content based on filters",
	Long:  `Fetch Gdrive content based on provided filters such as age, size.`,
	Run:   getGDriveContent,
}
