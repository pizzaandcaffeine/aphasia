// TODO: this maybe should be an external package into pkg/ directory
package aphasia

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// getAdjective retrieves a random adjective from a file.
// It reads the lines from the file located at "internal/aphasia/doc/adjectives.txt",
// selects a random line, and returns it as a string.
// If there is an error opening or reading the file, it returns an empty string and the error.
func getAdjective() (string, error) {
	filePath, _ := filepath.Abs("internal/aphasia/doc/adjectives.txt")
	adjective, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer adjective.Close()

	// Print a random line from adjective
	var lines []string
	scanner := bufio.NewScanner(adjective)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[rand.Intn(len(lines))], nil
}

// getNoun is a function that retrieves a random noun from a file.
// It reads the contents of the file located at "internal/aphasia/doc/nouns.txt",
// selects a random line, and returns it as a string.
// If there is an error opening the file or reading its contents, an error is returned.
func getNoun() (string, error) {
	filePath, _ := filepath.Abs("internal/aphasia/doc/nouns.txt")
	noun, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer noun.Close()

	// Print a random line from noun
	var lines []string
	scanner := bufio.NewScanner(noun)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[rand.Intn(len(lines))], nil
}

// Talk generates a mumbled string by combining an adjective and a noun.
// It calls the getAdjective and getNoun functions to retrieve the adjective and noun respectively.
// If any error occurs during the retrieval process, it panics.
// The mumbled string is formed by concatenating the adjective and noun with a hyphen in between.
// It returns the mumbled string.
func Talk() string {
	adjective, err := getAdjective()
	if err != nil {
		panic(err)
	}

	noun, err := getNoun()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s-%s", adjective, noun)
}
