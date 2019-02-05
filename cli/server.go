package cli

import (
	"github.com/g-dormoy/thoth/server"
	"github.com/spf13/cobra"
)

var port uint
var storageDir string

func init() {
	serverCmd.PersistentFlags().UintVarP(&port, "port", "p", 4242, "port the server is listening to")
	serverCmd.PersistentFlags().StringVarP(&storageDir, "storage-dir", "", "/var/tmp", "Path of the directory on which data will be saved on flush")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a thoth server",
	Long:  `Start a thoth server on the local machine`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := server.NewConf()
		conf.SetPort(port)
		conf.SetStorageDir(storageDir)
		server.Run(conf)
	},
}
