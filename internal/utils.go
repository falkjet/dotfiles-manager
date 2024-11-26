package internal

import (
	"os"
	"path"
	"unicode/utf8"
)

func readLinkAbs(link string) (location string, err error) {
	location, err = os.Readlink(link)
	if err != nil {
		return "", err
	}
	if !path.IsAbs(location) {
		location = path.Join(path.Dir(link), location)
	}
	return location, nil
}

// Replaces the start of the string ... if the string has more than n runes
func truncateStart(value string, n int) (result string) {
	i := len(value)
	for j := 3; j < n; j++ {
		if i == 0 {
			return value
		}
		_, w := utf8.DecodeRuneInString(value[:i])
		i -= w
	}
	k := i
	for j := 0; j < 3; j++ {
		_, w := utf8.DecodeRuneInString(value[:i])
		i -= w
		if i == 0 {
			return value
		}
	}
	return "..." + value[k:]
}
