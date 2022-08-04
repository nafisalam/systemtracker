package tracker

import (
	"encoding/json"
	"fmt"
	"log"
	"systemtracker/pkg/tracker"

	"github.com/spf13/cobra"
)

// subcommand to display the list of files and size
var trackerCmd = &cobra.Command{
	Use:     "/track",
	Aliases: []string{"."},
	Short:   "Track a mount folder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := tracker.FilesDetails(args[0])
		b, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println(string(b))
	},
}

// adding the subcommand tracker to the rootcmd.
func init() {
	rootCmd.AddCommand(trackerCmd)
}
