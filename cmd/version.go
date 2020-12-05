package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/leozhantw/gojira/internal/changelog"
)

var (
	Version   string
	BuildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gojira",
	Run: func(cmd *cobra.Command, args []string) {
		if BuildDate != "" {
			Version = fmt.Sprintf("%s (%s)", Version, BuildDate)
		}

		fmt.Printf("gojira version %s\n%s\n", Version, changelog.GetURL(Version))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
