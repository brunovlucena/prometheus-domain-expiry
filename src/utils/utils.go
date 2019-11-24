package utils

// Prometheur Exporter - Util.
// Author: Bruno Lucena <bvlg900f@gmail.com>

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// ForInterval helps the scraping
func ForInterval(fn func(), interval time.Duration) {
	for {
		fn()
		time.Sleep(interval)
	}
}

// FailOnError ...
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil { // FIX, Refactor
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
