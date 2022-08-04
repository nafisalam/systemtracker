package tracker

import (
	"encoding/json"
	"fmt"
	"systemtracker/pkg/tracker"

	"github.com/spf13/cobra"
)

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
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(trackerCmd)
}
