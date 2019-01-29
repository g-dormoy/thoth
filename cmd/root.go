package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Thoth",
	Short: "Thoth is a golang based messaging system for event driven architecture",
	Long:  `A messaging system to build event driven architectures coded with <3 and golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello Thoth")
	},
}

// Execute is used to launch the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
