package parsinglogfiles

import (
	"fmt"
	"regexp"
)

const (
	logLineRegex   = `^(\[TRC\])|^(\[DBG\])|^(\[INF\])|^(\[WRN\])|^(\[ERR\])|^(\[FTL\])\s+([a-zA-Z0-9_ ]*)$`
	fieldsRegex    = `<(?:\**|~*|=*|-*)*>`
	passwordRegex  = `(?mi)\"+[a-zA-Z0-9_ ]*password[a-zA-Z0-9_ ]*\"+`
	artifactsRegex = `(?im)end-of-line[0-9]*`
	userRegex      = `(?:User)\s+([a-zA-Z0-9]+)`
)

func IsValidLine(text string) bool {
	regex := regexp.MustCompile(logLineRegex)
	match := regex.FindStringSubmatch(text)
	if match == nil {
		return false
	}
	return true
}

func SplitLogLine(text string) []string {
	regex := regexp.MustCompile(fieldsRegex)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := regexp.MustCompile(passwordRegex)
	count := 0

	for _, line := range lines {
		match := regex.Find([]byte(line))
		if match != nil {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(artifactsRegex)
	return string(regex.ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	regex := regexp.MustCompile(userRegex)
	var taggedLines []string
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		if match != nil {
			newLine := fmt.Sprintf("[USR] %s %s", match[1], line)
			taggedLines = append(taggedLines, newLine)
		} else {
			taggedLines = append(taggedLines, line)
		}
	}
	return taggedLines
}
