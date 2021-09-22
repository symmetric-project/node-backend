package slug

import (
	"regexp"
	"strings"
)

var RE *regexp.Regexp

func Slugify(s string) string {
	RE = regexp.MustCompile("[^a-z0-9]+")
	return strings.Trim(RE.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
