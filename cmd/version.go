package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
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

		fmt.Printf("gojira version %s\n%s\n", Version, changelogURL())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func changelogURL() string {
	path := "https://github.com/leozhantw/gojira"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(Version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, Version)
	return url
}
