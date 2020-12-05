package changelog

import (
	"fmt"
	"regexp"
)

func GetURL(version string) string {
	path := "https://github.com/leozhantw/gojira"
	r := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, version)
	return url
}
