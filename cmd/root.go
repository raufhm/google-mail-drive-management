package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "googlecli",
	Short: "Gmail & Gdrive CLI",
	Long:  `A CLI tool to clean up and download Gmail/Gdrive content based on filters such as age, size, and type.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("email", "e", "", "Email address for authentication")
}
