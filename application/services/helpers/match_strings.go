package helpers

import (
	"fmt"
	"html"
	"os"
	"regexp"
)

// MatchHtmlStringsAndSaveTheFile находит строки мостов и заменяет HTML-сущности
func MatchHtmlStringsAndSaveTheFile(htmlContent []byte, file *os.File) (string, error) {
	htmlString := string(htmlContent)

	re := regexp.MustCompile(`obfs4 .*? iat-mode=\d`)

	matches := re.FindAllString(htmlString, -1)
	if matches == nil {
		return "no matching strings were found in the content", nil
	}

	for _, line := range matches {
		decodedLine := html.UnescapeString(line)

		if _, err := file.WriteString(decodedLine + "\n"); err != nil {
			return "", fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return "", nil
}
