package main

import (
	"errors"
	"regexp"
)

func urlValidator(input string) error {
	pattern := `^(http[s]?|ftp):\/\/(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(\/\S*)?$`
	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Check if the URL matches the pattern
	if regex.MatchString(input) {
		return nil
	}
	return errors.New("Invalid URL")
}
