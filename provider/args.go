package provider

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"strings"
)

type Commander interface {
	GetDownload(cmd *cobra.Command) bool
	GetPurge(cmd *cobra.Command) bool
	GetEmail(cmd *cobra.Command) string
	GetAccountInfo(cmd *cobra.Command) string
	GetQueryGmail(cmd *cobra.Command) string

	GetQueryGDrive(cmd *cobra.Command) string
}

func NewGmailArgs() *GMailObj {
	return &GMailObj{}
}

type GMailObj struct {
	Commander
}

func (a *GMailObj) GetDownload(cmd *cobra.Command) bool {
	return Download(cmd)
}

func (a *GMailObj) GetPurge(cmd *cobra.Command) bool {
	return Purge(cmd)
}

func (a *GMailObj) GetEmail(cmd *cobra.Command) string {
	return Email(cmd)
}

func (a *GMailObj) GetAccountInfo(cmd *cobra.Command) string {
	return AccountInfo(cmd)
}

func (a *GMailObj) GetQueryGmail(cmd *cobra.Command) string {
	return QueryGmail(cmd)
}

func Email(cmd *cobra.Command) string {
	email, _ := cmd.Flags().GetString("email")
	return email
}

func AccountInfo(cmd *cobra.Command) string {
	accountInfo, _ := cmd.Flags().GetString("account_info")
	return accountInfo
}

func Download(cmd *cobra.Command) bool {
	download, _ := cmd.Flags().GetBool("download")
	return download
}

func Purge(cmd *cobra.Command) bool {
	purge, _ := cmd.Flags().GetBool("purge")
	return purge
}

//{"Has Attachment", "Type 'yes' for filter", "has:attachment"},
//{"Google Drive/Docs Attachment", "Type 'yes' for filter", "has:drive"},
//{"Document Attachment", "Type 'yes' for filter", "has:document"},
//{"Filename", "Enter filename or file type", "filename:"},
//{"Category", "Enter category", "category:"},
//{"Size", "Enter size in bytes", "size:"},
//{"Newer than", "Enter duration (e.g., 2d, 3m, 1y)", "newer_than:"},
//{"In folder", "Enter folder name", "in:"},

func QueryGmail(cmd *cobra.Command) string {
	var query []string
	// size
	size, _ := cmd.Flags().GetString("size")
	if size != "" {
		query = append(query, "size:"+size)
	}

	// date range
	dateRange, _ := cmd.Flags().GetStringSlice("range")
	switch len(dateRange) {
	case 1:
		query = append(query, "older_than:"+dateRange[0])
	case 2:
		query = append(query, "older_than:"+dateRange[0]+" newer_than:"+dateRange[1])
	}

	// category
	category, _ := cmd.Flags().GetString("category")
	if category != "" {
		query = append(query, "category:"+category)
	}

	// folder
	in, _ := cmd.Flags().GetString("in")
	if in != "" {
		query = append(query, "in:"+in)
	}

	// filename
	filename, _ := cmd.Flags().GetString("filename")
	if filename != "" {
		query = append(query, "filename:"+filename)
	}

	finalQuery := strings.Join(query, " ")
	log.Infof("query statement: %s ", finalQuery)
	return finalQuery
}

/* ========================================== Gdrive ===================================================== */

func NewGDriveArgs() *GDriveObj {
	return &GDriveObj{}
}

type GDriveObj struct {
	Commander
}

func (a *GDriveObj) GetDownload(cmd *cobra.Command) bool {
	return Download(cmd)
}

func (a *GDriveObj) GetPurge(cmd *cobra.Command) bool {
	return Purge(cmd)
}

func (a *GDriveObj) GetEmail(cmd *cobra.Command) string {
	return Email(cmd)
}

func (a *GDriveObj) GetAccountInfo(cmd *cobra.Command) string {
	return AccountInfo(cmd)
}

func (a *GDriveObj) GetQueryGDrive(cmd *cobra.Command) string {
	return QueryGDrive(cmd)
}

//{"Filename", "Enter filename", "name contains "},
//{"MIME Type", "Enter MIME type", "mimeType contains "},
//{"Modified Time", "Enter modified time (e.g., '>2021-01-01')", "modifiedTime > "},
//{"Size", "Enter size in bytes (e.g., '>1000')", "size > "},
//{"Full Text", "Enter text to search within files", "fullText contains "},

func QueryGDrive(cmd *cobra.Command) string {
	var query []string
	// size
	size, _ := cmd.Flags().GetString("size")
	if size != "" {
		query = append(query, "size > "+size)
	}

	// modified time
	age, _ := cmd.Flags().GetString("age")
	if age != "" {
		query = append(query, "modifiedTime > "+age)
	}

	// full text
	text, _ := cmd.Flags().GetString("text")
	if text != "" {
		query = append(query, "fullText contains '"+text+"'")
	}

	// filename
	filename, _ := cmd.Flags().GetString("filename")
	if filename != "" {
		query = append(query, "name contains '"+filename+"'")
	}

	// mimeType
	mimeType, _ := cmd.Flags().GetString("mimeType")
	if mimeType != "" {
		query = append(query, "mimeType contains '"+mimeType+"'")
	}

	finalQuery := strings.Join(query, " ")
	log.Infof("query statement: %s ", finalQuery)
	return finalQuery
}
