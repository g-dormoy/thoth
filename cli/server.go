package cli

import (
	"github.com/g-dormoy/thoth/server"
	"github.com/spf13/cobra"
)

var port uint

func init() {
	serverCmd.PersistentFlags().UintVarP(&port, "port", "p", 4242, "port the server is listening to")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a thoth server",
	Long:  `Start a thoth server on the local machine`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(port)
	},
}
