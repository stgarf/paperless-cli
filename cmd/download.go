package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Document",
	Long:  "Download a remote document from a paperless server",

	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Called Download with args %v", args)
		if len(args) < 1 {
			log.Info("Missing filename to download")
		}
		if viper.ConfigFileUsed() == "" {
			fmt.Println("No configuration file found! Try 'config create'")
		} else {
			for index := range args {
				PaperInst.DownloadFiles(args[index])
			}
		}
		log.Debug("Done downloading")
	},
}

func init() {
	documentsCmd.AddCommand(downloadCmd)
}
