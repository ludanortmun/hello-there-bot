package parser

import (
	"regexp"
	"strings"
)

// Parse reads a string message and returns true if it is a variation of "Hello, there"; false otherwise.
// A message is considered a variation of "Hello, there" if, after removing any non alphabetical characters
// and changing it to lowercase, the resulting string is "hellothere", case insensitive.
func Parse(text string) bool {
	alphaCharsPattern := regexp.MustCompile("[^A-Za-z]")
	processed := alphaCharsPattern.ReplaceAllLiteralString(text, "")
	return strings.ToLower(processed) == "hellothere"
}
