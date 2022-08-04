package tracker

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tracker",
	Short: "tracker - a simple CLI to track a mount folder files and size",
	Long: `tracker is used to monitor the disk usage and return
   
	a list of all the files on the mountpoint and their disk usage in bytes in json format`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("welcome to the tracker !!")
	return nil
}
